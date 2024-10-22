package server

import (
	"log"
	"net/http"
)

type Listener struct {
	captures chan Capture
	closed   chan error
	host     string
}

func NewListener(host string, captures chan Capture, closed chan error) *Listener {
	return &Listener{
		captures: captures,
		closed:   closed,
		host:     host,
	}
}

func (l *Listener) Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		capture, err := NewCapture(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusNoContent)

		l.captures <- capture
	})

	log.Printf("Server starting on %s\n", l.host)
	if err := http.ListenAndServe(l.host, nil); err != nil {
		l.closed <- err
	}
}
