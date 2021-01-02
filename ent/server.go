// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/responserms/server/ent/maptype"
	"github.com/responserms/server/ent/metadata"
	"github.com/responserms/server/ent/server"
	"github.com/responserms/server/ent/servertype"
)

// Server is the model entity for the Server schema.
type Server struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// APIUsername holds the value of the "api_username" field.
	APIUsername *string `json:"api_username,omitempty"`
	// APISecret holds the value of the "api_secret" field.
	APISecret *string `json:"api_secret,omitempty"`
	// APIAddress holds the value of the "api_address" field.
	APIAddress *string `json:"api_address,omitempty"`
	// APIPort holds the value of the "api_port" field.
	APIPort *string `json:"api_port,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServerQuery when eager-loading is set.
	Edges              ServerEdges `json:"edges"`
	server_metadata    *int
	server_server_type *int
	server_map_type    *int
}

// ServerEdges holds the relations/edges for other nodes in the graph.
type ServerEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata
	// ServerType holds the value of the server_type edge.
	ServerType *ServerType
	// MapType holds the value of the map_type edge.
	MapType *MapType
	// Players holds the value of the players edge.
	Players []*Player
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerEdges) MetadataOrErr() (*Metadata, error) {
	if e.loadedTypes[0] {
		if e.Metadata == nil {
			// The edge metadata was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: metadata.Label}
		}
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// ServerTypeOrErr returns the ServerType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerEdges) ServerTypeOrErr() (*ServerType, error) {
	if e.loadedTypes[1] {
		if e.ServerType == nil {
			// The edge server_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: servertype.Label}
		}
		return e.ServerType, nil
	}
	return nil, &NotLoadedError{edge: "server_type"}
}

// MapTypeOrErr returns the MapType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerEdges) MapTypeOrErr() (*MapType, error) {
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

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e ServerEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[3] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Server) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case server.FieldID:
			values[i] = &sql.NullInt64{}
		case server.FieldName, server.FieldAPIUsername, server.FieldAPISecret, server.FieldAPIAddress, server.FieldAPIPort:
			values[i] = &sql.NullString{}
		case server.FieldCreatedAt, server.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		case server.ForeignKeys[0]: // server_metadata
			values[i] = &sql.NullInt64{}
		case server.ForeignKeys[1]: // server_server_type
			values[i] = &sql.NullInt64{}
		case server.ForeignKeys[2]: // server_map_type
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Server", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Server fields.
func (s *Server) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case server.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case server.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case server.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case server.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case server.FieldAPIUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_username", values[i])
			} else if value.Valid {
				s.APIUsername = new(string)
				*s.APIUsername = value.String
			}
		case server.FieldAPISecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_secret", values[i])
			} else if value.Valid {
				s.APISecret = new(string)
				*s.APISecret = value.String
			}
		case server.FieldAPIAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_address", values[i])
			} else if value.Valid {
				s.APIAddress = new(string)
				*s.APIAddress = value.String
			}
		case server.FieldAPIPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field api_port", values[i])
			} else if value.Valid {
				s.APIPort = new(string)
				*s.APIPort = value.String
			}
		case server.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field server_metadata", value)
			} else if value.Valid {
				s.server_metadata = new(int)
				*s.server_metadata = int(value.Int64)
			}
		case server.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field server_server_type", value)
			} else if value.Valid {
				s.server_server_type = new(int)
				*s.server_server_type = int(value.Int64)
			}
		case server.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field server_map_type", value)
			} else if value.Valid {
				s.server_map_type = new(int)
				*s.server_map_type = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMetadata queries the metadata edge of the Server.
func (s *Server) QueryMetadata() *MetadataQuery {
	return (&ServerClient{config: s.config}).QueryMetadata(s)
}

// QueryServerType queries the server_type edge of the Server.
func (s *Server) QueryServerType() *ServerTypeQuery {
	return (&ServerClient{config: s.config}).QueryServerType(s)
}

// QueryMapType queries the map_type edge of the Server.
func (s *Server) QueryMapType() *MapTypeQuery {
	return (&ServerClient{config: s.config}).QueryMapType(s)
}

// QueryPlayers queries the players edge of the Server.
func (s *Server) QueryPlayers() *PlayerQuery {
	return (&ServerClient{config: s.config}).QueryPlayers(s)
}

// Update returns a builder for updating this Server.
// Note that, you need to call Server.Unwrap() before calling this method, if this Server
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Server) Update() *ServerUpdateOne {
	return (&ServerClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Server) Unwrap() *Server {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Server is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Server) String() string {
	var builder strings.Builder
	builder.WriteString("Server(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	if v := s.APIUsername; v != nil {
		builder.WriteString(", api_username=")
		builder.WriteString(*v)
	}
	if v := s.APISecret; v != nil {
		builder.WriteString(", api_secret=")
		builder.WriteString(*v)
	}
	if v := s.APIAddress; v != nil {
		builder.WriteString(", api_address=")
		builder.WriteString(*v)
	}
	if v := s.APIPort; v != nil {
		builder.WriteString(", api_port=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Servers is a parsable slice of Server.
type Servers []*Server

func (s Servers) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
