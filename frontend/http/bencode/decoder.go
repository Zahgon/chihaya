package bencode

import (
	"bufio"
	"io"
)

// A Decoder reads bencoded objects from an input stream.
type Decoder struct {
	r *bufio.Reader
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder { _ = "STUB: not implemented"; return nil }

// Decode unmarshals the next bencoded value in the stream.
func (dec *Decoder) Decode() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unmarshal deserializes and returns the bencoded value in buf.
		nil
}

func Unmarshal(buf []byte) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// unmarshal reads bencoded values from a bufio.Reader
func unmarshal(r *bufio.Reader) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func readTerminator(r io.ByteScanner, term byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func readTerminatedInt(r *bufio.Reader, term byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func readList(r *bufio.Reader) (List, error) { _ = "STUB: not implemented"; return *new(List), nil }

func readDict(r *bufio.Reader) (Dict, error) { _ = "STUB: not implemented"; return *new(Dict), nil }
