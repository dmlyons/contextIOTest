package contextio

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/garyburd/go-oauth/oauth"
)

const (
	GetUsersEndPoint       = "https://api.context.io/lite/users"
	GetAttachmentsEndPoint = "https://api.context.io/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/attachments/%s"
)

type User struct {
	Created        int            `json:"created"`
	EmailAccounts  []EmailAccount `json:"email_accounts"`
	EmailAddresses []string       `json:"email_addresses"`
	FirstName      string         `json:"first_name"`
	ID             string         `json:"id"`
	LastName       string         `json:"last_name"`
	ResourceURL    string         `json:"resource_url"`
}

type EmailAccount struct {
	AuthenticationType string `json:"authentication_type"`
	Label              string `json:"label"`
	Port               int    `json:"port"`
	ResourceURL        string `json:"resource_url"`
	Server             string `json:"server"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	UseSsl             bool   `json:"use_ssl"`
	Username           string `json:"username"`
}

type Message struct {
	Addresses      EmailAddresses `json:"addresses"`
	Attachments    []Attachment   `json:"attachments"`
	Bodies         []Body         `json:"bodies"`
	EmailMessageID string         `json:"email_message_id"`
	Folders        []string       `json:"folders"`
	InReplyTo      interface{}    `json:"in_reply_to"`
	ListHeaders    struct {
		List_Unsubscribe string `json:"list-unsubscribe"`
	} `json:"list_headers"`
	MessageID  string `json:"message_id"`
	PersonInfo struct {
		Bbcnews_Email_Bbc_Com struct {
			Thumbnail string `json:"thumbnail"`
		} `json:"bbcnews@email.bbc.com"`
		Dmlreturnpath_Gmail_Com struct {
			Thumbnail string `json:"thumbnail"`
		} `json:"dmlreturnpath@gmail.com"`
	} `json:"person_info"`
	ReceivedHeaders []string      `json:"received_headers"`
	References      []interface{} `json:"references"`
	ResourceURL     string        `json:"resource_url"`
	SentAt          float64       `json:"sent_at"`
	Subject         string        `json:"subject"`
}

type Attachment struct {
	AttachmentID       float64 `json:"attachment_id"`
	BodySection        string  `json:"body_section"`
	ContentDisposition string  `json:"content_disposition"`
	FileName           string  `json:"file_name"`
	Size               float64 `json:"size"`
	Type               string  `json:"type"`
}
type Body struct {
	BodySection string  `json:"body_section"`
	Size        float64 `json:"size"`
	Type        string  `json:"type"`
}

type EmailAddresses struct {
	Bcc     []EmailAddress `json:"bcc"`
	Cc      []EmailAddress `json:"cc"`
	From    []EmailAddress `json:"from"`
	ReplyTo []EmailAddress `json:"reply_to"`
	Sender  []EmailAddress `json:"sender"`
	To      []EmailAddress `json:"to"`
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ContextIO struct {
	key    string
	secret string
	client *oauth.Client
}

func NewContextIO(key, secret string) *ContextIO {
	c := &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  key,
			Secret: secret,
		},
	}

	return &ContextIO{
		key:    key,
		secret: secret,
		client: c,
	}
}

const (
	apiHost = `api.context.io`
)

// returns an *http.Response, the body must be defer response.Body.close()
func (c *ContextIO) Do(method, q string, params url.Values, body io.Reader) (response *http.Response, err error) {
	// Cannot use http.NewRequest because of the possibility of encoded data in the url
	req := &http.Request{
		Method: method,
		Host:   apiHost, // takes precendence over Request.URL.Host
		URL: &url.URL{
			Host:     apiHost,
			Scheme:   "https",
			Opaque:   q,
			RawQuery: params.Encode(),
		},
		Header: http.Header{
			"User-Agent": {"GoContextIO Simple library"},
		},
	}

	fmt.Print("req.URL ")
	fmt.Println(req.URL)
	fmt.Print("req.URL.Opaque ")
	fmt.Println(req.URL.Opaque)
	if err != nil {
		return
	}
	err = c.client.SetAuthorizationHeader(req.Header, nil, req.Method, req.URL, nil)
	fmt.Println("HL:", req.Header)
	if err != nil {
		return
	}
	return http.DefaultClient.Do(req)
}

func (c *ContextIO) DoJson(method, u string, params url.Values, body io.Reader) (j []byte, err error) {
	response, err := c.Do(method, u, params, body)
	defer response.Body.Close()
	j, err = ioutil.ReadAll(response.Body)
	//	json = string(bytes)
	return j, err
}

//func (c *ContextIO) GetUsers(params map[string]string) (users []User, err error) {
//	if err != nil {
//		return
//	}
//	defer r.Body.Close()
//
//	d := json.NewDecoder(r.Body)
//	err = d.Decode(&users)
//	if err != nil {
//		return
//	}
//	return
//}
//
//func (c *ContextIO) GetAttachment() (f io.Reader, err error) {
//	return
//}
