package typescript

import (
	"fmt"
	"go/ast"
	"go2ts/internal/typescript/tsast"
	"strconv"
)

func (p *printer) specs(name string, specs []ast.Spec, node tsast.Specs) {
	if len(specs) == 0 {
		return
	}
	p.printDebugf("%s: []ast.Spec\n", name)

	p.indentDebug++
	for i, spec := range specs {
		nodeSpec := &tsast.TsSpec{}
		p.spec(strconv.Itoa(i), spec, nodeSpec)
		node.AddSpec(nodeSpec)
	}
	p.indentDebug--
}

func (p *printer) spec(name string, spec ast.Spec, node tsast.Spec) {
	var nodeSpec interface{}

	p.printDebugf("%s: ast.Spec\n", name)
	p.indentDebug++
	switch s := spec.(type) {
	case *ast.ImportSpec:
		nodeSpec = &tsast.TsImportSpec{}
		p.importSpec(name, s, nodeSpec.(tsast.ImportSpec))
	case *ast.ValueSpec:
		nodeSpec = &tsast.TsValueSpec{}
		p.valueSpec(name, s, nodeSpec.(tsast.ValueSpec))
	case *ast.TypeSpec:
		nodeSpec = &tsast.TsTypeSpec{}
		p.typeSpec(name, s, nodeSpec.(tsast.TypeSpec))
	default:
		panic(fmt.Sprint("Spec not defined %T", s))
	}
	p.indentDebug--
}

func (p *printer) typeSpec(name string, typeSpec *ast.TypeSpec, node tsast.TypeSpec) {
	p.printDebugf("%s: *ast.TypeSpec\n", name)

	p.indentDebug++
	p.commentGroup("Doc", typeSpec.Doc, node)
	p.commentGroup("Comment", typeSpec.Comment, node)
	p.ident("Name", typeSpec.Name, node)
	p.expr("Type", typeSpec.Type, node)
	p.printDebugf("Assign: %d\n", typeSpec.Assign)
	p.indentDebug--
}

func (p *printer) importSpecs(name string, importSpec []*ast.ImportSpec, node tsast.ImportSpecs) {
	if len(importSpec) == 0 {
		return
	}

	p.printDebugf("%s: []*ast.ImportSpec\n", name)
	p.indentDebug++
	for i, imp := range importSpec {
		nodeImportSpec := &tsast.TsImportSpec{}
		p.importSpec(strconv.Itoa(i), imp, nodeImportSpec)
		node.AddImportSpec(nodeImportSpec)
	}
	p.indentDebug--
}

func (p *printer) importSpec(name string, importSpec *ast.ImportSpec, node tsast.ImportSpec) {
	p.printDebugf("%s: *ast.ImportSpec\n", name)

	p.indentDebug++
	p.commentGroup("Doc", importSpec.Doc, node)
	p.commentGroup("Comment", importSpec.Comment, node)
	p.ident("Name", importSpec.Name, node)
	p.basicLit("Path", importSpec.Path, node)
	p.indentDebug--

	// path = strings.Trim(path, "\"")
	// pathParts := strings.Split(path, "/")
	// if len(pathParts) > 1 {
	// 	path = fmt.Sprintf("import %s from \"%s\"", pathParts[len(pathParts)-1], strings.Join(pathParts[:len(pathParts)-1], "."))
	// } else {
	// 	path = fmt.Sprintf("import \"%s\"", path)
	// }

}

func (p *printer) valueSpec(name string, valueSpec *ast.ValueSpec, node tsast.ValueSpec) {
	p.printDebugf("%s: *ast.ValueSpec\n", name)
	p.indentDebug++
	p.commentGroup("Doc", valueSpec.Doc, node)
	p.commentGroup("Comment", valueSpec.Comment, node)
	p.idents("Names", valueSpec.Names, node)
	p.expr("Type", valueSpec.Type, node)
	p.exprs("Values", valueSpec.Values, node)
	p.indentDebug--
}
