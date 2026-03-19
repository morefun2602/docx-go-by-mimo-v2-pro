package comments

import (
	"fmt"
	"time"

	"github.com/docx-go/oxml"
)

type Comment struct {
	element   *oxml.Element
	commentID int
	author    string
	initials  string
	date      time.Time
	text      string
}

func NewComment(element *oxml.Element) *Comment {
	c := &Comment{element: element}
	c.parse()
	return c
}

func (c *Comment) parse() {
	if c.element == nil {
		return
	}

	if val, ok := c.element.GetAttr("id"); ok {
		c.commentID = parseInt(val)
	}
	if val, ok := c.element.GetAttr("author"); ok {
		c.author = val
	}
	if val, ok := c.element.GetAttr("initials"); ok {
		c.initials = val
	}
	if val, ok := c.element.GetAttr("date"); ok {
		t, err := time.Parse(time.RFC3339, val)
		if err == nil {
			c.date = t
		}
	}

	c.text = c.element.Text()
}

func (c *Comment) Element() *oxml.Element {
	return c.element
}

func (c *Comment) CommentID() int {
	return c.commentID
}

func (c *Comment) Author() string {
	return c.author
}

func (c *Comment) SetAuthor(author string) {
	c.author = author
	c.element.SetAttr("author", author)
}

func (c *Comment) Initials() string {
	return c.initials
}

func (c *Comment) SetInitials(initials string) {
	c.initials = initials
	c.element.SetAttr("initials", initials)
}

func (c *Comment) Date() time.Time {
	return c.date
}

func (c *Comment) SetDate(date time.Time) {
	c.date = date
	c.element.SetAttr("date", date.Format(time.RFC3339))
}

func (c *Comment) Text() string {
	return c.text
}

func (c *Comment) SetText(text string) {
	c.text = text
	c.element.RemoveAllChildren()
	p := oxml.NewElement("w:p")
	c.element.AddChild(p)
	r := oxml.NewElement("w:r")
	p.AddChild(r)
	t := oxml.NewElement("w:t")
	r.AddChild(t)
	t.SetText(text)
}

type Comments struct {
	element  *oxml.Element
	parent   interface{}
	comments []*Comment
	nextID   int
}

func NewComments(element *oxml.Element, parent interface{}) *Comments {
	c := &Comments{
		element: element,
		parent:  parent,
		nextID:  1,
	}
	c.parse()
	return c
}

func (c *Comments) parse() {
	c.comments = nil
	maxID := 0
	for _, child := range c.element.FindAll("w:comment") {
		comment := NewComment(child)
		c.comments = append(c.comments, comment)
		if comment.CommentID() > maxID {
			maxID = comment.CommentID()
		}
	}
	c.nextID = maxID + 1
}

func (c *Comments) Element() *oxml.Element {
	return c.element
}

func (c *Comments) Len() int {
	return len(c.comments)
}

func (c *Comments) Get(index int) *Comment {
	if index < 0 || index >= len(c.comments) {
		return nil
	}
	return c.comments[index]
}

func (c *Comments) GetByID(id int) *Comment {
	for _, comment := range c.comments {
		if comment.CommentID() == id {
			return comment
		}
	}
	return nil
}

func (c *Comments) AddComment(text, author, initials string) *Comment {
	elem := oxml.NewElement("w:comment")
	elem.SetAttr("id", formatInt(c.nextID))
	elem.SetAttr("author", author)
	if initials == "" {
		initials = author[:1]
	}
	elem.SetAttr("initials", initials)
	elem.SetAttr("date", time.Now().Format(time.RFC3339))

	p := oxml.NewElement("w:p")
	elem.AddChild(p)
	if text != "" {
		r := oxml.NewElement("w:r")
		p.AddChild(r)
		t := oxml.NewElement("w:t")
		r.AddChild(t)
		t.SetText(text)
	}

	c.element.AddChild(elem)
	comment := NewComment(elem)
	c.comments = append(c.comments, comment)
	c.nextID++

	return comment
}

func (c *Comments) RemoveComment(id int) {
	for i, comment := range c.comments {
		if comment.CommentID() == id {
			c.element.RemoveChild(comment.element)
			c.comments = append(c.comments[:i], c.comments[i+1:]...)
			return
		}
	}
}

func (c *Comments) All() []*Comment {
	return c.comments
}

func parseInt(s string) int {
	result := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
		}
	}
	return result
}

func formatInt(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = fmt.Sprintf("%c%s", '0'+n%10, result)
		n /= 10
	}
	return result
}
