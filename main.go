package main

import (
	"net/http"
	"syscall"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/shutdown", func(_ http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			return
		}

		err := syscall.Shutdown(syscall.SYS_SHUTDOWN, 0)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
