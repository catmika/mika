package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Where is goddamn light?")
}
