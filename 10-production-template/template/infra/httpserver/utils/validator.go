package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func init() {
	// Use json tag names in validation errors (e.g. "name" not "Name")
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "" || name == "-" {
			return fld.Name
		}
		return name
	})
}

// FieldError is one client-facing validation problem.
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationError is returned by DecodeAndValidate.
type ValidationError struct {
	Kind   string       // "json" | "validation"
	Fields []FieldError // populated for validation failures
}

func (e *ValidationError) Error() string {
	if e.Kind == "json" {
		return "invalid json body"
	}
	return "validation failed"
}

// DecodeAndValidate decodes JSON body into dst, then runs struct validation.
func DecodeAndValidate(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // optional but good for strict APIs

	if err := dec.Decode(dst); err != nil {
		return mapJSONDecodeError(err)
	}

	// reject trailing garbage: {"name":"x"} extra text
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return &ValidationError{Kind: "json"}
		}
		return mapJSONDecodeError(err)
	}

	if err := validate.Struct(dst); err != nil {
		return mapValidationError(err)
	}

	return nil
}

func mapJSONDecodeError(err error) error {
	var syntaxErr *json.SyntaxError
	var typeErr *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxErr):
		return &ValidationError{Kind: "json"}

	case errors.As(err, &typeErr):
		field := typeErr.Field
		if field == "" {
			field = "body"
		}
		return &ValidationError{
			Kind: "validation",
			Fields: []FieldError{{
				Field:   field,
				Message: fmt.Sprintf("%s must be %s", field, typeErr.Type.String()),
			}},
		}

	default:
		if field, ok := parseUnknownField(err); ok {
			return &ValidationError{
				Kind: "validation",
				Fields: []FieldError{{
					Field:   field,
					Message: fmt.Sprintf("%s is not allowed", field),
				}},
			}
		}
		return &ValidationError{Kind: "json"}
	}
}

func parseUnknownField(err error) (string, bool) {
	// Go returns: json: unknown field "age"
	msg := err.Error()
	const marker = `unknown field "`
	i := strings.Index(msg, marker)
	if i == -1 {
		return "", false
	}
	rest := msg[i+len(marker):]
	j := strings.Index(rest, `"`)
	if j == -1 {
		return "", false
	}
	return rest[:j], true
}

func mapValidationError(err error) error {
	var verrs validator.ValidationErrors
	if !errors.As(err, &verrs) {
		return &ValidationError{Kind: "validation"}
	}

	fields := make([]FieldError, 0, len(verrs))
	for _, fe := range verrs {
		fields = append(fields, FieldError{
			Field:   fe.Field(),
			Message: validationMessage(fe),
		})
	}

	return &ValidationError{
		Kind:   "validation",
		Fields: fields,
	}
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be >= %s", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be <= %s", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s is invalid", fe.Field())
	}
}
