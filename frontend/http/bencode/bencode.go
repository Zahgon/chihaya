// Package bencode implements bencoding of data as defined in BEP 3 using
// type assertion over reflection for performance.
package bencode

// Enforce that Dict implements the Marshaler interface.
var _ Marshaler = Dict{}

// Dict represents a bencode dictionary.
type Dict map[string]interface{}

// NewDict allocates the memory for a Dict.
func NewDict() Dict {
	_ = "STUB: not implemented"

	// MarshalBencode implements the Marshaler interface for Dict.
	return *new(Dict)
}

func (d Dict) MarshalBencode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Enforce that List implements the Marshaler interface.
var _ Marshaler = List{}

// List represents a bencode list.
type List []interface{}

// MarshalBencode implements the Marshaler interface for List.
func (l List) MarshalBencode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NewList allocates the memory for a List.
func NewList() List { _ = "STUB: not implemented"; return *new(List) }
