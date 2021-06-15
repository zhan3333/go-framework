package storage

import (
	"encoding/base64"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go-framework/conf"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Storage storage

type storage struct {
	// public 目录的文件夹路径
	AbsPath  string
	DiskName string
	Uri      string
}

func Init(storagePath string) {
	Storage = storage{
		AbsPath:  filepath.Join(storagePath, conf.Filesystems.Disks.Local.Root),
		DiskName: conf.Filesystems.Default,
		Uri:      conf.Url,
	}
}

func (s *storage) Disk(diskName string) *storage {
	s.DiskName = diskName
	return s
}

func (s *storage) FullPath(path string) string {
	return filepath.Join(s.AbsPath, path)
}

func (s *storage) Url(path string) string {
	if path == "" {
		return ""
	}
	return fmt.Sprintf("%s/public/%s", s.Uri, path)
}

// Exists 判断文件是否存在, 使用的是文件的保存路径 storage/app/public/ 下
func (s *storage) Exists(path string) bool {
	fullPath := s.FullPath(path)
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// StoreBase64RandomName 保存base64字符串为文件
func (s *storage) StoreBase64RandomName(path string, ImgBase64 string) (string, error) {
	b64data := ImgBase64[strings.IndexByte(ImgBase64, ',')+1:]
	fileContent, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return "", err
	}
	fullPath := s.FullPath(path)
	fileName := fmt.Sprintf("%s.jpg", uuid.NewV4().String())
	filePath := filepath.Join(fullPath, fileName)
	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filePath, fileContent, 0777)
	if err != nil {
		return "", err
	}
	return filepath.Join(path, fileName), nil
}

// FileToBase64 将文件转换为 base64 字符串
func (s *storage) FileToBase64(path string) (string, error) {
	fileContent, err := ioutil.ReadFile(s.FullPath(path))
	if err != nil {
		return "", err
	}
	base64String := base64.StdEncoding.EncodeToString(fileContent)
	return base64String, nil
}

// GetBase64SizeMb 获取base64图片的尺寸 Mb 为单位
func (s *storage) GetBase64SizeMb(base64 string) float64 {
	deviation := 0.0
	length := float64(len(base64))
	return length / 1024.0 / 1024.0 * (1.0 - deviation)
}

// GetBase64SizeKb 获取base64图片的尺寸 kb 为单位
func (s *storage) GetBase64SizeKb(base64 string) float64 {
	deviation := 0.0
	length := float64(len(base64))
	return length / 1024.0 * (1.0 - deviation)
}
