package derive

import (
	"bytes"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type Value struct {
	OriginalName string // The name of the constant.
	Name         string // The name with trimmed prefix.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or an uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	Value     uint64 // Will be converted to int64 when needed.
	Signed    bool   // Whether the constant is a signed type.
	Str       string // The string representation given by the "go/constant" package.
	BasicType string // value of basic Type, for: int int64 uint etc
}

type File struct {
	Pkg      *Package // Package to which this file belongs.
	FileName string
	AstFile  *ast.File // Parsed AST.
}

type Package struct {
	Name    string // 包名
	PkgPath string
	Defs    map[*ast.Ident]types.Object // 一个包的所有定义共享
	File    []*File                     // 一个包可能有多个文件
}

type Generator struct {
	Buf bytes.Buffer // Accumulated output.
	Pkg *Package     // Package we are scanning.
}

func (g *Generator) AddPackage(pkg *packages.Package) {
	g.Pkg = &Package{
		Name:    pkg.Name,
		PkgPath: pkg.PkgPath,
		Defs:    pkg.TypesInfo.Defs,
		File:    make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.Pkg.File[i] = &File{
			Pkg:      g.Pkg,
			FileName: pkg.Fset.File(file.Pos()).Name(),
			AstFile:  file,
		}
	}
}
