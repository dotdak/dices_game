package pkg

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
	Code  int         `json:"code"`
}

func (r *Response) Marshal() []byte {
	data, _ := json.Marshal(r)
	return data
}

func (r *Response) Write(rw http.ResponseWriter) {
	rw.WriteHeader(r.Code)
	_, _ = rw.Write(r.Marshal())
}
