package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"

	contextio "github.com/dmlyons/contextIOTest/goContextIO"
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
	j, err := c.DoJson("GET", q, params, nil)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	var out bytes.Buffer
	json.Indent(&out, j, "", "  ")
	fmt.Println(out.String())
}
