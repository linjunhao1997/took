package restful

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeQueryId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return 0, ErrBadRoute
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, ErrInvalidArgument
	}
	return id, nil
}
