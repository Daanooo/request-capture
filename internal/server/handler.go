package server

import (
	"fmt"
	"io"
	"net/http"
)

func HandleRequest(r *http.Request) error {
	defer r.Body.Close()

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(buf))

	return nil
}
