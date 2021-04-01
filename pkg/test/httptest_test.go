package test_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-framework/pkg/test"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		t.Log("ok")
		c.String(http.StatusOK, "ok")
	})
	n := test.New(router)
	w := n.Get("/test", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestHttp_GetWithBody(t *testing.T) {
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		type T struct {
			Test string `json:"test" form:"test"`
		}
		assert.Equal(t, "application/json", c.ContentType())
		t.Log("ok")
		t.Logf("path: %s", c.FullPath())
		var tt T
		assert.Nil(t, c.ShouldBind(&tt))
		assert.Equal(t, "test", tt.Test)
		c.String(http.StatusOK, "ok")
	})
	n := test.New(router)
	w := n.Get("/test", gin.H{
		"test": "test",
	})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}
