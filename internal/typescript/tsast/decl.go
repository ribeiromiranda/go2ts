package tsast

type (
	Decls interface {
		AddDecl(Decl)
	}
	Decl interface {
		SetDecl(interface{})
	}
	BadDecl interface{}
	GenDecl interface {
		CommentGroup
		Specs
	}
	FuncDecl interface {
		CommentGroup
		Identity
		FieldList
	}
	AddDecl interface{}

	TsDecls    struct{}
	TsDecl     struct{}
	TsBadDecl  struct{}
	TsGenDecl  struct{}
	TsFuncDecl struct{}
	TsAddDecl  struct{}
)

// Decls
func (d *TsDecls) AddDecl(decl Decl) {}

// Decl
func (d *TsDecl) SetDecl(decl interface{})     {}
func (d *TsDeclStmt) SetDecl(decl interface{}) {}
