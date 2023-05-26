package i18n

import (
	"fmt"
	"strings"

	. "github.com/dave/jennifer/jen"
)

// TransFnTpl 翻译函数模板

type Lang struct {
	Lang  string
	Value any
}

type ErrorCode struct {
	Value any
	Type  string
}

type Data struct {
	Type        string
	Code        ErrorCode
	Langs       []Lang
	DefaultLang Lang
}

func GenerateCode(data Data, f *File) {

	f.Comment(fmt.Sprintf("// %v_ %v [%v]", data.Type, data.DefaultLang.Value, data.Code.Value))
	f.Var().Id(data.Type + "_").Id(data.Type).Op("=").Lit(data.Code.Value)

	f.Var().Id(strings.ToLower(data.Type) + "Locales").Op("=").Map(String()).String().Values(DictFunc(func(d Dict) {
		for _, v := range data.Langs {
			d[Lit(v.Lang)] = Lit(v.Value)
		}
	}))

	f.Func().Params(Id(data.Type)).Id("Trans").Params(Id("langOrArgs").Op("...").Any()).String().Block(
		List(Id("lang"), Id("args")).Op(":=").Qual("github.com/wule61/derive/utils", "ParseLangArgs").Call(Id("langOrArgs").Op("...")),

		If(List(Id("msg"), Id("ok")).Op(":=").Id(strings.ToLower(data.Type)+"Locales").Index(Id("lang")).Op(";").Id("ok")).Block(
			If(Len(Id("args")).Op(">").Lit(0)).Block(
				Return(Qual("fmt", "Sprintf").Call(Id("msg"), Id("args").Op("..."))),
			),
			Return(Id("msg")),
		),
		Return(Id(strings.ToLower(data.Type)+"Locales").Index(Lit(data.DefaultLang.Lang))),
	)

	f.Func().Params(Id(data.Type)).Id("Code").Params().Id(data.Code.Type).Block(
		Return(Lit(data.Code.Value)),
	)
}
