package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")

	//var buf bytes.Buffer
	//cmd.Stdin = &buf
	cmd := exec.Command("ls")
	out, err := cmd.Output()
	marshal, err := json.Marshal(out)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(marshal)
	}
}
