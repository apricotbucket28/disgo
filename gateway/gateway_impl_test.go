package gateway

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/disgoorg/json"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

// Messages encoded in either zlib or zstd using streaming contexts.
// HeartbeatACK is {"t":null,"s":null,"op":11,"d":null} while
// InvalidSession is {"t":null,"s":null,"op":9,"d":false}.
var (
	zlibHeartbeatACK = [][]byte{
		{120, 156, 170, 86, 42, 81},
		{178, 202, 43, 205, 201, 209, 81, 42},
		{134, 49, 242, 11, 148, 172, 12, 13},
		{117, 148, 82, 32, 2, 181, 0, 0, 0, 0, 255, 255},
	}
	zlibInvalidSession = []byte{
		194, 165, 198, 18, 172, 36, 45, 49, 167, 56, 181, 22, 0, 0, 0, 255, 255,
	}

	zstdHeartbeatACK = []byte{
		40, 181, 47, 253, 4, 104, 32, 1, 0, 123, 34, 116, 34, 58, 110, 117, 108,
		108, 44, 34, 115, 34, 58, 110, 117, 108, 108, 44, 34, 111, 112, 34, 58,
		49, 49, 44, 34, 100, 34, 58, 110, 117, 108, 108, 125,
	}
	zstdInvalidSession = []byte{
		156, 0, 0, 96, 57, 44, 34, 100, 34, 58, 102, 97, 108, 115, 101, 125, 1, 84, 0, 5, 21, 39,
	}
)

func heartbeatACK() *Message {
	return &Message{
		Op:   OpcodeHeartbeatACK,
		RawD: json.RawMessage([]byte(`null`)),
	}
}

func invalidSession() *Message {
	return &Message{
		Op:   OpcodeInvalidSession,
		D:    MessageDataInvalidSession(false),
		RawD: json.RawMessage([]byte(`false`)),
	}
}

func TestParseMessageZlib(t *testing.T) {
	g := gatewayImpl{
		config: Config{
			Logger:      slog.Default(),
			Compression: CompressionZlib,
		},
	}

	for i, chunk := range zlibHeartbeatACK {
		message, err := g.parseMessage(websocket.BinaryMessage, bytes.NewReader(chunk))

		assert.NoError(t, err)

		last := i == len(zlibHeartbeatACK)-1
		if last {
			assert.Equal(t, heartbeatACK(), message)
		} else {
			assert.Nil(t, message)
		}
	}

	message, err := g.parseMessage(websocket.BinaryMessage, bytes.NewReader(zlibInvalidSession))
	assert.NoError(t, err)
	assert.Equal(t, invalidSession(), message)
}

func TestParseMessageZstd(t *testing.T) {
	g := gatewayImpl{
		config: Config{
			Logger:      slog.Default(),
			Compression: CompressionZstd,
		},
	}

	message, err := g.parseMessage(websocket.BinaryMessage, bytes.NewReader(zstdHeartbeatACK))
	assert.NoError(t, err)
	assert.Equal(t, heartbeatACK(), message)

	message, err = g.parseMessage(websocket.BinaryMessage, bytes.NewReader(zstdInvalidSession))
	assert.NoError(t, err)
	assert.Equal(t, invalidSession(), message)
}
