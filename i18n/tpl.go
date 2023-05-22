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

const TransFnTpl = `var {{lower .Type}}Locales = map[string]string{
	{{- range $index, $element := .Langs}}
	"{{$element.Lang}}": "{{$element.Value}}",
	{{- end}}
}

func ({{.Type}}) Trans(langOrArgs ...any) string {

	lang, args := utils.ParseLangArgs(langOrArgs...)
  msg := {{lower .Type}}Locales[lang]
	if _, ok := {{lower .Type}}Locales[lang]; !ok {
		msg = {{lower .Type}}Locales["{{.DefaultLang.Lang}}"]
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return msg
}

func ({{.Type}}) Code() {{.Code.Type}} {
	return {{.Code.Value}}
} 
`
