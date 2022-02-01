package typescript

import (
	"fmt"
	"go/ast"
	"go2ts/internal/typescript/tsast"
	"strconv"
)

func (p *printer) stmts(name string, stmts []ast.Stmt, node tsast.Stmts) {
	if len(stmts) == 0 {
		return
	}
	p.printDebugf("%s: []ast.Stmt\n", name)

	p.indentDebug++
	for i, stmt := range stmts {
		nodeStmt := &tsast.TsStmt{}
		p.stmt(strconv.Itoa(i), stmt, nodeStmt)
		node.AddStmt(nodeStmt)
	}
	p.indentDebug--
}

func (p *printer) stmt(name string, stmt ast.Stmt, node tsast.Stmt) {
	var stmtNode interface{}
	p.printDebugf("%s: ast.Stmt\n", name)

	p.indentDebug++
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		stmtNode = tsast.TsExprStmt{}
		p.exprStmt(name, s, stmtNode.(tsast.ExprStmt))
	case *ast.ReturnStmt:
		stmtNode = tsast.TsReturnStmt{}
		p.returnStmt(name, s, stmtNode.(tsast.ReturnStmt))
	case *ast.IfStmt:
		stmtNode = tsast.TsIfStmt{}
		p.ifStmt(name, s, stmtNode.(tsast.IfStmt))
	case *ast.RangeStmt:
		stmtNode = tsast.TsRangeStmt{}
		p.rangeStmt(name, s, stmtNode.(tsast.RangeStmt))
	case *ast.AssignStmt:
		stmtNode = tsast.TsAssignStmt{}
		p.assignStmt(name, s, stmtNode.(tsast.AssignStmt))
	case *ast.ForStmt:
		stmtNode = tsast.TsForStmt{}
		p.forStmt(name, s, stmtNode.(tsast.ForStmt))
	case *ast.BlockStmt:
		stmtNode = tsast.TsBlockStmt{}
		p.blockStmt(name, s, stmtNode.(tsast.BlockStmt))
	case *ast.IncDecStmt:
		stmtNode = tsast.TsIncDecStmt{}
		p.incDecStmt(name, s, stmtNode.(tsast.IncDecStmt))
	case *ast.DeclStmt:
		stmtNode = tsast.TsDeclStmt{}
		p.declStmt(name, s, stmtNode.(tsast.DeclStmt))
	case *ast.TypeSwitchStmt:
		stmtNode = tsast.TsTypeSwitchStmt{}
		p.typeSwitchStmt(name, s, stmtNode.(tsast.TypeSwitchStmt))
	case *ast.CaseClause:
		stmtNode = tsast.TsCaseClause{}
		p.caseClause(name, s, stmtNode.(tsast.CaseClause))
	case *ast.BranchStmt:
		stmtNode = tsast.TsBranchStmt{}
		p.branchStmt(name, s, stmtNode.(tsast.BranchStmt))
	case *ast.LabeledStmt:
		stmtNode = tsast.TsLabeledStmt{}
		p.labeledStmt(name, s, stmtNode.(tsast.LabeledStmt))
	case *ast.SwitchStmt:
		stmtNode = tsast.TsSwitchStmt{}
		p.switchStmt(name, s, stmtNode.(tsast.SwitchStmt))
	case nil:
		p.printDebugf("nil\n")
	default:
		panic(fmt.Sprintf("Stmts not defined: %T", s))
	}
	p.indentDebug--

	node.SetStmt(stmtNode)
}

func (p *printer) switchStmt(name string, stmt *ast.SwitchStmt, node tsast.SwitchStmt) {
	p.printDebugf("%s: *ast.SwitchStmt\n", name)

	p.indentDebug++
	p.printDebugf("Switch: %d\n", stmt.Switch)
	p.stmt("Init", stmt.Init, node)
	p.expr("Tag", stmt.Tag, node)
	p.blockStmt("Body", stmt.Body, node)
	p.indentDebug--
}

func (p *printer) branchStmt(name string, stmt *ast.BranchStmt, node tsast.BranchStmt) {
	p.printDebugf("%s: *ast.BranchStmt\n", name)

	p.indentDebug++
	p.printDebugf("TokPos: %d\n", stmt.TokPos)
	p.printDebugf("Tok: %d\n", stmt.Tok)
	p.ident("Label", stmt.Label, node)
	p.indentDebug--
}

func (p *printer) typeSwitchStmt(name string, stmt *ast.TypeSwitchStmt, node tsast.TypeSwitchStmt) {
	p.printDebugf("%s: *ast.TypeSwitchStmt\n", name)

	p.indentDebug++
	p.printDebugf("Switch: %d\n", stmt.Switch)
	p.stmt("Init", stmt.Init, node)
	p.stmt("Assign", stmt.Assign, node)
	p.blockStmt("Body", stmt.Body, node)
	p.indentDebug--
}

func (p *printer) declStmt(name string, stmt *ast.DeclStmt, node tsast.DeclStmt) {
	p.printDebugf("%s: *ast.DeclStmt\n", name)

	p.indentDebug++
	p.decl("Decl", stmt.Decl, node)
	p.indentDebug--
}

func (p *printer) incDecStmt(name string, stmt *ast.IncDecStmt, node tsast.IncDecStmt) {
	p.printDebugf("%s: *ast.DeclStmt\n", name)

	p.indentDebug++
	p.printDebugf("TokPos: %d\n", stmt.TokPos)
	p.printDebugf("Tok: %d\n", stmt.Tok)
	p.expr("X", stmt.X, node)

	p.indentDebug--

}

func (p *printer) forStmt(name string, stmt *ast.ForStmt, node tsast.ForStmt) {
	p.printDebugf("%s: *ast.DeclStmt\n", name)

	p.indentDebug++
	p.printDebugf("For: %d\n", stmt.For)
	p.stmt("Init", stmt.Init, node)
	p.expr("Cond", stmt.Cond, node)
	p.stmt("Post", stmt.Post, node)
	p.blockStmt("Body", stmt.Body, node)
	p.indentDebug--
}

func (p *printer) rangeStmt(name string, rangeStmt *ast.RangeStmt, node tsast.RangeStmt) {
	p.printDebugf("%s: *ast.DeclStmt\n", name)
	p.indentDebug++

	p.printDebugf("For: %d\n", rangeStmt.For)
	p.printDebugf("TokPos: %d\n", rangeStmt.TokPos)
	p.printDebugf("Tok: %d\n", rangeStmt.Tok)
	p.expr("Key", rangeStmt.Key, node)
	p.expr("Value", rangeStmt.Value, node)
	p.expr("X", rangeStmt.X, node)
	p.blockStmt("Body", rangeStmt.Body, node)

	p.indentDebug--
}

func (p *printer) ifStmt(name string, ifStmt *ast.IfStmt, node tsast.IfStmt) {
	p.printDebugf("%s: *ast.DeclStmt\n", name)

	p.indentDebug++
	p.printDebugf("If: %d\n", ifStmt.If)
	p.stmt("Init", ifStmt.Init, node)
	p.expr("Cond", ifStmt.Cond, node)
	p.blockStmt("Body", ifStmt.Body, node)
	p.stmt("Else", ifStmt.Else, node)
	p.indentDebug--

}

func (p *printer) returnStmt(name string, returnStmt *ast.ReturnStmt, node tsast.ReturnStmt) {
	p.printDebugf("%s: *ast.DeclStmt\n", name)

	p.indentDebug++
	p.printDebugf("Return: %T\n", returnStmt.Return)
	p.exprs("Results", returnStmt.Results, node)
	p.indentDebug--
}

func (p *printer) exprStmt(name string, exprStmt *ast.ExprStmt, node tsast.ExprStmt) {
	p.printDebugf("%s: *ast.DeclStmt\n", name)

	p.indentDebug++
	p.expr("X", exprStmt.X, node)
	p.indentDebug--
}

func (p *printer) blockStmt(name string, blockStmt *ast.BlockStmt, node tsast.BlockStmt) {
	if len(blockStmt.List) == 0 {
		return
	}
	p.printDebugf("%s: *ast.BlockStmt\n", name)

	p.indentDebug++
	p.stmts("List", blockStmt.List, node)
	p.indentDebug--
}

func (p *printer) labeledStmt(name string, labeledStmt *ast.LabeledStmt, node tsast.LabeledStmt) {
	p.printDebugf("%s: *ast.LabeledStmt\n", name)

	p.indentDebug++
	p.printDebugf("Colon: %d", labeledStmt.Colon)
	p.ident("Label", labeledStmt.Label, node)
	p.stmt("Stmt", labeledStmt.Stmt, node)
	p.indentDebug--
}

func (p *printer) assignStmt(name string, assignStmt *ast.AssignStmt, node tsast.AssignStmt) {
	p.printDebugf("%s: *ast.AssignStmt", name)

	p.indentDebug++
	p.printDebugf("Tok: %d", assignStmt.Tok)
	p.printDebugf("TokPos: %d", assignStmt.TokPos)
	p.exprs("Lhs", assignStmt.Lhs, node)
	p.exprs("Rhs", assignStmt.Rhs, node)
	p.indentDebug--
}
