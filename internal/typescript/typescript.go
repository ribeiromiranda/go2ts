// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains printing support for ASTs.

package typescript

import (
	"fmt"
	"go/ast"
	"go/token"
	"go2ts/internal/typescript/tsast"
	"os"
	"reflect"
	"strings"
)

// IsExported reports whether name starts with an upper-case letter.
//
func IsExported(name string) bool { return token.IsExported(name) }

// A FieldFilter may be provided to Fprint to control the output.
type FieldFilter func(name string, value reflect.Value) bool

// NotNilFilter returns true for field values that are not nil;
// it returns false otherwise.
func NotNilFilter(_ string, v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return !v.IsNil()
	}
	return true
}

type printer struct {
	fileOut     *os.File
	fset        *token.FileSet
	filter      FieldFilter
	ptrmap      map[interface{}]int // *T -> line number
	indent      int                 // current indentation level
	indentDebug int
	line        int // current line number
}

// localError wraps locally caught errors so we can distinguish
// them from genuine panics which we don't want to return as errors.
type localError struct {
	err error
}

// Fprint prints the (sub-)tree starting at AST node x to w.
// If fset != nil, position information is interpreted relative
// to that file set. Otherwise positions are printed as integer
// values (file set specific offsets).
//
// A non-nil FieldFilter f may be provided to control the output:
// struct fields for which f(fieldname, fieldvalue) is true are
// printed; all others are filtered from the output. Unexported
// struct fields are never printed.
func Fprint(fset *token.FileSet, x *ast.File, f FieldFilter) error {
	return fprint(fset, x, f)
}

func fprint(fset *token.FileSet, x *ast.File, f FieldFilter) (err error) {
	fOut, err := os.Create("/tmp/golang")
	if err != nil {
		return err
	}

	// setup printer
	p := printer{
		fileOut: fOut,
		fset:    fset,
		filter:  f,
		ptrmap:  make(map[interface{}]int),
	}

	// install error handler
	defer func() {
		if e := recover(); e != nil {
			err = e.(localError).err // re-panics if it's not a localError
		}
	}()

	// print x
	if x == nil {
		fmt.Printf("nil\n")
		return
	}
	p.file(x)
	fmt.Printf("\n")

	return
}

// Print prints x to standard output, skipping nil fields.
// Print(fset, x) is the same as Fprint(os.Stdout, fset, x, NotNilFilter).
func Print(fset *token.FileSet, x *ast.File) error {
	return Fprint(fset, x, NotNilFilter)
}

func (p *printer) printf(text string, a ...interface{}) {
	ident := strings.Repeat(" ", p.indent)

	p.fileOut.WriteString(fmt.Sprintf("%s", ident))
	p.fileOut.WriteString(fmt.Sprintf(text, a...))
}

func (p *printer) printDebugf(text string, a ...interface{}) {
	ident := strings.Repeat(".", (p.indentDebug-1)*2)
	fmt.Printf("%s", ident)
	fmt.Printf(text, a...)
}

// Implementation note: Print is written for AST nodes but could be
// used to print arbitrary data structures; such a version should
// probably be in a different package.
//
// Note: This code detects (some) cycles created via pointers but
// not cycles that are created via slices or maps containing the
// same slice or map. Code for general data structures probably
// should catch those as well.
func (p *printer) file(f *ast.File) {
	sourceFile := &tsast.TsFile{}

	p.indentDebug++
	p.printDebugf("*ast.File\n")
	p.indentDebug++

	p.commentGroup("Doc", f.Doc, sourceFile)
	p.commentGroups("Comments", f.Comments, sourceFile)
	p.ident("Name", f.Name, sourceFile)
	p.printDebugf("Package: %d\n", f.Package)
	p.importSpecs("Imports", f.Imports, sourceFile)
	p.scope("Scope", f.Scope, sourceFile)

	// p.printDebugf("Decls: \n")
	// decls := p.decls(f.Decls)

	// p.printDebugf("Unresolved: \n")
	// unresolved := p.scope(f.Unresolved)
	p.indentDebug--
	p.indentDebug--
}

func (p *printer) scope(name string, scope *ast.Scope, node tsast.Scope) {
	p.printDebugf("%s: *ast.Scope\n", name)
	p.indentDebug++
	// if scope.Outer != nil {
	// 	p.scope("Outer", scope.Outer, node)
	// }

	p.objects("Objects", scope.Objects, node)
	p.indentDebug--
}

func (p *printer) objects(name string, objects map[string]*ast.Object, node tsast.Objects) {
	if len(objects) == 0 {
		return
	}

	p.printDebugf("%s: map[string]*ast.Object\n", name)
	p.indentDebug++
	for i, obj := range objects {
		nodeObject := &tsast.TsObject{}
		p.object(i, obj, nodeObject)
		node.AddObject(nodeObject)
	}
	p.indentDebug--
}

func (p *printer) object(name string, object *ast.Object, node tsast.Object) {
	if object == nil {
		return
	}

	p.printDebugf("%s: *ast.Object\n", name)

	p.indentDebug++
	p.printDebugf("Name: %s\n", object.Name)
	p.printDebugf("Data: %T\n", object.Data)
	p.printDebugf("Type: %T\n", object.Type)
	kindStr := ""
	switch object.Kind {
	case ast.Bad: // for error handling
		kindStr = ""
	case ast.Pkg: // package
		kindStr = "package"
	case ast.Con: // constant
		kindStr = "constat"
	case ast.Typ: // type
		kindStr = "type"
	case ast.Var: // variable
		kindStr = "variable"
	case ast.Fun: // function or method
		kindStr = object.Name
	case ast.Lbl: // label
		kindStr = "label"
	default:
		panic(fmt.Sprintf("Kind invalid: %T", object.Kind))
	}
	p.printDebugf("Kind: %s\n", kindStr)

	if object.Decl != nil {
		var declNode interface{}
		switch d := object.Decl.(type) {
		case *ast.Field:
			declNode = &tsast.TsField{}
			p.field(name, d, declNode.(tsast.Field))
		case *ast.ImportSpec:
			declNode = &tsast.TsImportSpec{}
			p.importSpec(name, d, declNode.(tsast.ImportSpec))
		case *ast.ValueSpec:
			declNode = &tsast.TsValueSpec{}
			p.valueSpec(name, d, declNode.(tsast.ValueSpec))
		case *ast.TypeSpec:
			declNode = &tsast.TsTypeSpec{}
			p.typeSpec(name, d, declNode.(tsast.TypeSpec))
		case *ast.FuncDecl:
			declNode = &tsast.TsFuncDecl{}
			p.funcDecl(name, d, declNode.(tsast.FuncDecl))
		case *ast.LabeledStmt:
			declNode = &tsast.TsLabeledStmt{}
			p.labeledStmt(name, d, declNode.(tsast.LabeledStmt))
		case *ast.AssignStmt:
			declNode = &tsast.TsAssignStmt{}
			p.assignStmt(name, d, declNode.(tsast.AssignStmt))
		case *ast.Scope:
			declNode = &tsast.TsScope{}
			p.scope("object", d, declNode.(tsast.Scope))
		default:
			panic(fmt.Sprintf("Type not defined: %T", object.Decl))
		}
	}

	p.indentDebug--
}

func (p *printer) caseClause(name string, caseClause *ast.CaseClause, node tsast.CaseClause) {
	p.printDebugf("%s: *ast.CaseClause\n", name)

	p.indentDebug++
	p.printDebugf("Case: %d\n", caseClause.Case)
	p.printDebugf("Colon: %d\n", caseClause.Colon)
	p.exprs("Init", caseClause.List, node)
	p.stmts("Body", caseClause.Body, node)
	p.indentDebug--
}

func (p *printer) ellipsis(name string, ellipsis *ast.Ellipsis, node tsast.Ellipsis) {
	p.printDebugf("%s: \n", name)

	p.indentDebug++
	p.printDebugf("Ellipsis: %T\n", ellipsis.Ellipsis)
	p.expr("Elt", ellipsis.Elt, node)
	p.indentDebug--
}
