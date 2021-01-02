// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/responserms/server/ent/maptype"
	"github.com/responserms/server/ent/metadata"
	"github.com/responserms/server/ent/metadataschema"
	"github.com/responserms/server/ent/user"
)

// Metadata is the model entity for the Metadata schema.
type Metadata struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Data holds the value of the "data" field.
	Data map[string]interface{} `json:"data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MetadataQuery when eager-loading is set.
	Edges             MetadataEdges `json:"edges"`
	map_type_metadata *int
	metadata_schema   *int
	user_metadata     *int
}

// MetadataEdges holds the relations/edges for other nodes in the graph.
type MetadataEdges struct {
	// Schema holds the value of the schema edge.
	Schema *MetadataSchema
	// User holds the value of the user edge.
	User *User
	// MapType holds the value of the map_type edge.
	MapType *MapType
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SchemaOrErr returns the Schema value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MetadataEdges) SchemaOrErr() (*MetadataSchema, error) {
	if e.loadedTypes[0] {
		if e.Schema == nil {
			// The edge schema was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: metadataschema.Label}
		}
		return e.Schema, nil
	}
	return nil, &NotLoadedError{edge: "schema"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MetadataEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MapTypeOrErr returns the MapType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MetadataEdges) MapTypeOrErr() (*MapType, error) {
	if e.loadedTypes[2] {
		if e.MapType == nil {
			// The edge map_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: maptype.Label}
		}
		return e.MapType, nil
	}
	return nil, &NotLoadedError{edge: "map_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Metadata) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case metadata.FieldData:
			values[i] = &[]byte{}
		case metadata.FieldID:
			values[i] = &sql.NullInt64{}
		case metadata.ForeignKeys[0]: // map_type_metadata
			values[i] = &sql.NullInt64{}
		case metadata.ForeignKeys[1]: // metadata_schema
			values[i] = &sql.NullInt64{}
		case metadata.ForeignKeys[2]: // user_metadata
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Metadata", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Metadata fields.
func (m *Metadata) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case metadata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case metadata.FieldData:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %v", err)
				}
			}
		case metadata.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field map_type_metadata", value)
			} else if value.Valid {
				m.map_type_metadata = new(int)
				*m.map_type_metadata = int(value.Int64)
			}
		case metadata.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field metadata_schema", value)
			} else if value.Valid {
				m.metadata_schema = new(int)
				*m.metadata_schema = int(value.Int64)
			}
		case metadata.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_metadata", value)
			} else if value.Valid {
				m.user_metadata = new(int)
				*m.user_metadata = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySchema queries the schema edge of the Metadata.
func (m *Metadata) QuerySchema() *MetadataSchemaQuery {
	return (&MetadataClient{config: m.config}).QuerySchema(m)
}

// QueryUser queries the user edge of the Metadata.
func (m *Metadata) QueryUser() *UserQuery {
	return (&MetadataClient{config: m.config}).QueryUser(m)
}

// QueryMapType queries the map_type edge of the Metadata.
func (m *Metadata) QueryMapType() *MapTypeQuery {
	return (&MetadataClient{config: m.config}).QueryMapType(m)
}

// Update returns a builder for updating this Metadata.
// Note that, you need to call Metadata.Unwrap() before calling this method, if this Metadata
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Metadata) Update() *MetadataUpdateOne {
	return (&MetadataClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Metadata) Unwrap() *Metadata {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Metadata is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Metadata) String() string {
	var builder strings.Builder
	builder.WriteString("Metadata(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", data=")
	builder.WriteString(fmt.Sprintf("%v", m.Data))
	builder.WriteByte(')')
	return builder.String()
}

// MetadataSlice is a parsable slice of Metadata.
type MetadataSlice []*Metadata

func (m MetadataSlice) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
