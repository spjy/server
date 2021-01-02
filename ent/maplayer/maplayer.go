// Code generated by entc, DO NOT EDIT.

package maplayer

import (
	"time"
)

const (
	// Label holds the string label denoting the maplayer type in the database.
	Label = "map_layer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldURLTemplate holds the string denoting the url_template field in the database.
	FieldURLTemplate = "url_template"
	// FieldIsPublic holds the string denoting the is_public field in the database.
	FieldIsPublic = "is_public"

	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeMapType holds the string denoting the map_type edge name in mutations.
	EdgeMapType = "map_type"

	// Table holds the table name of the maplayer in the database.
	Table = "map_layers"
	// MetadataTable is the table the holds the metadata relation/edge.
	MetadataTable = "map_layers"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "map_layer_metadata"
	// MapTypeTable is the table the holds the map_type relation/edge.
	MapTypeTable = "map_layers"
	// MapTypeInverseTable is the table name for the MapType entity.
	// It exists in this package in order to avoid circular dependency with the "maptype" package.
	MapTypeInverseTable = "map_types"
	// MapTypeColumn is the table column denoting the map_type relation/edge.
	MapTypeColumn = "map_layer_map_type"
)

// Columns holds all SQL columns for maplayer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldURLTemplate,
	FieldIsPublic,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the MapLayer type.
var ForeignKeys = []string{
	"map_layer_metadata",
	"map_layer_map_type",
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

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)
