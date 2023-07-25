package bom

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"sync"
)

var (
	mutex   = &sync.RWMutex{}
	clients = map[string]*Client{}
)

type Client struct {
	Rest *resty.Client
}

func New(cli *Client) *Client {
	mutex.Lock()
	defer mutex.Unlock()
	if cli == nil {
		log.Fatal("rubix-os client cli can not be empty")
		return nil
	}
	baseURL := "https://api.weather.bom.gov.au/v1"
	if client, found := clients[baseURL]; found {
		return client
	}
	cli.Rest = resty.New()
	cli.Rest.SetBaseURL(baseURL)
	clients[baseURL] = cli
	return cli
}

// An Error maps to Form3 API error responses
type Error struct {
	Code    int    `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

type Message struct {
	Message string `json:"message"`
}

type FoundMessage struct {
	Found bool `json:"found"`
}

// composeErrorMsg it helps to create a clean output error message; we used to have JSON message with nested key
func composeErrorMsg(resp *resty.Response) error {
	message := Message{}
	rawMessage := resp.String()
	_ = json.Unmarshal([]byte(rawMessage), &message)

	if message.Message == "" {
		// if we do not have => `{"message": <message>}`
		message.Message = fmt.Sprintf("%s %s [%d]: %s",
			resp.Request.Method,
			resp.Request.URL,
			resp.StatusCode(),
			rawMessage)
	} else if message.Message == "not found" {
		message.Message = fmt.Sprintf("%s %s [%d]: %s",
			resp.Request.Method,
			resp.Request.URL,
			resp.StatusCode(),
			message.Message)
	}
	e := fmt.Errorf(message.Message)
	return e
}

func FormatRestyResponse(resp *resty.Response, err error) (*resty.Response, error) {
	if err != nil {
		return resp, err
	}
	if resp.IsError() {
		return resp, composeErrorMsg(resp)
	}
	return resp, nil
}
