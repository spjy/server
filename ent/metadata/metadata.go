// Code generated by entc, DO NOT EDIT.

package metadata

const (
	// Label holds the string label denoting the metadata type in the database.
	Label = "metadata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"

	// EdgeSchema holds the string denoting the schema edge name in mutations.
	EdgeSchema = "schema"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeMapType holds the string denoting the map_type edge name in mutations.
	EdgeMapType = "map_type"

	// Table holds the table name of the metadata in the database.
	Table = "metadata"
	// SchemaTable is the table the holds the schema relation/edge.
	SchemaTable = "metadata"
	// SchemaInverseTable is the table name for the MetadataSchema entity.
	// It exists in this package in order to avoid circular dependency with the "metadataschema" package.
	SchemaInverseTable = "metadata_schemas"
	// SchemaColumn is the table column denoting the schema relation/edge.
	SchemaColumn = "metadata_schema"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "metadata"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_metadata"
	// MapTypeTable is the table the holds the map_type relation/edge.
	MapTypeTable = "metadata"
	// MapTypeInverseTable is the table name for the MapType entity.
	// It exists in this package in order to avoid circular dependency with the "maptype" package.
	MapTypeInverseTable = "map_types"
	// MapTypeColumn is the table column denoting the map_type relation/edge.
	MapTypeColumn = "map_type_metadata"
)

// Columns holds all SQL columns for metadata fields.
var Columns = []string{
	FieldID,
	FieldData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Metadata type.
var ForeignKeys = []string{
	"map_type_metadata",
	"metadata_schema",
	"user_metadata",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
