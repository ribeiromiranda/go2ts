package typescript

import (
	"go/ast"
	"go2ts/internal/typescript/tsast"
)

func (p *printer) funcType(name string, funcType *ast.FuncType, node tsast.FuncType) {
	p.printDebugf("%s: *ast.FuncType\n", name)

	p.indentDebug++
	p.fieldList("Params", funcType.Params, node)
	p.fieldList("Results", funcType.Results, node)
	p.indentDebug--
}

func (p *printer) interfaceType(name string, interfaceType *ast.InterfaceType, node tsast.InterfaceType) {
	p.printDebugf("%s: *ast.InterfaceType\n", name)

	p.indentDebug++
	p.printDebugf("Interface: %d\n", interfaceType.Interface)
	p.printDebugf("Incomplete: %t\n", interfaceType.Incomplete)
	p.fieldList("Methods", interfaceType.Methods, node)
	p.indentDebug--
}

func (p *printer) arrayType(name string, arrayType *ast.ArrayType, node tsast.ArrayType) {
	p.printDebugf("%s: *ast.ArrayType\n", name)

	p.indentDebug++
	p.printDebugf("Lbrack: %d\n", arrayType.Lbrack)
	p.expr("Len", arrayType.Len, node)
	p.expr("Elt", arrayType.Elt, node)
	p.indentDebug--
}

func (p *printer) structType(name string, structType *ast.StructType, node tsast.StructType) {
	p.printDebugf("%s: *ast.StructType\n", name)

	p.indentDebug++
	p.printDebugf("Incomplete: %t\n", structType.Incomplete)
	p.printDebugf("Struct: %d\n", structType.Struct)
	p.fieldList("Fields", structType.Fields, node)
	p.indentDebug--
}
