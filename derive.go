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
			switch call.(type) {
			case *ast.CallExpression:
				derive := DeriveType{
					Name: call.(*ast.CallExpression).Function.(*ast.Identifier).Value,
				}
				arguments := call.(*ast.CallExpression).Arguments
				for _, arg := range arguments {
					deriveParam := Field{
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
					derive.Args = append(derive.Args, deriveParam)
				}

				derives = append(derives, derive)
			case *ast.Identifier:
				derives = append(derives, DeriveType{
					Name: call.(*ast.Identifier).Value,
					Args: []Field{},
				})
			}
		}
	}

	return derives
}
