package docx

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/docx-go/opc"
	"github.com/docx-go/oxml"
)

func Open(pathOrStream interface{}) (*Document, error) {
	switch v := pathOrStream.(type) {
	case string:
		if v == "" {
			return New(), nil
		}
		return openDocxFile(v)
	case nil:
		return New(), nil
	default:
		return nil, fmt.Errorf("unsupported document source")
	}
}

func openDocxFile(path string) (*Document, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", path)
	}

	pkg, err := opc.Open(path)
	if err != nil {
		return nil, err
	}

	mainPart, err := pkg.MainDocumentPart()
	if err != nil {
		return nil, err
	}

	if mainPart.ContentType() != opc.ContentTypeWMLDocumentMain {
		return nil, fmt.Errorf("file '%s' is not a Word file, content type is '%s'", path, mainPart.ContentType())
	}

	data := mainPart.Blob()
	element, err := oxml.ParseXML(data)
	if err != nil {
		return nil, err
	}

	return NewDocument(element, mainPart), nil
}

func defaultDocxPath() string {
	return filepath.Join(getExecDir(), "templates", "default.docx")
}

func getExecDir() string {
	ex, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(ex)
}
