package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"
)

type handler interface {
	Handle(resp *http.Response, err error, f func(*http.Response, error))
	User(resp *http.Response, err error)
	Signup(resp *http.Response, err error)
	Login(resp *http.Response, err error)
	Logout(resp *http.Response, err error)
}

// Handler to handle http response
type Handler struct{}

// Handle handles a response by specific handle function
func (h *Handler) Handle(resp *http.Response, err error, f func(*http.Response, error)) {
	f(resp, err)
}

// User handles the response from '/user' api
func (h *Handler) User(resp *http.Response, err error) {
	check(err)
	readBody(resp)
}

// Login handles the response from '/login' api
func (h *Handler) Login(resp *http.Response, err error) {
	check(err)
	body := readBody(resp)
	var data loginResp
	err = json.Unmarshal(body, &data)
	check(err)
	syscall.Umask(0)
	err = os.MkdirAll("/tmp/ppe", 0755)
	check(err)
	f, err := os.Create("/tmp/ppe/token")
	check(err)
	token := data.Token
	n3, _ := f.WriteString(token)
	fmt.Printf("wrote %d bytes\n", n3)
}

// Logout handles the response from '/logout' api
func (h *Handler) Logout(resp *http.Response, err error) {
	check(err)
	readBody(resp)
	err = os.Remove("/tmp/ppe/token")
	check(err)
}

// Signup handles the response from '/signup' api
func (h *Handler) Signup(resp *http.Response, err error) {
	check(err)
	readBody(resp)
}

// Deploy handles the response from '/deploy' api
func (h *Handler) Deploy(resp *http.Response, err error) {
	check(err)
	readBody(resp)
}
