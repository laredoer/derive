package parser

import (
	"testing"

	"github.com/wule61/derive/interpreter/lexer"
)

func Test_Parser(t *testing.T) {
	str := `#[i18n(zh-HK = "你好",en = "hello", code = 400), debug] #[Clone]`
	//str := `i18n`
	l := lexer.New(str)
	p := New(l)
	program := p.ParseProgram()
	t.Log(program)
}
