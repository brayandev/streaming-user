package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xeipuuv/gojsonschema"
)

// Validator validator methods.
type Validator interface {
	Validate(content interface{}) (bool, error)
}

// SchemaValidator validator dependencies.
type SchemaValidator struct {
	schemaFolder string
	schemaName   string
}

// NewValidator validator constructor.
func NewValidator(schemaFolder, schemaName string) *SchemaValidator {
	return &SchemaValidator{
		schemaFolder: schemaFolder,
		schemaName:   schemaName,
	}
}

// Validate validate input values with json schema.
func (sv *SchemaValidator) Validate(content interface{}) (bool, error) {
	schema, err := sv.getSchema(sv.schemaName)
	if err != nil {
		return false, NewUnprocessableEntityError(err.Error())
	}

	b, mErr := json.Marshal(content)
	if mErr != nil {
		return false, mErr
	}

	schemaLoader := gojsonschema.NewBytesLoader(schema)
	msgLoader := gojsonschema.NewBytesLoader(b)

	result, vErr := gojsonschema.Validate(schemaLoader, msgLoader)
	if vErr != nil {
		return false, vErr
	}

	return result.Valid(), nil
}

func (sv *SchemaValidator) getSchema(schemaName string) ([]byte, error) {
	fileName := fmt.Sprintf("./%s/%s.json", sv.schemaFolder, sv.schemaName)
	schema, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return schema, nil
}
