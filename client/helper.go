package client

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// user account
type account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// response from login
type loginResp struct {
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

func readBody(resp *http.Response) []byte {
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	return data
}

func getAccount() account {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSuffix(username, "\n")
	fmt.Print("Enter your password: ")
	passwordmid, _ := terminal.ReadPassword(0)
	password := strings.TrimSuffix(string(passwordmid[:]), "\n")
	fmt.Print("\n")
	return account{username, password}
}

func check(err error) {
	if err != nil {
		// fmt.Print(err)
	}
}
