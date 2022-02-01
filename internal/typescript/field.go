package typescript

import (
	"go/ast"
	"go2ts/internal/typescript/tsast"
	"strconv"
)

func (p *printer) fieldList(name string, fieldList *ast.FieldList, node tsast.FieldList) {
	if fieldList == nil || len(fieldList.List) == 0 {
		return
	}

	p.printDebugf("%s: ast.FieldList\n", name)
	p.indentDebug++
	p.fields("List", fieldList.List, node)
	p.indentDebug--
}

func (p *printer) fields(name string, fieldList []*ast.Field, node tsast.Fields) {
	if len(fieldList) == 0 {
		return
	}

	p.printDebugf("%s: []*ast.Field\n", name)
	p.indentDebug++
	for i, f := range fieldList {
		nodeField := &tsast.TsField{}
		p.field(strconv.Itoa(i), f, nodeField)
		node.AddField(nodeField)
	}
	p.indentDebug--
}

func (p *printer) field(name string, field *ast.Field, node tsast.Field) {
	p.printDebugf("%s: *ast.Field\n", name)

	p.indentDebug++
	p.commentGroup("Doc", field.Doc, node)
	p.commentGroup("Comment", field.Comment, node)
	p.idents("Names", field.Names, node)
	p.expr("Type", field.Type, node)
	p.basicLit("Tag", field.Tag, node)
	p.indentDebug--
}
