package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"regexp"
	"strings"
)

var (
	V     *validator.Validate
	trans ut.Translator
)

func Setup() {
	V = validator.New()
	registerValidation()
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	uni := ut.New(enT, zhT, enT)
	var ok bool
	trans, ok = uni.GetTranslator("zh")
	if !ok {
		return
	}
	// 验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(V, trans)
	if err != nil {
		return
	}
	registerTranslation()
}

// 翻译错误消息
func TranslateErr(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}

// 自定义验证方法
func registerValidation() {
	_ = V.RegisterValidation("checkUsername", checkUsername)
}

// 根据自定义的标记注册翻译
func registerTranslation() {
	_ = V.RegisterTranslation(
		"checkUsername",
		trans,
		registerTranslator("checkUsername", "{0}必须是由字母开头的4-16位字母和数字组成的字符串"),
		translate)
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func checkUsername(fl validator.FieldLevel) bool {
	if ok, _ := regexp.MatchString(`^[a-zA-Z]{1}[a-zA-Z0-9]{3,15}$`, fl.Field().String()); !ok {
		return false
	}
	return true
}
