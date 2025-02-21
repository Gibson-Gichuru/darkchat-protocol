package protocol

import (
	"encoding/binary"
	"io"
)

// Decode reads a message from the reader with a message type header, and
// decodes it into a Payload. It returns the Payload and any error encountered.
// If the message type is unknown, it returns an error.
func Decode(r io.Reader) (Payload, error) {

	var typ uint8

	err := binary.Read(r, binary.BigEndian, &typ)

	if err != nil {
		return nil, err
	}

	switch typ {
	case HeartBeat:
		return decodeBeat(r)
	case MessageType:
		return decodeMessage(r)
	case Error:
		return decodeError(r)
	default:
		return nil, ErrorUnknownType
	}
}

// Encode writes a message to the writer with a message type header, and returns
// the number of bytes written and any error encountered. If the message type is
// unknown, it returns an error.
func Encode(w io.Writer, payload Payload, payloadType uint8) (int64, error) {

	switch payloadType {
	case HeartBeat:
		return encodeBeat(w, payload)
	case MessageType:
		return encodeMessage(w, payload)
	case Error:
		return encodeError(w, payload)

	default:
		return 0, ErrorUnknownType
	}

}
