package typescript

import (
	"fmt"
	"go/ast"
	"go2ts/internal/typescript/tsast"
	"strconv"
)

func (p *printer) decls(name string, decls []ast.Decl, node tsast.Decls) {
	if len(decls) == 0 {
		return
	}

	p.printDebugf("%s: []ast.Decl\n", name)
	p.indentDebug++
	for i, decl := range decls {
		nodeDecl := &tsast.TsDecl{}
		p.decl(strconv.Itoa(i), decl, nodeDecl)
		node.AddDecl(nodeDecl)
	}
	p.indentDebug--
}

func (p *printer) decl(name string, decl ast.Decl, node tsast.Decl) {
	p.printDebugf("%s: ast.Decl\n", name)
	var declNode interface{}

	p.indentDebug++
	switch d := decl.(type) {
	case *ast.BadDecl:
		declNode = &tsast.TsBadDecl{}
		p.badDecl(name, d, declNode.(tsast.BadDecl))
	case *ast.GenDecl:
		declNode = &tsast.TsGenDecl{}
		p.genDecl(name, d, node.(tsast.GenDecl))
	case *ast.FuncDecl:
		declNode = &tsast.TsFuncDecl{}
		p.funcDecl(name, d, declNode.(tsast.FuncDecl))
	default:
		panic(fmt.Sprintf("ast.Decl invalid: %T", d))
	}
	p.indentDebug--

	node.SetDecl(declNode)
}

func (p *printer) badDecl(name string, badDecl *ast.BadDecl, node tsast.BadDecl) {
	p.printDebugf("%s: ast.BadDecl\n", name)
	p.indentDebug++
	p.indentDebug--
}

func (p *printer) genDecl(name string, genDecl *ast.GenDecl, node tsast.GenDecl) {
	p.printDebugf("%s: ast.GenDecl\n", name)

	p.indentDebug++
	p.commentGroup("Doc", genDecl.Doc, node)
	p.printDebugf("TokPos: %d\n", genDecl.TokPos)
	p.printDebugf("Tok: %d\n", genDecl.Tok)
	p.printDebugf("Lparen: %d\n", genDecl.Lparen)
	p.printDebugf("Rparen: %d\n", genDecl.Rparen)
	p.specs("Specs", genDecl.Specs, node)
	p.indentDebug--
}

func (p *printer) funcDecl(name string, funcDecl *ast.FuncDecl, node tsast.FuncDecl) {
	p.printDebugf("%s: *ast.FuncDecl\n", name)

	p.indentDebug++
	p.commentGroup("Doc", funcDecl.Doc, node)
	p.fieldList("Recv", funcDecl.Recv, node)
	p.ident("Name", funcDecl.Name, node)
	p.funcType("Type", funcDecl.Type, node)
	p.indentDebug--
}
