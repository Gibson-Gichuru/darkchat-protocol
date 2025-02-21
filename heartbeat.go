package protocol

import (
	"encoding/binary"
	"io"
)

type Beat string

func (b Beat) String() string { return "" }
func (b Beat) Byte() []byte   { return []byte("") }

// WriteTo implements the io.WriterTo interface.
// It writes the heartbeat message to the writer.
// It returns the number of bytes written and any error encountered.
func (b Beat) WriteTo(w io.Writer) (int64, error) {

	o, err := w.Write([]byte(""))

	if err != nil {

		return 0, err
	}

	return int64(o), nil
}

// ReadFrom implements the io.ReaderFrom interface.
// It reads a heartbeat message from the reader.
// It returns the number of bytes read and any error encountered.
func (b *Beat) ReadFrom(r io.Reader) (int64, error) {

	return 0, nil
}

// encodeBeat writes a heartbeat message to the writer with a message type header.
// It returns the number of bytes written and any error encountered.
func encodeBeat(w io.Writer, payload Payload) (int64, error) {
	err := binary.Write(w, binary.BigEndian, HeartBeat)

	if err != nil {
		return 0, err
	}

	var n int64 = 1

	o, err := payload.WriteTo(w)

	return n + o, err
}

// decodeBeat reads a heartbeat message from the reader with a message type header,
// and decodes it into a Payload. It returns the Payload and any error encountered.
func decodeBeat(_ io.Reader) (Payload, error) {
	payload := new(Beat)

	return payload, nil
}
