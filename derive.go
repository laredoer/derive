package derive

import (
	"github.com/wule61/derive/interpreter/ast"
	"github.com/wule61/derive/interpreter/lexer"
	"github.com/wule61/derive/interpreter/parser"
)

type DeriveType struct {
	Name string
	Args []Field
}

type Field struct {
	Name  string
	Type  string
	Value any
}

func ParseCommentToDerive(comment string) []DeriveType {

	l := lexer.New(comment)
	p := parser.New(l)
	program := p.ParseProgram()

	var derives []DeriveType
	for _, stmt := range program.Statements {
		shap, ok := stmt.(*ast.ShapStatement)
		if !ok {
			continue
		}

		for _, call := range shap.Elements {
			switch callType := call.(type) {
			case *ast.CallExpression:
				derive := DeriveType{
					Name: callType.Function.(*ast.Identifier).Value,
				}
				arguments := callType.Arguments
				for _, arg := range arguments {
					deriveParam := Field{}
					switch argType := arg.(type) {
					case *ast.AssignLiteral:
						deriveParam.Name = argType.Name.(*ast.Identifier).Value
						switch argType.Expression.(type) {
						case *ast.StringLiteral:
							deriveParam.Value = argType.Expression.(*ast.StringLiteral).Value
							deriveParam.Type = "string"
						case *ast.IntegerLiteral:
							deriveParam.Value = argType.Expression.(*ast.IntegerLiteral).Value
							deriveParam.Type = "int"
						}
					case *ast.Identifier:
						deriveParam.Name = argType.Value
					}

					derive.Args = append(derive.Args, deriveParam)
				}

				derives = append(derives, derive)
			case *ast.Identifier:
				derives = append(derives, DeriveType{
					Name: callType.Value,
					Args: []Field{},
				})
			}
		}
	}

	return derives
}
