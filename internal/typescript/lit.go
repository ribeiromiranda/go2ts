package typescript

import (
	"go/ast"
	"go2ts/internal/typescript/tsast"
)

func (p *printer) basicLit(name string, basicLit *ast.BasicLit, node tsast.BasicLit) {
	if basicLit == nil {
		return
	}

	p.printDebugf("%s: ast.BasicLit\n", name)

	p.indentDebug++
	value := basicLit.Value
	p.printDebugf("Value: %s\n", value)
	p.indentDebug--

	node.SetValue(value)
}

func (p *printer) compositeLit(name string, compositeLit *ast.CompositeLit, node tsast.CompositeLit) {
	p.printDebugf("%s: *ast.CompositeLit\n", compositeLit)

	p.indentDebug++
	p.printDebugf("Lbrace: %T\n", compositeLit.Lbrace)
	p.printDebugf("Rbrace: %T\n", compositeLit.Rbrace)
	p.printDebugf("Incomplete: %T\n", compositeLit.Incomplete)
	p.exprs("Elts", compositeLit.Elts, node)
	p.indentDebug--
}
