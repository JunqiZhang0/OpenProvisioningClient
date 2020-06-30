package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/openprovclient/config"
)

type client struct {
	Client  *http.Client
	BaseURL string
	Handler Handler
}

// NewClient return a new client
func NewClient(url string) client {
	return client{&http.Client{}, url, Handler{}}
}

// AuthRequest return a request with a token in header
func (c *client) AuthRequest(method, url string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(url, "http") {
		url = c.BaseURL + url
	}
	token, err := ioutil.ReadFile("/tmp/ppe/token")
	if err != nil {
		fmt.Println("Auth failed(token not found), please login")
		return nil, err
	}
	req, err := http.NewRequest(method, url, body)
	check(err)
	req.Header.Add("token", string(token))
	return req, err
}

// HTTPRequest return a normal http request
func (c *client) HTTPRequest(method, url string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(url, "http") {
		url = c.BaseURL + url
	}
	return http.NewRequest(method, url, body)
}

// User send request to '/user' api and handle response
func (c *client) User() {
	req, err := c.AuthRequest("GET", "/user", nil)
	if err != nil {
		return
	}
	resp, err := c.Client.Do(req)
	c.Handler.Handle(resp, err, c.Handler.User)
}

// User send request to '/signup' api and handle response
func (c *client) Signup() {
	account := getAccount()
	jsonValue, err := json.Marshal(account)
	check(err)
	req, err := c.HTTPRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	check(err)
	resp, err := c.Client.Do(req)
	c.Handler.Handle(resp, err, c.Handler.Signup)
}

// Login send request to '/login' api and handle response
func (c *client) Login() {
	account := getAccount()
	jsonValue, err := json.Marshal(account)
	check(err)
	req, err := c.HTTPRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	// fmt.Printf("Login success!")
	// The line below should be left here for debug use
	// fmt.Printf("%v\n\n\n", string(jsonValue))
	check(err)
	resp, err := c.Client.Do(req)
	c.Handler.Handle(resp, err, c.Handler.Login)
}

// Logout send request to '/logout' api and handle response
func (c *client) Logout() {
	req, err := c.AuthRequest("POST", "/logout", nil)
	check(err)
	resp, err := c.Client.Do(req)
	c.Handler.Handle(resp, err, c.Handler.Logout)
}

//
func (c *client) Deploy(filename string) {
	config, err := config.LoadConfig(filename)
	check(err)
	jsonValue, err := json.Marshal(config)
	check(err)
	fmt.Println(string(jsonValue))
	req, err := c.AuthRequest("POST", "/deploy", bytes.NewBuffer(jsonValue))
	check(err)
	resp, err := c.Client.Do(req)
	c.Handler.Handle(resp, err, c.Handler.Deploy)
}
