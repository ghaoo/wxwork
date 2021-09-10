package main

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/sirupsen/logrus"
	r "math/rand"
	"os"
	"path"
	"path/filepath"
	"time"
)

func MkDir(filepath string) error {

	if _, err := os.Stat(filepath); err != nil {

		if os.IsNotExist(err) {

			err = os.MkdirAll(filepath, os.ModePerm)

			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func CreateFile(dir string, name string) (string, error) {
	src := path.Join(dir, name)

	_, err := os.Stat(src)

	if os.IsExist(err) {
		return src, nil
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		if os.IsPermission(err) {
			return "", errors.New("你不够权限创建文件")
		}
		return "", err
	}

	_, err = os.Create(src)
	if err != nil {
		return "", err
	}

	return src, nil
}

func WriteFile(file string, text []byte) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.Write(text)
	if err != nil {
		return err
	}
	return w.Flush()
}

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

// Xmd5 计算字符串的md5值
func Xmd5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func storeResult(filename string, data []byte) string {
	basePath := os.Getenv("DATA_PATH")
	if basePath == "" {
		basePath = ".data"
	}

	fpath := filepath.Join(basePath, "tmp", filename)

	os.Remove(fpath)

	MkDir(filepath.Dir(fpath))

	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.Write(data)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	w.Flush()

	return fpath
}
