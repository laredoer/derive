// Package utils ...
package utils

// ParseLangArgs 解析语言和参数
func ParseLangArgs(langOrArgs ...any) (lang string, args []any) {

	if len(langOrArgs) > 0 {
		//如果第一个参数是字符串，则认为是语言
		if lang, ok := langOrArgs[0].(string); ok {
			return lang, langOrArgs[1:]
		}
	}

	//否则，返回默认值
	return "zh-HK", nil
}
