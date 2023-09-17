// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPrivateKey holds the string denoting the private_key field in the database.
	FieldPrivateKey = "private_key"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeChats holds the string denoting the chats edge name in mutations.
	EdgeChats = "chats"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "user_messages"
	// ChatsTable is the table that holds the chats relation/edge. The primary key declared below.
	ChatsTable = "chat_members"
	// ChatsInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	ChatsInverseTable = "chats"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldCreatedAt,
	FieldPassword,
	FieldPrivateKey,
	FieldPublicKey,
}

var (
	// ChatsPrimaryKey and ChatsColumn2 are the table columns denoting the
	// primary key for the chats relation (M2M).
	ChatsPrimaryKey = []string{"chat_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultPrivateKey holds the default value on creation for the "private_key" field.
	DefaultPrivateKey string
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByPrivateKey orders the results by the private_key field.
func ByPrivateKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivateKey, opts...).ToFunc()
}

// ByMessagesCount orders the results by messages count.
func ByMessagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMessagesStep(), opts...)
	}
}

// ByMessages orders the results by messages terms.
func ByMessages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMessagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByChatsCount orders the results by chats count.
func ByChatsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChatsStep(), opts...)
	}
}

// ByChats orders the results by chats terms.
func ByChats(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
	)
}
func newChatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ChatsTable, ChatsPrimaryKey...),
	)
}
