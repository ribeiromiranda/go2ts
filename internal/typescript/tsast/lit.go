package tsast

type (
	BasicLit interface {
		SetValue(string)
	}
	CompositeLit interface {
		Exprs
	}

	TsBasicLit    struct{}
	TsCompositLit struct{}
)

// BasicLit
func (b *TsBasicLit) SetValue(value string)   {}
func (b *TsObject) SetValue(value string)     {}
func (s *TsImportSpec) SetValue(value string) {}
func (s *TsField) SetValue(value string)      {}
