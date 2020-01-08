package wxwork

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type Media struct {
	baseCaller
	Type      string `json:"type"`       // 文件类型,image、voice、video、file
	MediaId   string `json:"media_id"`   // 唯一标识，3天内有效
	CreatedAt string `json:"created_at"` // 上传时间戳
}

// UploadMediaWithType 上传临时素材
func (a *Agent) UploadMediaWithType(mediaType string, buf []byte, info os.FileInfo) (*Media, error) {

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	fw, err := writer.CreateFormFile("media", info.Name())
	if err != nil {
		return nil, err
	}

	fw.Write(buf)

	writer.WriteField("filename", info.Name())
	writer.WriteField("filelength", strconv.FormatInt(info.Size(), 10))
	writer.Close()

	accessToken, err := a.GetAccessToken()
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("access_token", accessToken)
	query.Set("type", mediaType)

	u, err := url.Parse(BaseURL + "media/upload")
	if err != nil {
		return nil, err
	}

	u.RawQuery = query.Encode()

	resp, err := a.client.Post(u.String(), writer.FormDataContentType(), buffer)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	media := &Media{}

	if err = json.NewDecoder(resp.Body).Decode(&media); err != nil {
		return nil, err
	}

	if !media.Success() {
		return nil, media.Error()
	}

	return media, nil
}

// MediaUpload 上传临时素材并获取素材信息
// 参数 file 为素材位置
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90253
func (a *Agent) MediaUpload(file string) (*Media, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
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
