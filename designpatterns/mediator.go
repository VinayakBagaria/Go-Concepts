/*
	Behavioural design pattern
	Reduce communication b/w objects and use a mediator for collaboration
*/

package designpatterns

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mediator interface {
	Get(url string) ([]byte, error)
	Post(url string, jSONPayload string) ([]byte, error)
}

// Implements Mediator interface
type HttpServiceCaller struct {
	httpClient *http.Client
}

func NewHttpServiceCaller() *HttpServiceCaller {
	return &HttpServiceCaller{httpClient: http.DefaultClient}
}

func (sc *HttpServiceCaller) Get(url string) ([]byte, error) {
	resp, err := sc.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (sc *HttpServiceCaller) Post(url string, jSONPayload string) ([]byte, error) {
	postData := bytes.NewBuffer([]byte(jSONPayload))
	resp, err := sc.httpClient.Post(url, "application/json", postData)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type UserServiceHandler struct {
	sc Mediator
}

type PostServiceHandler struct {
	sc Mediator
}

func (h *UserServiceHandler) GetAllUsers() {
	resp, err := h.sc.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%#v\n\n", string(resp))
	}
}

func (p *PostServiceHandler) SavePost(payload string) {
	resp, err := p.sc.Post("https://jsonplaceholder.typicode.com/posts", payload)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%#v\n\n", string(resp))
	}
}

func usingMediator() {
	httpServiceCaller := NewHttpServiceCaller()

	userServiceHandler := UserServiceHandler{sc: httpServiceCaller}
	postServiceHandler := PostServiceHandler{sc: httpServiceCaller}

	payload := `{"title": "btree.dev", "body": "mediator pattern", "userId": 1}`
	postServiceHandler.SavePost(payload)

	userServiceHandler.GetAllUsers()
}
