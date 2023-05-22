package main

import (
	"flag"
	"fmt"
	"go/ast"
	"os"

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
		for _, file := range g.Pkg.File {
			ast.Inspect(file.AstFile, file.GenDecl)
			// 生成代码
			g.Generate(file)
		}
	}
}
