# docx-go

> **Important Notice**  
> This repository is translated from [python-openxml/python-docx](https://github.com/python-openxml/python-docx) and serves as its Go implementation.  
> [中文版](README.zh.md)

---

### Overview

**docx-go** is a Go library for reading, creating, and updating Microsoft Word 2007+ (.docx) files.

This repository is **translated from [python-openxml/python-docx](https://github.com/python-openxml/python-docx)**. It ports the core functionality and API design of the original Python library to Go, offering Go developers a similar document manipulation experience.

### Features

- Create new documents or open existing .docx files
- Add paragraphs, headings, and tables
- Support for styles, page breaks, and sections
- Handle images, shapes, and drawings
- Comments support
- Low-level implementation based on OPC (Open Packaging Conventions) and OOXML

### Installation

```bash
go get github.com/docx-go
```

### Quick Example

```go
package main

import (
    "github.com/docx-go/docx"
)

func main() {
    // Create a new document
    doc := docx.New()
    doc.AddParagraph("It was a dark and stormy night.", "")
    doc.Save("output.docx")

    // Open an existing document
    doc2, _ := docx.Open("output.docx")
    paras := doc2.Paragraphs()
    // paras[0].Text() returns "It was a dark and stormy night."
}
```

### Project Structure

- `docx/` - Core API: document, paragraph, table, section
- `text/` - Text, paragraph, run
- `opc/` - OPC package and parts
- `oxml/` - OOXML parsing and generation
- `styles/` - Styles
- `comments/` - Comments
- `image/`, `drawing/`, `dml/` - Images and drawings

---

## License

MIT (aligned with [python-docx](https://github.com/python-openxml/python-docx))
