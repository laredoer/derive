# derive

Derive is a project written in Go that allows users to utilize macro functionality similar to that of Rust. With Derive, developers can simplify their codebase by defining and invoking macros to generate repetitive or boilerplate code. Derive aims to improve the developer experience by reducing code redundancy and increasing code readability. It provides an intuitive interface for defining and invoking macros, making it an excellent tool for developers looking to improve their productivity in Go programming.

### i18n


```go
// .../message/message.go
package message

// #[i18n(code = 400, zh-HK = "年卡", zh-CN = "年卡", en = "Year card")]
type YearCard int32
```

```go
// .../message/message_i18n.go

// Code generated by derive; DO NOT EDIT.
package message

import (
	"fmt"
	"github.com/wule61/derive/utils"
)

var YearCard_ YearCard = 400
var yearcardLocales = map[string]string{
	"en":    "Year card",
	"zh-CN": "年卡",
	"zh-HK": "年卡",
}

func (YearCard) Trans(langOrArgs ...any) string {
	lang, args := utils.ParseLangArgs(langOrArgs...)
	if msg, ok := yearcardLocales[lang]; ok {
		if len(args) > 0 {
			return fmt.Sprintf(msg, args...)
		}
		return msg
	}
	return yearcardLocales["zh-HK"]
}

func (YearCard) Code() int32 {
	return 400
}
```

**i18n** is a **derive** that can automatically generate code for internationalization based on the annotations in the structure. In this code, the **YearCard** structure contains the necessary multilingual parameters and error codes for internationalization. derive will parse the information in the annotation and generate the **Trans** and **Code** methods. This makes it easier and more convenient to use multilingualism in the code with the help of **github copilot**, without having to manually translate and write multilingual versions of the code or yaml files. And the multilingual parameters are free, so you can use parameters like **zh_HK**, **en-US** to represent your language.