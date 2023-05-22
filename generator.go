package derive

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/wule61/derive/i18n"

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
	Types    map[string]Type
}

type Type struct {
	Type    string
	Derives []Derive
}

func (f *File) AddDerive(tName, tType string, derives []Derive) {
	f.Types[tName] = Type{
		Type:    tType,
		Derives: derives,
	}
}

func (f *File) GenDecl(node ast.Node) bool {

	fileNode, ok := node.(*ast.File)
	if !ok {
		// We only care about const declarations.
		return true
	}

	for _, spec := range fileNode.Decls {
		gDecl, ok := spec.(*ast.GenDecl)
		if !ok {
			continue
		}
		var comments string
		if gDecl.Doc == nil || len(gDecl.Doc.List) == 0 {
			continue
		}

		if gDecl.Doc != nil {
			for _, comment := range gDecl.Doc.List {
				comments += comment.Text
			}
		}

		var typ string
		var tType string
		if spec := gDecl.Specs[0]; spec != nil {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			typ = typeSpec.Name.Name
			tType = typeSpec.Type.(*ast.Ident).Name
		}

		f.AddDerive(typ, tType, ParseCommentToDerive(comments))
	}

	return false
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
			Types:    make(map[string]Type),
		}
	}
}

func (g *Generator) Generate(file *File) {

	var buffer bytes.Buffer
	buffer.WriteString(`// Code generated by derive; DO NOT EDIT.`)
	buffer.Write([]byte("\n\n"))
	buffer.WriteString(`package ` + file.Pkg.Name)
	// import
	buffer.Write([]byte("\n\n"))
	buffer.WriteString("import (\n \"github.com/wule61/derive/utils\" \n \"fmt\" \n)")
	for typ, derives := range file.Types {
		for _, derive := range derives.Derives {
			if derive.Name == "i18n" {
				data := i18n.TransFnTplData{
					Type: typ,
					Code: i18n.Code{},
				}
				for _, v := range derive.Params {
					if v.Name == "code" {
						data.Code = i18n.Code{
							Type:  derives.Type,
							Value: v.Value,
						}
						continue
					}
					if v.Name == "zh-HK" {
						data.DefaultLang = i18n.Lang{
							Lang:  v.Name,
							Value: v.Value,
						}
					}

					data.Langs = append(data.Langs, i18n.Lang{
						Lang:  v.Name,
						Value: v.Value,
					})
				}

				buffer.Write([]byte("\n\n"))
				buffer.WriteString(fmt.Sprintf("var %v_ %v = %d", typ, typ, data.Code.Value))
				buffer.Write([]byte("\n\n"))

				tmpl, err := template.New("i18n_trans_fn").Parse(i18n.TransFnTpl)
				if err != nil {
					panic(err)
				}

				err = tmpl.Execute(&buffer, data)
				if err != nil {
					panic(err)
				}

				err = WriteToFile(g.GetFileName(file.FileName, derive.Name), buffer.Bytes())
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// GetFileName  获取要生成的文件名称
func (g *Generator) GetFileName(fileName, deriveName string) string {

	arr := strings.Split(fileName, ".")
	// 以最后一个点号为分割
	if len(arr) > 1 {
		fileName = strings.Join(arr[:len(arr)-1], ".")
	}

	return fileName + "_" + deriveName + ".go"

}

func WriteToFile(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, io.SeekEnd)
		_, err = f.WriteAt(content, n)
		fmt.Println("write succeed!")
		defer f.Close()
	}
	return err
}
