package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	BUF_LEN = 1024
)

func handler(w http.ResponseWriter, r *http.Request) {
	content := readFile(r.URL.Query().Get("dir"))
	log.Println("got content", content)

	cmd := exec.Command("ls", "-a")
	pipeReader, pipeWriter := io.Pipe()
	cmd.Stdout = pipeWriter
	cmd.Stderr = pipeWriter
	go writeCmdOutput(w, pipeReader)
	cmd.Run()
	pipeWriter.Close()
}

func writeCmdOutput(res http.ResponseWriter, pipeReader *io.PipeReader) {
	buffer := make([]byte, BUF_LEN)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			pipeReader.Close()
			break
		}

		data := buffer[0:n]
		res.Write(data)
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
		//reset buffer
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}

func readFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(data))
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("listening to port *:8080. press ctrl + c to cancel.")
	http.ListenAndServe(":8080", nil)
}
