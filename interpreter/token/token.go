// Package token defines the token type and the token struct.
package token

// TokenType represents the type of a token.
type TokenType string

// Token represents a token in the source code.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 123456
	STRING = "STRING" // "hello. this is string."

	ASSIGN = "="
	COMMA  = ","
	SHAP   = "#"

	BANG  = "!"
	MINUS = "-"

	LPAREN = "("
	RPAREN = ")"

	LBRACKET  = "["
	RBRACKET  = "]"
	SEMICOLON = ";"

	TRUE  = "true"
	FALSE = "false"
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
