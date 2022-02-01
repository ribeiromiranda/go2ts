package tsast

type (
	Specs interface {
		AddSpec(Spec)
	}
	Spec interface {
		SetSpec()
	}
	TypeSpec interface {
		CommentGroup
		Expr
	}
	ImportSpecs interface {
		AddImportSpec(ImportSpec)
	}
	ImportSpec interface {
		Identity
		CommentGroup
		BasicLit
		SetImportSec()
	}
	ValueSpec interface {
		CommentGroup
		Identities
		Exprs
		Expr
	}

	TsSpecs       struct{}
	TsSpec        struct{}
	TsTypeSpec    struct{}
	TsImportSpecs struct{}
	TsImportSpec  struct{}
	TsValueSpec   struct{}
)

// Specs
func (s *TsSpecs) AddSpec(spec Spec)   {}
func (g *TsGenDecl) AddSpec(spec Spec) {}

// Spec
func (s *TsSpec) SetSpec() {}

// ImportSpecs
func (i *TsImportSpecs) AddImportSpec(importSpec ImportSpec) {}
func (f *TsFile) AddImportSpec(importSpec ImportSpec)        {}

// ImportSpec
func (i *TsImportSpec) SetImportSec() {}
func (o *TsObject) SetImportSec()     {}
