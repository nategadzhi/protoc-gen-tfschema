package builder

import "google.golang.org/protobuf/reflect/protoreflect"

// SchemaMap string -> *Schema
type SchemaMap map[string]*Schema

// BuildSchemaMapFromMessage reads and parses resource into schema map
func BuildSchemaMapFromMessage(message *protoreflect.MessageDescriptor, resource *Resource, resourceMap ResourceMap) {
	fields := (*message).Fields()

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		schema := BuildSchemaFromField(&field, resourceMap)
		resource.Schema[schema.Name] = schema
	}
}
