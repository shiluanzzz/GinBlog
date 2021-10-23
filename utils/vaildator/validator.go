package vaildator

import (
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	uTran "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	// 自动翻译语言 固定写法
	uni := uTran.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_Cn")
	// 注册校验中的语言
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	// 返回错误的时候就是返回的标签值而不是username这种变量名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})
	// 验证
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
