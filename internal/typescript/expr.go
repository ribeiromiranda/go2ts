package typescript

import (
	"fmt"
	"go/ast"
	"go2ts/internal/typescript/tsast"
	"strconv"
)

func (p *printer) exprs(name string, expr []ast.Expr, node tsast.Exprs) {
	if len(expr) == 0 {
		return
	}
	p.printDebugf("%s: []ast.Expr\n", name)

	p.indentDebug++
	for i, e := range expr {
		nodeExpr := &tsast.TsExpr{}
		p.expr(strconv.Itoa(i), e, nodeExpr)
		node.AddExpr(nodeExpr)
	}
	p.indentDebug--
}

func (p *printer) expr(name string, expr ast.Expr, node tsast.Expr) {
	p.printDebugf("%s: ast.Expr\n", name)

	p.indentDebug++
	var exprNode interface{}
	switch e := expr.(type) {
	case *ast.Ident:
		p.ident(name, e, node)
	case *ast.Ellipsis:
		exprNode = &tsast.TsEllipsis{}
		p.ellipsis(name, e, exprNode.(tsast.Ellipsis))
	case *ast.CompositeLit:
		exprNode = &tsast.TsCompositLit{}
		p.compositeLit(name, e, exprNode.(tsast.CompositeLit))
	case *ast.BasicLit:
		exprNode = &tsast.TsBasicLit{}
		p.basicLit(name, e, exprNode.(tsast.BasicLit))
	case *ast.InterfaceType:
		exprNode = &tsast.TsInterfaceType{}
		p.interfaceType(name, e, exprNode.(tsast.InterfaceType))
	case *ast.StructType:
		exprNode = &tsast.TsStructType{}
		p.structType(name, e, exprNode.(tsast.StructType))
	case *ast.ArrayType:
		exprNode = &tsast.TsArrayType{}
		p.arrayType(name, e, exprNode.(tsast.ArrayType))
	case *ast.FuncType:
		exprNode = &tsast.TsFuncType{}
		p.funcType(name, e, exprNode.(tsast.FuncType))
	case *ast.CallExpr:
		exprNode = &tsast.TsCallExpr{}
		p.callExpr(name, e, exprNode.(tsast.CallExpr))
	case *ast.SelectorExpr:
		exprNode = &tsast.TsSelectorExpr{}
		p.selectExpr(name, e, exprNode.(tsast.SelectorExpr))
	case *ast.StarExpr:
		exprNode = &tsast.TsStarExpr{}
		p.starExpr(name, e, exprNode.(tsast.StarExpr))
	case *ast.UnaryExpr:
		exprNode = &tsast.TsUnaryExpr{}
		p.unaryExpr(name, e, exprNode.(tsast.UnaryExpr))
	case *ast.BinaryExpr:
		exprNode = &tsast.TsBinaryExpr{}
		p.binaryExpr(name, e, exprNode.(tsast.BinaryExpr))
	case *ast.KeyValueExpr:
		exprNode = &tsast.TsKeyValueExpr{}
		p.keyValueExpr(name, e, exprNode.(tsast.KeyValueExpr))
	case *ast.SliceExpr:
		exprNode = &tsast.TsSliceExpr{}
		p.sliceExpr(name, e, exprNode.(tsast.SliceExpr))
	case *ast.IndexExpr:
		exprNode = &tsast.TsIndexExpr{}
		p.indexExpr(name, e, exprNode.(tsast.IndexExpr))
	case *ast.ParenExpr:
		exprNode = &tsast.TsParenExpr{}
		p.parenExpr(name, e, exprNode.(tsast.ParenExpr))
	case *ast.TypeAssertExpr:
		exprNode = &tsast.TsTypeAssertExpr{}
		p.typeAssertExpr(name, e, exprNode.(tsast.TypeAssertExpr))
	case nil:
		p.printDebugf("nil\n")
	default:
		panic(fmt.Sprintf("Expr not defined: %T", e))
	}
	p.indentDebug--

	node.SetExpr(exprNode)
}

func (p *printer) typeAssertExpr(name string, expr *ast.TypeAssertExpr, node tsast.TypeAssertExpr) {
	p.printDebugf("%s: *ast.TypeAssertExpr\n", name)

	p.indentDebug++
	p.printDebugf("Lparen: %d\n", expr.Lparen)
	p.printDebugf("Rparen: %d\n", expr.Rparen)
	p.expr("Type", expr.Type, node)
	p.expr("X", expr.X, node)
	p.indentDebug--
}

func (p *printer) parenExpr(name string, expr *ast.ParenExpr, node tsast.ParenExpr) {
	p.printDebugf("%s: *ast.ParenExpr\n", name)
	p.indentDebug++

	p.printDebugf("Lparen: %d\n", expr.Lparen)
	p.printDebugf("Rparen: %d\n", expr.Rparen)
	p.expr("X", expr.X, node)

	p.indentDebug--
}

func (p *printer) indexExpr(name string, expr *ast.IndexExpr, node tsast.IndexExpr) {
	p.printDebugf("%s: *ast.IndexExpr\n", name)

	p.indentDebug++
	p.printDebugf("Lbrack: %d\n", expr.Lbrack, node)
	p.printDebugf("Rbrack: %d\n", expr.Rbrack, node)
	p.expr("X", expr.X, node)
	p.expr("Index", expr.Index, node)
	p.indentDebug--
}

func (p *printer) sliceExpr(name string, sliceExpr *ast.SliceExpr, node tsast.SliceExpr) {
	p.printDebugf("%s: *ast.SliceExpr\n", name)

	p.indentDebug++
	p.printDebugf("Lbrack: %d\n", sliceExpr.Lbrack)
	p.printDebugf("Rbrack: %d\n", sliceExpr.Rbrack)
	p.printDebugf("Slice3: %T\n", sliceExpr.Slice3)
	p.expr("X", sliceExpr.X, node)
	p.expr("Low", sliceExpr.Low, node)
	p.expr("High", sliceExpr.High, node)
	p.expr("Max", sliceExpr.Max, node)
	p.indentDebug--
}

func (p *printer) keyValueExpr(name string, keyValueExpr *ast.KeyValueExpr, node tsast.KeyValueExpr) {
	p.printDebugf("%s: *ast.KeyValueExpr\n", name)

	p.indentDebug++
	p.printDebugf("Colon: %d\n", keyValueExpr.Colon)
	p.expr("Key", keyValueExpr.Key, node)
	p.expr("Value", keyValueExpr.Value, node)
	p.indentDebug--
}

func (p *printer) binaryExpr(name string, binaryExpr *ast.BinaryExpr, node tsast.BinaryExpr) {
	p.printDebugf("%s: *ast.BinaryExpr\n", name)

	p.indentDebug++
	p.printDebugf("OpPos: %T\n", binaryExpr.OpPos)
	p.printDebugf("Op: %T\n", binaryExpr.Op)
	p.expr("X", binaryExpr.X, node)
	p.expr("Y", binaryExpr.Y, node)
	p.indentDebug--
}

func (p *printer) unaryExpr(name string, unaryExpr *ast.UnaryExpr, node tsast.UnaryExpr) {
	p.printDebugf("%s: *ast.UnaryExpr\n", name)

	p.indentDebug++
	p.printDebugf("OpPos: %T\n", unaryExpr.OpPos)
	p.printDebugf("Op: %T\n", unaryExpr.Op)
	p.expr("X", unaryExpr.X, node)
	p.indentDebug--
}

func (p *printer) starExpr(name string, starExpr *ast.StarExpr, node tsast.StarExpr) {
	p.printDebugf("%s: *ast.StarExpr\n", name)

	p.indentDebug++
	p.printDebugf("Star: %d\n", starExpr.Star)
	p.expr("X", starExpr.X, node)
	p.indentDebug--
}

func (p *printer) callExpr(name string, expr *ast.CallExpr, node tsast.CallExpr) {
	p.printDebugf("%s: *ast.CallExpr\n", name)

	p.indentDebug++
	p.expr("Fun", expr.Fun, node)
	p.exprs("Args", expr.Args, node)
	p.printDebugf("Lparen: %d\n", expr.Lparen)
	p.printDebugf("Rparen: %d\n", expr.Rparen)
	p.printDebugf("Ellipsis: %d\n", expr.Ellipsis)
	p.indentDebug--
}

func (p *printer) selectExpr(name string, expr *ast.SelectorExpr, node tsast.SelectorExpr) {
	p.printDebugf("%s: *ast.SelectorExpr\n", name)

	p.indentDebug++
	p.ident("Sel", expr.Sel, node)
	p.expr("X", expr.X, node)
	p.indentDebug--
}
