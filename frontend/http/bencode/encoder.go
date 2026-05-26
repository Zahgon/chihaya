package bencode

import (
	"io"
)

// An Encoder writes bencoded objects to an output stream.
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder { _ = "STUB: not implemented"; return nil }

// Encode writes the bencoding of v to the stream.
func (enc *Encoder) Encode(v interface{}) error { _ = "STUB: not implemented"; return nil }

// Marshal returns the bencoding of v.
func Marshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Marshaler is the interface implemented by objects that can marshal
// themselves.
type Marshaler interface {
	MarshalBencode() ([]byte, error)
}

// marshal writes types bencoded to an io.Writer.
func marshal(w io.Writer, data interface{}) (err error) { _ = "STUB: not implemented"; return nil }

// Assume seconds

func marshalInt(w io.Writer, v int64) error { _ = "STUB: not implemented"; return nil }

func marshalUint(w io.Writer, v uint64) error { _ = "STUB: not implemented"; return nil }

func marshalBytes(w io.Writer, v []byte) error { _ = "STUB: not implemented"; return nil }

func marshalString(w io.Writer, v string) error { _ = "STUB: not implemented"; return nil }

func marshalStringSlice(w io.Writer, v []string) error { _ = "STUB: not implemented"; return nil }

func marshalList(w io.Writer, v []interface{}) error { _ = "STUB: not implemented"; return nil }

func marshalMap(w io.Writer, v map[string]interface{}) error { _ = "STUB: not implemented"; return nil }
