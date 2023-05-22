// Code generated by derive; DO NOT EDIT.

package err

import (
 "github.com/wule61/derive/utils" 
 "fmt" 
)

var cardIDNotSpecified_ cardIDNotSpecified = 400

var cardIDNotSpecifiedLocales = map[string]string{
	"zh-HK": "卡券 id 未选择",
	"zh-CN": "卡券 id 未选择",
	"en": "卡券 id 未选择",
}

func (cardIDNotSpecified) Trans(langOrArgs ...any) string {

	lang, args := utils.ParseLangArgs(langOrArgs...)
  msg := cardIDNotSpecifiedLocales[lang]
	if _, ok := cardIDNotSpecifiedLocales[lang]; !ok {
		msg = cardIDNotSpecifiedLocales["zh-HK"]
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return msg
}

func (cardIDNotSpecified) Code() int32 {
	return 400
} 


var TeamNotFound_ TeamNotFound = 400

var TeamNotFoundLocales = map[string]string{
	"zh-HK": "团队未找到",
	"zh-CN": "团队未找到",
	"en": "团队未找到",
}

func (TeamNotFound) Trans(langOrArgs ...any) string {

	lang, args := utils.ParseLangArgs(langOrArgs...)
  msg := TeamNotFoundLocales[lang]
	if _, ok := TeamNotFoundLocales[lang]; !ok {
		msg = TeamNotFoundLocales["zh-HK"]
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return msg
}

func (TeamNotFound) Code() int32 {
	return 400
} 
