// Code generated by derive; DO NOT EDIT.

package err

import (
 "github.com/wule61/derive/utils" 
 "fmt" 
)

var cardIDNotSpecified_ cardIDNotSpecified = 400

func (cardIDNotSpecified) Trans(langOrArgs ...any) string {

	// 解析语言和参数
	lang, args := utils.ParseLangArgs(langOrArgs...)
	// 返回翻译结果
	switch lang {
	
	case "zh-HK":
		return fmt.Sprintf("卡券 id 未选择", args...)
	
	case "zh-CN":
		return fmt.Sprintf("卡券 id 未选择", args...)
	
	case "en":
		return fmt.Sprintf("卡券 id 未选择", args...)
	
	default:
		return fmt.Sprintf("卡券 id 未选择", args...)
	}
}

func (cardIDNotSpecified) Code() int {
	return 400
} 


var TeamNotFound_ TeamNotFound = 400

func (TeamNotFound) Trans(langOrArgs ...any) string {

	// 解析语言和参数
	lang, args := utils.ParseLangArgs(langOrArgs...)
	// 返回翻译结果
	switch lang {
	
	case "zh-HK":
		return fmt.Sprintf("团队未找到", args...)
	
	case "zh-CN":
		return fmt.Sprintf("团队未找到", args...)
	
	case "en":
		return fmt.Sprintf("团队未找到", args...)
	
	default:
		return fmt.Sprintf("团队未找到", args...)
	}
}

func (TeamNotFound) Code() int {
	return 400
} 
