package gogen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"testing"
)

func dirFilter(f os.FileInfo) bool { return nameFilter(f.Name()) }

func nameFilter(filename string) bool {
	switch filename {
	case "gogen.go", "gogen_test.go":
		return false
	}
	return true
}

const (
	conf, bg                                 = 0, 0
	testRed, testGreen, testBlue, testPurple = 31, 32, 34, 35
)

func FmtColor(format string, clr int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, clr, format, 0x1B)
}

type FindContext struct {
	/*File      string
	Package   string
	LocalFunc *ast.FuncDecl*/
	structName string
}

var GFset *token.FileSet // 全局存储token的position

func fixTypeInTag(string2 string) string {
	split := strings.Split(string2, `"`)
	if len(split) > 1 {
		return split[1]
	}
	return ""
}

func (f *FindContext) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return f
	}

	switch ele := n.(type) {
	case *ast.Field:
		if ele == nil || ele.Tag == nil || !strings.Contains(ele.Tag.Value, "fixType") {
			// tag中不包含fixType的
			break
		}
		fixValue := fixTypeInTag(ele.Tag.Value)
		fmt.Println(fixValue)
		fmt.Printf("变长字段 on line %v:\n", GFset.Position(ele.Pos()))
		ast.Print(GFset, ele)

		// feildName := ele.Names[0].Name // 不重要
		// optional struct
		// var struct
		// var []byte
		// var string
		switch typeInfo := ele.Type.(type) {
		case *ast.Ident: // 单个
			if typeInfo.Name == "string" {
				// string ---> [XLen]byte
			}
		case *ast.ArrayType: // array
			// []xxx
			switch arrEle := typeInfo.Elt.(type) {
			case *ast.SelectorExpr:
				// []xxx.xxx 外部包
			case *ast.Ident:
				// 内部
				switch arrEle.Name {
				case "byte":
					// []byte
				default:
					// []others
				}
			default:
				log.Printf("%T not accept", arrEle)
			}
		default:
			log.Printf("%T not accept", typeInfo)
		}
	case *ast.StructType:
	// structName
	// ast.Print(GFset, ele)
	case *ast.Ident:
		fmt.Printf("*ast.Ident on line %v:\n", GFset.Position(ele.Pos()))
		if ele.Name == "A"{
			ast.Print(GFset, ele)
		}
	}
	//ret, ok := n.(*ast.Field)
	//if ok && ret != nil && ret.Tag != nil {
	//	if strings.Contains(ret.Tag.Value, "fixType") && strings.Contains(ret.Tag.Value, "optional") {
	//		fmt.Printf("变长字段 on line %v:\n", GFset.Position(ret.Pos()))
	//		fmt.Printf("tag %+v found on line %v:\n", FmtColor(fixTypeInTag(ret.Tag.Value), testRed), GFset.Position(ret.Tag.Pos()))
	//		_ = printer.Fprint(os.Stdout, GFset, ret)
	//		_ = printer.Fprint(os.Stdout, GFset, ret.Tag)
	//		fmt.Printf("\n")
	//		ast.Print(GFset, ret)
	//
	//		//if arryDef, ok := ret.Type.(*ast.ArrayType); ok {
	//		//	ret.Type = f.arrayToOptional(arryDef)
	//		//}
	//	}
	//}

	return f
}

func (f *FindContext) arrayToOptional(arrayType *ast.ArrayType) ast.Expr {
	elt := arrayType.Elt
	switch elt.(type) {
	case *ast.Ident:
		ident := elt.(*ast.Ident)
		switch ident.Obj.Decl.(type) {
		case *ast.TypeSpec:
			switch varr := ident.Obj.Decl.(*ast.TypeSpec).Type; varr.(type) {
			case *ast.StructType:
				structType := varr.(*ast.StructType)
				structType.Fields.List[0].Pos()
				validBool := &ast.Ident{
					NamePos: structType.Fields.List[0].Pos(),
					Name:    "valid",
					Obj:     nil,
				}
				feild := &ast.Field{
					Names: []*ast.Ident{validBool},
					Type: &ast.Ident{
						NamePos: structType.Fields.List[0].Pos() + 10,
						Name:    "bool",
					},
				}
				structType.Fields.List = append(structType.Fields.List, feild)
			}
		}
		return elt
	default:
		log.Println("assert fail")
	}

	return elt
}

func Test_Parse(t *testing.T) {
	fset := token.NewFileSet()
	// path := `D:\goProject\src\code.huawei.com\5gcore\cp\domain\session\smc\exec\do\datas\`
	// path := `D:\goProject\src\code.huawei.com\5gcore\cp\domain\session\smc\exec\do\datas\elements\gtpeles\`
	path := `D:\gopath\src\gogen\astruct\`
	dir, _ := parser.ParseDir(fset, path, dirFilter, 0)
	GFset = fset

	for _, a := range dir {
		for fname := range a.Files {
			// Create the AST by parsing src.
			fsetInner := token.NewFileSet() // positions are relative to fsetInner
			f, err := parser.ParseFile(fsetInner, fname, nil, parser.ParseComments)
			if err != nil {
				log.Printf("ParseFile %s error:%v", fname, err)
				return
			} else {
				GFset = fsetInner
			}

			fix := &FindContext{}

			ast.Walk(fix, f)

			/*var buf bytes.Buffer
			printer.Fprint(&buf, fsetInner, f)

			genFile(fname, buf)*/
		}
	}

	// ast.Print(fset, dir)
}
func genFile(file string, buf bytes.Buffer) {
	// 替换原文件
	newFile, err := os.Create(file)
	defer newFile.Close()
	if err != nil {
		log.Printf("os.Create %s error:%v", file, err)
		return
	} else {
		source, _ := format.Source(buf.Bytes())
		newFile.Write(source)
	}

	// cmd := fmt.Sprintf("go fmt %s;goimports -w %s", file, file)
	// runCmd("/bin/sh", "-c", cmd)
}
