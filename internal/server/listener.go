package server

import (
	"log"
	"net/http"
)

type Listener struct {
	closed chan error
	host   string
}

func NewListener(host string, closed chan error) *Listener {
	return &Listener{
		closed: closed,
		host:   host,
	}
}

func (l *Listener) Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := HandleRequest(r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusNoContent)
	})

	log.Printf("Server starting on %s\n", l.host)
	if err := http.ListenAndServe(l.host, nil); err != nil {
		l.closed <- err
	}
}

func (l *Listener) Closed() chan error {
	return l.closed
}
