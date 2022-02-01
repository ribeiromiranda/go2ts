package tsast

type (
	Scope interface {
		Objects
		AddObjects(Objects)
	}
	Objects interface{ AddObject(Object) }
	Object  interface {
		Field
	}

	Identities interface{ AddIdentity(Identity) }
	Identity   interface{ SetIdentity(string) }
	Import     interface{ AddImport(string) }

	CaseClause interface {
		Exprs
		Stmts
	}
	Ellipsis interface {
		Expr
	}

	TsScope      struct{}
	TsFile       struct{}
	TsObjects    struct{}
	TsObject     struct{}
	TsIdentities struct{}
	TsIdentity   struct{}
	TsImport     struct{}
	TsCaseClause struct{}
	TsEllipsis   struct{}
)

// Scope
func (o *TsScope) AddObjects(objects Objects) {}
func (o *TsFile) AddObjects(objects Objects)  {}

// Identities
func (o *TsObject) AddIdentity(name Identity) {}
func (f *TsField) AddIdentity(name Identity)  {}

// Identity
func (f *TsIdentity) SetIdentity(name string)         {}
func (f *TsFile) SetIdentity(name string)             {}
func (i *TsImportSpec) SetIdentity(name string)       {}
func (f *TsFuncDecl) SetIdentity(identity string)     {}
func (e *TsExpr) SetIdentity(identity string)         {}
func (s *TsSelectorExpr) SetIdentity(identity string) {}
func (f *TsField) SetIdentity(identity string)        {}
func (f *TsSwitchStmt) SetIdentity(identity string)   {}
func (b *TsBranchStmt) SetIdentity(identity string)   {}
func (b *TsIncDecStmt) SetIdentity(identity string)   {}
func (b *TsIfStmt) SetIdentity(identity string)       {}
func (b *TsLabeledStmt) SetIdentity(identity string)  {}
func (b *TsObject) SetIdentity(identity string)       {}
func (b *TsTypeSpec) SetIdentity(identity string)     {}

// Objects
func (o *TsFile) AddObject(object Object)  {}
func (o *TsScope) AddObject(object Object) {}
