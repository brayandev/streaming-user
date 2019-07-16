package user

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

// Validators implements methods of validators package.
type Validators interface {
	JSONValidator(content interface{}) (bool, error)
}

// ValidatorsImpl validators dependencies.
type ValidatorsImpl struct{}

// NewValidators constructor of validators.
func NewValidators() *ValidatorsImpl {
	return &ValidatorsImpl{}
}

// JSONValidator validate input of user.
func (v ValidatorsImpl) JSONValidator(content interface{}) (bool, error) {
	schema := gojsonschema.NewReferenceLoader("file://../json-schema/user-schema.json")

	data := gojsonschema.NewGoLoader(content)
	result, vErr := gojsonschema.Validate(schema, data)
	if vErr != nil {
		return false, vErr
	}

	if !result.Valid() {
		for _, err := range result.Errors() {
			return false, errors.New(err.Description())
		}
	}
	return true, nil
}
