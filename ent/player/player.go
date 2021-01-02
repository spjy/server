// Code generated by entc, DO NOT EDIT.

package player

import (
	"time"
)

const (
	// Label holds the string label denoting the player type in the database.
	Label = "player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTotalMinutes holds the string denoting the total_minutes field in the database.
	FieldTotalMinutes = "total_minutes"
	// FieldSessionStartedAt holds the string denoting the session_started_at field in the database.
	FieldSessionStartedAt = "session_started_at"
	// FieldSessionEndedAt holds the string denoting the session_ended_at field in the database.
	FieldSessionEndedAt = "session_ended_at"
	// FieldLastSeenAt holds the string denoting the last_seen_at field in the database.
	FieldLastSeenAt = "last_seen_at"

	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeServers holds the string denoting the servers edge name in mutations.
	EdgeServers = "servers"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeIdentifiers holds the string denoting the identifiers edge name in mutations.
	EdgeIdentifiers = "identifiers"

	// Table holds the table name of the player in the database.
	Table = "players"
	// MetadataTable is the table the holds the metadata relation/edge.
	MetadataTable = "players"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "player_metadata"
	// ServersTable is the table the holds the servers relation/edge. The primary key declared below.
	ServersTable = "player_servers"
	// ServersInverseTable is the table name for the Server entity.
	// It exists in this package in order to avoid circular dependency with the "server" package.
	ServersInverseTable = "servers"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "players"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "player_user"
	// IdentifiersTable is the table the holds the identifiers relation/edge.
	IdentifiersTable = "player_identifiers"
	// IdentifiersInverseTable is the table name for the PlayerIdentifier entity.
	// It exists in this package in order to avoid circular dependency with the "playeridentifier" package.
	IdentifiersInverseTable = "player_identifiers"
	// IdentifiersColumn is the table column denoting the identifiers relation/edge.
	IdentifiersColumn = "player_identifier_player"
)

// Columns holds all SQL columns for player fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldTotalMinutes,
	FieldSessionStartedAt,
	FieldSessionEndedAt,
	FieldLastSeenAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Player type.
var ForeignKeys = []string{
	"player_metadata",
	"player_user",
}

var (
	// ServersPrimaryKey and ServersColumn2 are the table columns denoting the
	// primary key for the servers relation (M2M).
	ServersPrimaryKey = []string{"player_id", "server_id"}
)

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
	// DefaultTotalMinutes holds the default value on creation for the total_minutes field.
	DefaultTotalMinutes int
)
