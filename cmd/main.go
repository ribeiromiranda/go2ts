package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go2ts/internal/typescript"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "./tests/interface.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// print.Print(fset, file)
	typescript.Print(fset, file)

	return
	ast.Inspect(file, func(n ast.Node) bool {
		fmt.Printf("%T\n", n)
		// Find Function Call Statements
		funcCall, ok := n.(*ast.CallExpr)
		if ok {
			fmt.Println(funcCall.Fun)
		}
		return true
	})
}
