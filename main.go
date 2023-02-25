package main

import (
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/shutdown", func(_ http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			return
		}

		binary, lookErr := exec.LookPath("shutdown")
		if lookErr != nil {
			panic(lookErr)
		}

		err := syscall.Exec(binary, []string{"shutdown", "-P", "now"}, os.Environ())
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
