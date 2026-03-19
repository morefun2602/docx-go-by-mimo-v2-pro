package opc

import "fmt"

type ErrInvalidPackURI string

func (e ErrInvalidPackURI) Error() string {
	return fmt.Sprintf("invalid pack URI: %s", string(e))
}
