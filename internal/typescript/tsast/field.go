package tsast

type (
	FieldList interface {
		Fields
	}
	Fields interface{ AddField(Field) }
	Field  interface {
		CommentGroup
		Expr
		Identities
		BasicLit
	}

	TsFieldList struct{}
	TsFields    struct{}
	TsField     struct{}
)

// FieldList
func (f *TsFieldList) AddField(field Field)     {}
func (f *TsFuncDecl) AddField(field Field)      {}
func (f *TsFuncType) AddField(field Field)      {}
func (f *TsInterfaceType) AddField(field Field) {}
func (f *TsStructType) AddField(field Field)    {}

// Fields
func (f *TsFields) AddField(field Field) {}

// Field
