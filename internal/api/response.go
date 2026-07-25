package api

import (
	"encoding/json"
	"net/http"
)

// writeJSON serializa data para JSON e escreve no ResponseWriter
// com o status code e Content-Type definidos.
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

// writeError escreve uma resposta de erro em JSON com o status
// code e a mensagem informados.
func writeError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
