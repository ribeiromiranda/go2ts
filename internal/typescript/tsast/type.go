package tsast

type (
	FuncType interface {
		FieldList
	}
	InterfaceType interface {
		FieldList
	}
	ArrayType interface {
		Expr
	}
	StructType interface {
		FieldList
	}

	TsFuncType      struct{}
	TsInterfaceType struct{}
	TsArrayType     struct{}
	TsStructType    struct{}
)
