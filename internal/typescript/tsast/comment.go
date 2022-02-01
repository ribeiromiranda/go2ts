package tsast

type (
	CommentGroups interface {
		AddCommentGroup(CommentGroup)
	}

	CommentGroup interface {
		AddComment(Comment)
	}

	Comment interface {
		SetComment(string)
	}

	TsCommentGroups struct{}

	TsCommentGroup struct{}

	TsComment struct{}
)

// CommentGroups
func (cg *TsCommentGroups) AddCommentGroup(commentGroup CommentGroup) {}
func (f *TsFile) AddCommentGroup(commentGroup CommentGroup)           {}

// CommentGroup
func (c *TsCommentGroup) AddComment(comment Comment) {}
func (f *TsFile) AddComment(Comment)                 {}
func (i *TsImportSpec) AddComment(comment Comment)   {}
func (g *TsGenDecl) AddComment(comment Comment)      {}
func (d *TsFuncDecl) AddComment(comment Comment)     {}
func (d *TsField) AddComment(comment Comment)        {}
func (d *TsTypeSpec) AddComment(comment Comment)     {}
func (d *TsObject) AddComment(comment Comment)       {}

// Comment
func (cg *TsComment) SetComment(comment string) {}
