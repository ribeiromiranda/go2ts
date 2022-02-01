package typescript

import (
	"go/ast"
	"go2ts/internal/typescript/tsast"
	"strconv"
)

type comment struct {
	t string
}

func (c *comment) Set(t string) {
	c.t = t
}

func (p *printer) commentGroups(name string, commentGroups []*ast.CommentGroup, node tsast.CommentGroups) {
	if len(commentGroups) == 0 {
		return
	}

	p.printDebugf("%s: []*ast.CommentGroup\n", name)
	p.indentDebug++
	for _, commentGroup := range commentGroups {
		nodeCommentGroup := &tsast.TsCommentGroup{}
		p.commentGroup("c", commentGroup, nodeCommentGroup)
		node.AddCommentGroup(nodeCommentGroup)
	}
	p.indentDebug--
}

func (p *printer) commentGroup(name string, commentGroup *ast.CommentGroup, node tsast.CommentGroup) {
	if commentGroup == nil {
		return
	}
	if len(commentGroup.List) == 0 {
		return
	}

	p.printDebugf("%s: *ast.CommentGroup\n", name)
	p.indentDebug++
	for i, c := range commentGroup.List {
		nodeComment := &tsast.TsComment{}
		p.comment(strconv.Itoa(i), c, nodeComment)
		node.AddComment(nodeComment)
	}
	p.indentDebug--
}

func (p *printer) comment(name string, comment *ast.Comment, node tsast.Comment) {
	p.printDebugf("%s: *ast.Comment\n", name)
	p.indentDebug++
	p.printDebugf("Slash: %d\n", comment.Slash)
	text := comment.Text
	p.printDebugf("Text: %s\n", text)
	p.indentDebug--
	p.indentDebug--

	node.SetComment(text)
}
