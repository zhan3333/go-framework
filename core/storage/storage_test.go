package storage_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-framework/app"
	"go-framework/core/boot"
	storage2 "go-framework/core/storage"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	if err := boot.New(
		boot.WithConfigFile(os.Getenv("LGO_TEST_FILE")),
		boot.WithRoutePrint(false),
	); err != nil {
		panic(err)
	}
	m.Run()
}

func TestStorageSaveBase64(t *testing.T) {
	frontImgBase64, _ := ioutil.ReadFile(filepath.Join(app.TestPath, "testdata", "id-card-back-base64.txt"))
	savePath, err := storage2.Storage.StoreBase64RandomName(fmt.Sprintf("id_card_imgs/%s", "123456"), string(frontImgBase64))
	log.Println(savePath, err)
	assert.Nil(t, err)
	assert.NotEqual(t, "", savePath)
	assert.True(t, storage2.Storage.Exists(savePath))
}

func TestGetStorageAbsPath(t *testing.T) {
	log.Println(storage2.Storage.AbsPath)
}

func TestFileToBase64(t *testing.T) {
	frontImgBase64, _ := ioutil.ReadFile("./testdata/id-card-back-base64.txt")
	savePath, err := storage2.Storage.StoreBase64RandomName(fmt.Sprintf("id_card_imgs/%s", "123456"), string(frontImgBase64))
	assert.Nil(t, err)
	assert.NotEqual(t, "", savePath)
	assert.True(t, storage2.Storage.Exists(savePath))
	base64String, err := storage2.Storage.FileToBase64(savePath)
	assert.Nil(t, err)
	assert.NotEmpty(t, base64String)
}
