package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

type CommandRequest struct {
	Command string   `json:"command"`
	Params  []string `json:"params"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	var cmdReq CommandRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cmdReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	cmd := exec.Command(cmdReq.Command, cmdReq.Params...)
	out, err := cmd.Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("Command execution failed: %s", err), http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(string(out))

	if err != nil {
		http.Error(w, fmt.Sprintf("Error happened in JSON marshal: %s", err), http.StatusInternalServerError)
		return
	}
	w.Write(marshal)
}
