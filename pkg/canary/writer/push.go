package writer

import (
	"io"
	"sync"
)

var p io.WriteCloser = &Push{}

// Push is a io.Writer, that writes given long entries by pushing
// directly to the given loki server URL. Each `Push` instance handles single tenant.
// TODO(kavi): Add batching?
type Push struct {
	lokiURL  string
	tenantID string

	mu sync.RWMutex

	// batch represents a single batch of log entries that will be sent to loki at once.
	batch []string

	// state to identify that Push is sending it's last batch (if any)
	// to loki so that no more write will be accepted
	closing bool
}

func (p *Push) Write(buf []byte) (int, error) {

}

func (p *Push) sendBatch() (int, error) {

}

// Close makes sure the pending buffer is pushed to `loki` before
// returning. It's the responsibility of the client to call Close.
func (p *Push) Close() error {

}
