package derive

import (
	"github.com/wule61/derive/interpreter/ast"
	"github.com/wule61/derive/interpreter/lexer"
	"github.com/wule61/derive/interpreter/parser"
)

type Derive struct {
	Name   string
	Params []DeriveParam
}

type DeriveParam struct {
	Name  string
	Value any
	Type  string
}

func ParseCommentToDerive(comment string) []Derive {

	l := lexer.New(comment)
	p := parser.New(l)
	program := p.ParseProgram()

	var derives []Derive
	for _, stmt := range program.Statements {
		shap, ok := stmt.(*ast.ShapStatement)
		if !ok {
			continue
		}

		for _, call := range shap.Elements {
			switch call.(type) {
			case *ast.CallExpression:
				derive := Derive{
					Name: call.(*ast.CallExpression).Function.(*ast.Identifier).Value,
				}
				arguments := call.(*ast.CallExpression).Arguments
				for _, arg := range arguments {
					deriveParam := DeriveParam{
						Name: arg.(*ast.AssignLiteral).Name.(*ast.Identifier).Value,
					}
					switch arg.(*ast.AssignLiteral).Expression.(type) {
					case *ast.StringLiteral:
						deriveParam.Value = arg.(*ast.AssignLiteral).Expression.(*ast.StringLiteral).Value
						deriveParam.Type = "string"
					case *ast.IntegerLiteral:
						deriveParam.Value = arg.(*ast.AssignLiteral).Expression.(*ast.IntegerLiteral).Value
						deriveParam.Type = "int"
					}
					derive.Params = append(derive.Params, deriveParam)
				}

				derives = append(derives, derive)
			case *ast.Identifier:
				derives = append(derives, Derive{
					Name:   call.(*ast.Identifier).Value,
					Params: []DeriveParam{},
				})
			}
		}
	}

	return derives
}
