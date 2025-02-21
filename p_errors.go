package protocol

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"io"
)

/*
	The Error package has the following structure
	--------------------------------------------------------
	ErrorType(uint8)| Size(uint32)| ErrorMessage(string)
	--------------------------------------------------------

*/

type Error_ string

// String implements the Stringer interface.
// It returns a string representation of the error message.
func (e Error_) String() string {
	return string(e)
}

func (e Error_) Byte() []byte {
	return []byte(e)
}

// WriteTo implements the io.WriterTo interface.
// It writes an error message to the writer with a base64 encoded string.
// It returns the number of bytes written and any error encountered.
func (e Error_) WriteTo(w io.Writer) (int64, error) {

	encoded := base64.StdEncoding.EncodeToString(e.Byte())
	o, err := w.Write([]byte(encoded))

	return int64(o), err
}

// ReadFrom implements the io.ReaderFrom interface.
// It reads an error message from the reader with a size header, and decodes it into an _Error.
// It returns the number of bytes read and any error encountered.
func (e *Error_) ReadFrom(r io.Reader) (int64, error) {
	var size uint32

	err := binary.Read(r, binary.BigEndian, &size)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, size)

	n, err := r.Read(buf)

	if err != nil {
		return int64(n), err
	}

	decoded, err := base64.StdEncoding.DecodeString(string(buf))

	if err != nil {
		return int64(n), err
	}

	*e = Error_(decoded)

	return int64(n), nil
}

// encodeError writes an error message to the writer with a message type header.
// It returns the number of bytes written and any error encountered.
func encodeError(w io.Writer, payload Payload) (int64, error) {

	err := binary.Write(w, binary.BigEndian, Error)

	if err != nil {
		return 0, err
	}

	var n int64 = 1

	err = binary.Write(w, binary.BigEndian, uint32(len(
		base64.StdEncoding.EncodeToString(payload.Byte()),
	)))

	if err != nil {
		return n, err
	}
	n += 1

	o, err := payload.WriteTo(w)

	return n + o, err
}

// decodeError reads an error message from the reader with a message type header,
// and decodes it into a Payload. It returns the Payload and any error encountered.
// If the size of the error message exceeds MaxPayloadsize, it returns an error.
func decodeError(r io.Reader) (Payload, error) {

	var size uint32

	var payload = new(Error_)

	err := binary.Read(r, binary.BigEndian, &size)

	if err != nil {
		return nil, err
	}

	if uint32(size) > MaxPayloadsize {
		return nil, ErrorMaxPayloadSize
	}

	payloadSize := size

	errorBuf := new(bytes.Buffer)

	err = binary.Write(errorBuf, binary.BigEndian, payloadSize)

	if err != nil {
		return nil, err
	}

	_, err = payload.ReadFrom(io.MultiReader(errorBuf, r))

	if err != nil {
		return nil, err
	}

	return payload, nil

}
