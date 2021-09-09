package wxwork

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

// Media 素材
type Media struct {
	baseCaller
	Type      string `json:"type,omitempty"`       // 文件类型,image、voice、video、file
	MediaId   string `json:"media_id,omitempty"`   // 唯一标识，3天内有效
	CreatedAt string `json:"created_at,omitempty"` // 上传时间戳
}

// UploadMediaWithType 上传临时素材
func (a *Agent) UploadMediaWithType(mediaType string, buf []byte, info os.FileInfo) (media Media, err error) {

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	fw, err := writer.CreateFormFile("media", info.Name())
	if err != nil {
		return
	}

	_, err = fw.Write(buf)
	if err != nil {
		return
	}

	_ = writer.WriteField("filename", info.Name())
	_ = writer.WriteField("filelength", strconv.FormatInt(info.Size(), 10))
	_ = writer.Close()

	accessToken, err := a.GetAccessToken()
	if err != nil {
		return
	}

	query := url.Values{}
	query.Set("access_token", accessToken)
	query.Set("type", mediaType)

	u, err := url.Parse(BaseURL + "media/upload")
	if err != nil {
		return
	}

	u.RawQuery = query.Encode()

	resp, err := a.client.Post(u.String(), writer.FormDataContentType(), buffer)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return
	}

	if !media.Success() {
		return media, media.Error()
	}

	return media, nil
}

// MediaUpload 上传临时素材并获取素材信息
// 参数 file 为素材位置
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90253
func (a *Agent) MediaUpload(file string) (media Media, err error) {
	info, err := os.Stat(file)
	if err != nil {
		return
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	var mediaType string

	switch filepath.Ext(info.Name()) {
	case ".jpg", ".png":
		mediaType = "image"
	case ".arm":
		mediaType = "voice"
	case ".mp4":
		mediaType = "video"
	default:
		mediaType = "file"
	}

	return a.UploadMediaWithType(mediaType, buf, info)
}

// UploadImgWithType 上传图片
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90256
func (a *Agent) UploadImgWithType(buf []byte, info os.FileInfo) (string, error) {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	fw, err := writer.CreateFormFile("media", info.Name())
	if err != nil {
		return "", err
	}

	fw.Write(buf)

	writer.WriteField("name", info.Name())
	writer.WriteField("filename", info.Name())
	writer.WriteField("filelength", strconv.FormatInt(info.Size(), 10))
	writer.Close()

	accessToken, err := a.GetAccessToken()
	if err != nil {
		return "", err
	}

	query := url.Values{}
	query.Set("access_token", accessToken)

	u, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}

	u.Path = path.Join(u.Path, "media/uploadimg")

	u.RawQuery = query.Encode()

	resp, err := a.client.Post(u.String(), writer.FormDataContentType(), buffer)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var caller struct {
		baseCaller
		Url string
	}

	if err = json.NewDecoder(resp.Body).Decode(&caller); err != nil {
		return "", err
	}

	if !caller.Success() {
		return "", caller.Error()
	}

	return caller.Url, err
}

// UploadImg 上传图片
func (a *Agent) UploadImg(file string) (string, error) {
	info, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return a.UploadImgWithType(buf, info)
}
