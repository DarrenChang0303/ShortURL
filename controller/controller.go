package controller

import (
	"net/http"
)

// type RequestData struct {
// 	ID			string	'json:"id, omitempty"'
// 	Original	string	'json:"id, omitempty"'
// 	ID			string	'json:"id, omitempty"'
// }

type ResponseData struct {
	Original  string `json:"original, omitempty"`
	CreateURL string `json:"createURL, omitempty"`
	ID        string `json:"id, omitempty"`
	IsRequest bool   `json:"isRequest, omitempty"`
}

func CreateURL(w http.ResponseWriter, r *http.Request) {

}

func Original(w http.ResponseWriter, r *http.Request) {

}

func Redirect(w http.ResponseWriter, r *http.Request) {

}
