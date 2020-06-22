package gogen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {
	Main()
}

func dirFilter(f os.FileInfo) bool { return nameFilter(f.Name()) }

func nameFilter(filename string) bool {
	switch filename {
	case "gogen.go", "gogen_test.go":
		return false
	}
	return true
}

func Test_Parse(t *testing.T) {
	fset := token.NewFileSet()
	dir, _ := parser.ParseDir(fset, `D:\gopath\src\gogen\astruct`, dirFilter, 0)
	for _, a := range dir {
		for _, file := range a.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				// Find Return Statements
				ret, ok := n.(*ast.Field)
				if ok {
					if strings.Contains(ret.Tag.Value, "fixType") {
						fmt.Printf("return statement found on line %v:\n", fset.Position(ret.Pos()))
						fmt.Printf("tag with fixType statement found on line %v:\n", fset.Position(ret.Tag.Pos()))
						printer.Fprint(os.Stdout, fset, ret)
						printer.Fprint(os.Stdout, fset, ret.Tag)
						fmt.Printf("\n")
						return true
					}
				}
				return true
			})
		}
	}

	ast.Print(fset, dir)
}

func Test_PrintStringGen(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, `D:\gopath\src\gogen\astruct\ss_string_gen.go`, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}
	ast.Print(fset, file)
}
