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
}

func Init(storagePath string) {
	Storage = storage{
		AbsPath:  filepath.Join(storagePath, conf.Conf.Filesystems.Disks.Local.Root),
		DiskName: conf.Conf.Filesystems.Default,
	}
}

func (Storage *storage) Disk(diskName string) *storage {
	Storage.DiskName = diskName
	return Storage
}

func (Storage *storage) FullPath(path string) string {
	return filepath.Join(Storage.AbsPath, path)
}

/**
判断文件是否存在, 使用的是文件的保存路径 storage/app/public/ 下
*/
func (Storage *storage) Exists(path string) bool {
	fullPath := Storage.FullPath(path)
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
保存base64字符串为文件
*/
func (Storage *storage) StoreBase64RandomName(path string, ImgBase64 string) (string, error) {
	b64data := ImgBase64[strings.IndexByte(ImgBase64, ',')+1:]
	fileContent, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return "", err
	}
	fullPath := Storage.FullPath(path)
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

/**
文件转base64
*/
func (Storage *storage) FileToBase64(path string) (string, error) {
	fileContent, err := ioutil.ReadFile(Storage.FullPath(path))
	if err != nil {
		return "", err
	}
	base64String := base64.StdEncoding.EncodeToString(fileContent)
	return base64String, nil
}
