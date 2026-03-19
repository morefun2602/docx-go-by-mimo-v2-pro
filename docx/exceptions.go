package docx

import "fmt"

type InvalidFileFormatError struct {
	msg string
}

func NewInvalidFileFormatError(msg string) *InvalidFileFormatError {
	return &InvalidFileFormatError{msg: msg}
}

func (e *InvalidFileFormatError) Error() string {
	return fmt.Sprintf("invalid file format: %s", e.msg)
}

type UnexpectedFileFormatError struct {
	msg string
}

func NewUnexpectedFileFormatError(msg string) *UnexpectedFileFormatError {
	return &UnexpectedFileFormatError{msg: msg}
}

func (e *UnexpectedFileFormatError) Error() string {
	return fmt.Sprintf("unexpected file format: %s", e.msg)
}

type PackageNotFoundError struct {
	msg string
}

func NewPackageNotFoundError(msg string) *PackageNotFoundError {
	return &PackageNotFoundError{msg: msg}
}

func (e *PackageNotFoundError) Error() string {
	return fmt.Sprintf("package not found: %s", e.msg)
}

type RelationshipNotFoundError struct {
	msg string
}

func NewRelationshipNotFoundError(msg string) *RelationshipNotFoundError {
	return &RelationshipNotFoundError{msg: msg}
}

func (e *RelationshipNotFoundError) Error() string {
	return fmt.Sprintf("relationship not found: %s", e.msg)
}

type PartNotFoundError struct {
	msg string
}

func NewPartNotFoundError(msg string) *PartNotFoundError {
	return &PartNotFoundError{msg: msg}
}

func (e *PartNotFoundError) Error() string {
	return fmt.Sprintf("part not found: %s", e.msg)
}

type ValueError struct {
	msg string
}

func NewValueError(msg string) *ValueError {
	return &ValueError{msg: msg}
}

func (e *ValueError) Error() string {
	return fmt.Sprintf("value error: %s", e.msg)
}
