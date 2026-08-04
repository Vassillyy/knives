package errors

import "errors"

var (
	ErrNotFound   = errors.New("resource not found")
	ErrValidation = errors.New("validation error")
)

const (
	MsgInvalidRequestBody  = "invalid request body"
	MsgInternalServerError = "internal server error"
	MsgKnifeNotFound       = "knife not found"
	MsgPhotoNotFound       = "photo not found"
	MsgFileRequired        = "file is required"
)
