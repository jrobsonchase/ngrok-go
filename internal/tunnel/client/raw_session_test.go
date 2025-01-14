package client

import (
	"testing"

	"github.com/nikolay-ngrok/log15"

	"golang.ngrok.com/ngrok/internal/muxado"
)

type dummyStream struct{}

func (d *dummyStream) Read(bs []byte) (int, error)  { return 0, nil }
func (d *dummyStream) Write(bs []byte) (int, error) { return 0, nil }
func (d *dummyStream) Close() error                 { return nil }

func TestRawSessionDoubleClose(t *testing.T) {
	r := NewRawSession(log15.New(), muxado.Client(&dummyStream{}, nil), nil, nil)

	// Verify that closing the session twice doesn't cause a panic
	r.Close()
	r.Close()
}
