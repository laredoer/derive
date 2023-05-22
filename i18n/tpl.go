package i18n

// TransFnTpl 翻译函数模板

type Lang struct {
	Lang  string
	Value any
}

type Code struct {
	Value any
	Type  string
}

type TransFnTplData struct {
	Type        string
	Code        Code
	Langs       []Lang
	DefaultLang Lang
}

const TransFnTpl = `func ({{.Type}}) Trans(langOrArgs ...any) string {

	// 解析语言和参数
	lang, args := utils.ParseLangArgs(langOrArgs...)
	// 返回翻译结果
	switch lang {
	{{range $index, $element := .Langs}}
	case "{{$element.Lang}}":
		return fmt.Sprintf("{{$element.Value}}", args...)
	{{end}}
	default:
		return fmt.Sprintf("{{.DefaultLang.Value}}", args...)
	}
}

func ({{.Type}}) Code() {{.Code.Type}} {
	return {{.Code.Value}}
} 
`
