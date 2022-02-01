package tsast

type (
	Exprs interface {
		AddExpr(Expr)
	}
	Expr interface {
		Identity
		SetExpr(interface{})
	}
	TypeAssertExpr interface {
		Expr
	}
	ParenExpr interface {
		Expr
	}
	IndexExpr interface {
		Expr
	}
	SliceExpr interface {
		Expr
	}
	KeyValueExpr interface {
		Expr
	}
	BinaryExpr interface {
		Expr
	}
	UnaryExpr interface {
		Expr
	}
	StarExpr interface {
		Expr
	}
	CallExpr interface {
		Exprs
		Expr
	}
	SelectorExpr interface {
		Expr
		Identity
	}

	TsExprs          struct{}
	TsExpr           struct{}
	TsTypeAssertExpr struct{}
	TsParenExpr      struct{}
	TsIndexExpr      struct{}
	TsSliceExpr      struct{}
	TsKeyValueExpr   struct{}
	TsBinaryExpr     struct{}
	TsUnaryExpr      struct{}
	TsStarExpr       struct{}
	TsCallExpr       struct{}
	TsSelectorExpr   struct{}
)

// Exprs
func (e *TsExprs) AddExpr(expr Expr)              {}
func (e *TsCallExpr) AddExpr(expr interface{})    {}
func (e *TsCompositLit) AddExpr(expr interface{}) {}
func (e *TsValueSpec) AddExpr(expr interface{})   {}
func (e *TsReturnStmt) AddExpr(expr interface{})  {}
func (e *TsAssignStmt) AddExpr(expr interface{})  {}
func (e *TsCaseClause) AddExpr(expr interface{})  {}

// Expr
func (e *TsExpr) SetExpr(expr interface{})           {}
func (e *TsTypeAssertExpr) SetExpr(expr interface{}) {}
func (e *TsParenExpr) SetExpr(expr interface{})      {}
func (e *TsIndexExpr) SetExpr(expr interface{})      {}
func (e *TsSliceExpr) SetExpr(expr interface{})      {}
func (e *TsKeyValueExpr) SetExpr(expr interface{})   {}
func (e *TsBinaryExpr) SetExpr(expr interface{})     {}
func (e *TsUnaryExpr) SetExpr(expr interface{})      {}
func (e *TsStarExpr) SetExpr(expr interface{})       {}
func (e *TsCallExpr) SetExpr(expr interface{})       {}
func (e *TsSelectorExpr) SetExpr(expr interface{})   {}
func (e *TsField) SetExpr(expr interface{})          {}
func (e *TsValueSpec) SetExpr(expr interface{})      {}
func (e *TsSwitchStmt) SetExpr(expr interface{})     {}
func (e *TsIncDecStmt) SetExpr(expr interface{})     {}
func (e *TsForStmt) SetExpr(expr interface{})        {}
func (e *TsRangeStmt) SetExpr(expr interface{})      {}
func (e *TsIfStmt) SetExpr(expr interface{})         {}
func (e *TsExprStmt) SetExpr(expr interface{})       {}
func (e *TsArrayType) SetExpr(expr interface{})      {}
func (e *TsObject) SetExpr(expr interface{})         {}
func (e *TsEllipsis) SetExpr(expr interface{})       {}
func (e *TsTypeSpec) SetExpr(expr interface{})       {}
