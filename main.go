package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"

	"github.com/dmlyons/contextIOTest/goContextIO"
)

func main() {
	key := flag.String("key", "", "Your CIO User Key")
	secret := flag.String("secret", "", "Your CIO User Secret")
	flag.Parse()
	c := contextio.NewContextIO(*key, *secret)
	params := url.Values{}
	params.Set("limit", "2")
	j, err := c.DoJson("GET", `/2.0/accounts/551420ac615a99de12fee488/messages/%3Ccomments%2F267458040%2Fcreated%40basecamp.com%3E`, params, nil)
	//j, err := c.DoJson("POST", `http://requestb.in/1ivgz6m1`, params, nil)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println(string(j))
	var out bytes.Buffer
	json.Indent(&out, j, "", "  ")
	fmt.Println(out.String())

	//	c := oauth.Client{
	//		Credentials: oauth.Credentials{
	//			Token:  Key,
	//			Secret: Secret,
	//		},
	//	}
	//	var params url.Values
	//	req, err := http.NewRequest("GET", `https://api.context.io/lite/users/54f77600facadd646442ad4e/email_accounts/0/folders/INBOX/messages/<CA+fiddEp63mdueHZ0JdTezsrCxt1EMh6-m=FTD1RkKk4dBCNYg@mail.gmail.com>/attachments`, nil)
	//	err = c.SetAuthorizationHeader(req.Header, nil, req.Method, req.URL, params)
	//
	//	response, err := http.DefaultClient.Do(req)
	//
	//	if err != nil {
	//		fmt.Println("ERROR:", err)
	//	}
	//	fmt.Println("Response:", response.StatusCode, response.Status)
	//	defer response.Body.Close()
	//	body, err := ioutil.ReadAll(response.Body)
	//	if err != nil {
	//		fmt.Println("ERROR:", err)
	//	}
	//	fmt.Println("BODY:", string(body))
}
