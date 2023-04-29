// lexer/lexer_test.go

package lexer

import (
	"testing"

	"github.com/wule61/derive/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `#[derive(i18n(zh-CN="你好 %s", zh-HK="你好", en="hello", code=400))]`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SHAP, "#"},
		{token.LBRACKET, "["},
		{token.IDENT, "derive"},
		{token.LPAREN, "("},
		{token.IDENT, "i18n"},
		{token.LPAREN, "("},
		{token.IDENT, "zh-CN"},
		{token.ASSIGN, "="},
		{token.STRING, "你好 %s"},
		{token.COMMA, ","},
		{token.IDENT, "zh-HK"},
		{token.ASSIGN, "="},
		{token.STRING, "你好"},
		{token.COMMA, ","},
		{token.IDENT, "en"},
		{token.ASSIGN, "="},
		{token.STRING, "hello"},
		{token.COMMA, ","},
		{token.IDENT, "code"},
		{token.ASSIGN, "="},
		{token.INT, "400"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.RBRACKET, "]"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
