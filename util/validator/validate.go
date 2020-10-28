package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

var (
	V     *validator.Validate
	trans ut.Translator
)

// 直接使用validator库
func Setup() {
	V = validator.New()
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	uni := ut.New(enT, zhT, enT)
	var ok bool
	trans, ok = uni.GetTranslator("zh")
	if !ok {
		zap.L().Error("translator is not ok")
		return
	}
	// 验证器注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(V, trans)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	registerTagNameFunc()
	registerValidationTranslation([]validationTranslation{
		{
			tag: "checkUsername",
			Fun: checkUsername,
			msg: "用户名必须是由字母开头的4-16位字母和数字组成的字符串",
		},
	})
}

// 以msg方式翻译错误消息
func TranslateErrMsg(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}

// 以data方式翻译错误消息
func TranslateErrData(errs validator.ValidationErrors) map[string]string {
	return removeTopStruct(errs.Translate(trans))
}

// 注册一个获取json tag的自定义方法
func registerTagNameFunc() {
	V.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

type validationTranslation struct {
	tag string
	Fun validator.Func
	msg string
}

// 注册验证方法并翻译
func registerValidationTranslation(vs []validationTranslation) {
	for _, v := range vs {
		registerValidation(v.tag, v.Fun)
		registerTranslation(v.tag, registerTranslator(v.tag, v.msg))
	}
}

// 自定义验证方法
func registerValidation(tag string, fun validator.Func) {
	_ = V.RegisterValidation(tag, fun)
}

// 根据自定义的标记注册翻译
func registerTranslation(tag string, registerFn validator.RegisterTranslationsFunc) {
	_ = V.RegisterTranslation(
		tag,
		trans,
		registerFn,
		translate)
}

// 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.(error).Error()
	}
	return msg
}

// 去除字段名中的结构体名称标识
// refer from:https://github.com/go-playground/validator/issues/633#issuecomment-654382345
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
