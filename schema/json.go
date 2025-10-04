package schema

type SchemaJSONMarshaler interface {
	SchemaJSON() map[string]any
}
