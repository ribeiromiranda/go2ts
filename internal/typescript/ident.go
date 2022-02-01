package typescript

import (
	"go/ast"
	"go2ts/internal/typescript/tsast"
	"strconv"
)

func (p *printer) idents(name string, idents []*ast.Ident, node tsast.Identities) {
	if len(idents) == 0 {
		return
	}

	p.printDebugf("%s: []*ast.Ident\n", name)
	p.indentDebug++
	for i, ident := range idents {
		nodeIdent := &tsast.TsIdentity{}
		p.ident(strconv.Itoa(i), ident, nodeIdent)
		node.AddIdentity(nodeIdent)
	}
	p.indentDebug--
}

func (p *printer) ident(name string, ident *ast.Ident, node tsast.Identity) {
	if ident == nil {
		return
	}

	p.printDebugf("%s: *ast.Ident\n", name)
	p.indentDebug++
	nameIdent := ident.Name
	node.SetIdentity(nameIdent)
	p.printDebugf("Name: %s\n", nameIdent)
	p.printDebugf("NamePos: %d\n", ident.NamePos)
	// p.object("Obj", ident.Obj)
	p.indentDebug--
}
