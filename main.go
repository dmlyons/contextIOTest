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
	params.Set("include_body", "1")
	eid := "<comments/267458040/created@basecamp.com>"
	q := `/2.0/accounts/551420ac615a99de12fee488/messages/` + url.QueryEscape(eid)
	up, _ := url.Parse(q)
	fmt.Println("UP", up)
	j, err := c.DoJson("GET", q, params, nil)
	//j, err := c.DoJson("POST", `http://requestb.in/1ivgz6m1`, params, nil)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	//fmt.Println(string(j))
	var out bytes.Buffer
	json.Indent(&out, j, "", "  ")
	fmt.Println(out.String())

}
