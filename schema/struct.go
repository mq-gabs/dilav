package schema

import (
	"errors"
	"fmt"
)

type structSchema struct {
	baseSchema[any]
	schemas map[string]Schema[any]
}

func Struct(fields map[string]Schema[any]) *structSchema {
	return &structSchema{
		baseSchema: newBaseSchema[any](),
		schemas:    fields,
	}
}

func (ss *structSchema) Validate(structValue any) error {
	structValueTyped, ok := structValue.(SchemaJSONMarshaler)
	if !ok {
		return errors.New("struct must implement SchemaJSON")
	}

	json := structValueTyped.SchemaJSON()

	var err error
	for key, value := range json {
		schema, ok := ss.schemas[key]
		if !ok {
			return fmt.Errorf("no schema set for field: %v", key)
		}

		schemaErr := schema.Validate(value)
		if schemaErr != nil {
			err = errors.Join(err, fmt.Errorf("[%v]: %v", key, schemaErr.Error()))
		}
	}

	return err
}
