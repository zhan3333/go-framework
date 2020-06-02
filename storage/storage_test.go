package storage_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-framework/app"
	boot "go-framework/bootstrap"
	"go-framework/storage"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Bootstrap()
	m.Run()
}

func TestStorageSaveBase64(t *testing.T) {
	frontImgBase64, _ := ioutil.ReadFile(filepath.Join(app.TestPath, "testdata", "id-card-back-base64.txt"))
	savePath, err := storage.Storage.StoreBase64RandomName(fmt.Sprintf("id_card_imgs/%s", "123456"), string(frontImgBase64))
	log.Println(savePath, err)
	assert.Nil(t, err)
	assert.NotEqual(t, "", savePath)
	assert.True(t, storage.Storage.Exists(savePath))
}

func TestGetStorageAbsPath(t *testing.T) {
	log.Println(storage.Storage.AbsPath)
}

func TestFileToBase64(t *testing.T) {
	frontImgBase64, _ := ioutil.ReadFile("./testdata/id-card-back-base64.txt")
	savePath, err := storage.Storage.StoreBase64RandomName(fmt.Sprintf("id_card_imgs/%s", "123456"), string(frontImgBase64))
	log.Println(savePath, err)
	assert.Nil(t, err)
	assert.NotEqual(t, "", savePath)
	assert.True(t, storage.Storage.Exists(savePath))
	base64String, err := storage.Storage.FileToBase64(savePath)
	assert.Nil(t, err)
	assert.NotEmpty(t, base64String)
	log.Println(base64String)
}
