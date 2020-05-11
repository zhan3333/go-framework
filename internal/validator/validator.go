package validator

import (
	"github.com/gin-gonic/gin/binding"
	zh_cn "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"regexp"
)

var Trans ut.Translator
var Validate *validator.Validate

func Init() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		Validate = validate
		zh := zh_cn.New()
		uni := ut.New(zh, zh)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		Trans, _ = uni.GetTranslator("zh")
		_ = zh_translations.RegisterDefaultTranslations(validate, Trans)

		// 将 comment 标签注册为参数的翻译
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})

		register()
	}
}

// 定义验证规则
var Mobile validator.Func = func(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())
	return ok
}

func register() {
	// 注册验证规则
	_ = Validate.RegisterValidation("mobile", Mobile)

	// 注册验证规则的翻译
	_ = Validate.RegisterTranslation("mobile", Trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0}格式不正确", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})
}
