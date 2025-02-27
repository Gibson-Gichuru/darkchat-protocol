# Simple TCP Communication Protocol

This document outlines a lightweight communication protocol designed for two TCP nodes. The protocol supports distinct message types, each conveying a specific meaning to the communicating nodes. It prioritizes simplicity and efficiency in encoding and decoding messages while allowing flexibility for future extensions.

## Message Types

The protocol defines three primary message types:

1. **Heartbeat**
   - A simple "ping-pong" message used by nodes to confirm connectivity and advance session read/write deadlines. It carries no payload, keeping it lightweight.

2. **Message**
   - Enables nodes to exchange string-based messages. This type includes headers and a payload for extensible communication.

3. **Error**
   - Allows nodes to send predefined error types as strings, facilitating error reporting between nodes.

## Type Design

The protocol employs a straightforward design to differentiate message types and ensure seamless encoding and decoding:

- **Identification Byte**: Each message begins with a single `uint8` (1 byte) value that identifies its type. This simplicity minimizes overhead and aids in quick parsing.
  - For example, `0` might represent Heartbeat, `1` for Message, and `2` for Error (specific values to be assigned during implementation).

## Message Design

The structure for the **Message** type is as follows:

```
| type (uint8) | header size (uint32) | headers | payload |
```

- **type (uint8)**: 1 byte identifying the message type.
- **header size (uint32)**: 4 bytes specifying the length of the headers section in bytes.
- **headers**: A JSON-encoded payload describing the subsequent payload (e.g., its size). Future updates may extend this to include payload type and encoding details.
- **payload**: The actual data being transmitted, typically a string in the current design.

### Headers Structure
The headers are defined in Go as:

```go
type PayloadHeaders struct {
    Size     uint32  // Size of the payload in bytes
    Type     uint8   // Type of the payload (for future use)
    Encoding string  // Encoding of the payload (e.g., "utf-8", for future use)
}
```

The initial implementation uses `Size` to indicate payload length, with `Type` and `Encoding` reserved for future enhancements (e.g., supporting non-string payloads or alternative encodings).

## Heartbeat Design

The structure for the **Heartbeat** type is minimal:

```
| type (uint8) |
```

- **type (uint8)**: 1 byte identifying the Heartbeat type.
- **No payload**: Since this type only advances deadlines, no additional data is included.

**Note**: The lack of a size limiter introduces a potential vulnerability. A malicious node could append a large, unexpected payload after the type byte, potentially overwhelming the receiving node and leading to a Denial-of-Service (DoS) attack. To mitigate this, consider adding a check to reject any data beyond the expected 1 byte for Heartbeat messages.

## Error Design

The structure for the **Error** type is:

```
| type (uint8) | size (uint32) | payload |
```

- **type (uint8)**: 1 byte identifying the Error type.
- **size (uint32)**: 4 bytes specifying the length of the payload in bytes.
- **payload**: A string describing the error.

Unlike the Message type, Error omits headers since nodes expect the payload to be a string, simplifying the design.

## Payload Interface

All message types (Heartbeat, Message, Error, and future additions) must implement the `Payload` interface, defined in Go as:

```go
type Payload interface {
    fmt.Stringer       // Provides a string representation of the payload
    io.WriterTo        // Writes the payload to an io.Writer
    io.ReaderFrom      // Reads the payload from an io.Reader
    Byte() []byte      // Returns the payload as a byte slice
}
```

This interface ensures consistency in handling different message types across encoding and decoding operations.

## Package API Design

The protocol package exports two core functions for working with the protocol:

1. **`Encode(w io.Writer, payload Payload, payloadType uint8) (int64, error)`**
   - **Purpose**: Encodes a `Payload` into a writer, prefixed with the message type header.
   - **Returns**: The number of bytes written and any error encountered.
   - **Behavior**: If `payloadType` is unrecognized (e.g., not assigned to a known type), it returns an error.

2. **`Decode(r io.Reader) (Payload, error)`**
   - **Purpose**: Reads a message from a reader, identifies its type via the header, and decodes it into a `Payload`.
   - **Returns**: The decoded `Payload` and any error encountered.
   - **Behavior**: If the message type is unknown, it returns an error.

These functions provide a simple, unified API for encoding and decoding all supported message types.

---

### Additional Notes
- **Extensibility**: The use of a `uint8` type identifier (0–255) allows for up to 256 distinct message types, providing room for future expansion.
- **Security**: The Heartbeat DoS risk should be addressed in implementation, perhaps by enforcing a strict 1-byte read for this type.
- **Clarity**: The protocol assumes a reliable TCP connection; additional error handling (e.g., for partial reads) may need to be specified depending on use case.
