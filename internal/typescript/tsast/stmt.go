package tsast

type (
	Stmts      interface{ AddStmt(Stmt) }
	Stmt       interface{ SetStmt(interface{}) }
	SwitchStmt interface {
		Stmt
		Expr
		BlockStmt
	}
	BranchStmt interface {
		Identity
		Stmt
	}
	TypeSwitchStmt interface {
		Stmt
		BlockStmt
	}
	DeclStmt interface {
		Decl
	}
	IncDecStmt interface {
		Expr
	}
	ForStmt interface {
		Stmt
		Expr
		BlockStmt
	}
	RangeStmt interface {
		Expr
		BlockStmt
	}
	IfStmt interface {
		Stmt
		Expr
		BlockStmt
	}
	ReturnStmt interface {
		Exprs
	}
	ExprStmt interface {
		Expr
	}
	BlockStmt interface {
		Stmt
		Stmts
	}
	LabeledStmt interface {
		Identity
		Stmt
	}
	AssignStmt interface {
		Exprs
	}

	TsStmts          struct{}
	TsStmt           struct{}
	TsSwitchStmt     struct{}
	TsBranchStmt     struct{}
	TsTypeSwitchStmt struct{}
	TsDeclStmt       struct{}
	TsIncDecStmt     struct{}
	TsForStmt        struct{}
	TsRangeStmt      struct{}
	TsIfStmt         struct{}
	TsReturnStmt     struct{}
	TsExprStmt       struct{}
	TsBlockStmt      struct{}
	TsLabeledStmt    struct{}
	TsAssignStmt     struct{}
)

// Stmts
func (s *TsStmts) AddStmt(stmt Stmt)             {}
func (s *TsBlockStmt) AddStmt(stmt Stmt)         {}
func (s *TsSwitchStmt) AddStmt(stmt Stmt)        {}
func (s *TsCaseClause) AddStmt(stmt interface{}) {}

// Stmt
func (s *TsStmt) SetStmt(stmt interface{})           {}
func (s *TsSwitchStmt) SetStmt(stmt interface{})     {}
func (b *TsBranchStmt) SetStmt(stmt interface{})     {}
func (s *TsTypeSwitchStmt) SetStmt(stmt interface{}) {}
func (s *TsForStmt) SetStmt(stmt interface{})        {}
func (s *TsIfStmt) SetStmt(stmt interface{})         {}
func (s *TsRangeStmt) SetStmt(stmt interface{})      {}
func (s *TsBlockStmt) SetStmt(stmt interface{})      {}
