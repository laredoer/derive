package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/wule61/derive"
	"golang.org/x/tools/go/packages"
)

var (
	dir = flag.String("dir", "./...", "the directory you want to execute")
)

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Derive is a macro tool, like rust macro:\n")
	flag.PrintDefaults()
}

func main() {

	flag.Usage = usage
	flag.Parse()

	pkgs, err := packages.Load(&packages.Config{
		Mode:  packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedModule,
		Tests: false,
	}, *dir)
	if err != nil {
		panic(err)
	}

	// 一个包解析一次
	for _, pkg := range pkgs {
		g := &derive.Generator{}
		g.AddPackage(pkg)

		println("---------------------------")
		spew.Dump(g)
	}

	// ast.Inspect(pkg.Syntax[0], genDecl)
}

func genDecl(node ast.Node) bool {

	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.CONST {
		// We only care about const declarations.
		return true
	}

	var typ string
	for _, spec := range decl.Specs {
		vSpec := spec.(*ast.ValueSpec)
		if vSpec.Type == nil && len(vSpec.Values) > 0 {
			typ = ""
			bl, ok := vSpec.Values[0].(*ast.BasicLit)
			if !ok {
				continue
			}
			typ = strings.ToLower(bl.Kind.String())
			if typ == "float" {
				typ = "float64"
			}
		}

		if vSpec.Type != nil {
			ident, ok := vSpec.Type.(*ast.Ident)
			if !ok {
				continue
			}
			typ = ident.Name
		}

		println("----------------------------")
		if vSpec.Doc == nil || len(vSpec.Doc.List) == 0 {
			continue
		}

		var comments string
		for _, comment := range vSpec.Doc.List {
			comments += comment.Text
		}

		spew.Dump(derive.ParseCommentToDerive(comments))

		for _, name := range vSpec.Names {
			if name.Name == "_" {
				continue
			}

		}

		println("----------------------------")

	}

	return false
}
