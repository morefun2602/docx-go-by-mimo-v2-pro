package comments_test

import (
	"testing"

	"github.com/docx-go/comments"
	"github.com/docx-go/oxml"
)

func TestNewComment(t *testing.T) {
	elem := oxml.NewElement("w:comment")
	elem.SetAttr("id", "1")
	elem.SetAttr("author", "Test Author")
	elem.SetAttr("initials", "TA")

	p := oxml.NewElement("w:p")
	elem.AddChild(p)
	r := oxml.NewElement("w:r")
	p.AddChild(r)
	tElem := oxml.NewElement("w:t")
	r.AddChild(tElem)
	tElem.SetText("Test comment")

	comment := comments.NewComment(elem)
	if comment.CommentID() != 1 {
		t.Errorf("CommentID() = %d, want 1", comment.CommentID())
	}
	if comment.Author() != "Test Author" {
		t.Errorf("Author() = %s, want Test Author", comment.Author())
	}
	if comment.Initials() != "TA" {
		t.Errorf("Initials() = %s, want TA", comment.Initials())
	}
}

func TestCommentSetAuthor(t *testing.T) {
	elem := oxml.NewElement("w:comment")
	elem.SetAttr("id", "1")

	comment := comments.NewComment(elem)
	comment.SetAuthor("New Author")
	if comment.Author() != "New Author" {
		t.Errorf("Author() = %s, want New Author", comment.Author())
	}
}

func TestCommentSetText(t *testing.T) {
	elem := oxml.NewElement("w:comment")
	elem.SetAttr("id", "1")

	comment := comments.NewComment(elem)
	comment.SetText("New text")
	if comment.Text() != "New text" {
		t.Errorf("Text() = %s, want New text", comment.Text())
	}
}

func TestCommentsCollection(t *testing.T) {
	elem := oxml.NewElement("w:comments")

	comment1 := oxml.NewElement("w:comment")
	comment1.SetAttr("id", "1")
	comment1.SetAttr("author", "Author 1")
	p1 := oxml.NewElement("w:p")
	comment1.AddChild(p1)
	r1 := oxml.NewElement("w:r")
	p1.AddChild(r1)
	t1 := oxml.NewElement("w:t")
	r1.AddChild(t1)
	t1.SetText("Comment 1")
	elem.AddChild(comment1)

	comment2 := oxml.NewElement("w:comment")
	comment2.SetAttr("id", "2")
	comment2.SetAttr("author", "Author 2")
	p2 := oxml.NewElement("w:p")
	comment2.AddChild(p2)
	r2 := oxml.NewElement("w:r")
	p2.AddChild(r2)
	t2 := oxml.NewElement("w:t")
	r2.AddChild(t2)
	t2.SetText("Comment 2")
	elem.AddChild(comment2)

	commentsCol := comments.NewComments(elem, nil)
	if commentsCol.Len() != 2 {
		t.Errorf("Len() = %d, want 2", commentsCol.Len())
	}

	c := commentsCol.GetByID(1)
	if c == nil {
		t.Fatal("GetByID(1) returned nil")
	}
	if c.Author() != "Author 1" {
		t.Errorf("GetByID(1).Author() = %s, want Author 1", c.Author())
	}
}

func TestCommentsAddComment(t *testing.T) {
	elem := oxml.NewElement("w:comments")
	commentsCol := comments.NewComments(elem, nil)

	comment := commentsCol.AddComment("New comment", "Author", "A")
	if comment == nil {
		t.Fatal("AddComment() returned nil")
	}
	if comment.Author() != "Author" {
		t.Errorf("Author() = %s, want Author", comment.Author())
	}
	if commentsCol.Len() != 1 {
		t.Errorf("Len() = %d, want 1", commentsCol.Len())
	}
}

func TestCommentsRemoveComment(t *testing.T) {
	elem := oxml.NewElement("w:comments")
	comment := oxml.NewElement("w:comment")
	comment.SetAttr("id", "1")
	elem.AddChild(comment)

	commentsCol := comments.NewComments(elem, nil)
	commentsCol.RemoveComment(1)
	if commentsCol.Len() != 0 {
		t.Errorf("Len() = %d, want 0", commentsCol.Len())
	}
}
