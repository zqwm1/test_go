package handler

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")

	//var buf bytes.Buffer
	//cmd.Stdin = &buf
	cmd := exec.Command("df")
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(json.Marshal(out))
	}
}
