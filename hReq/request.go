package hReq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/flambra/helpers/hError"
)

type Request struct {
	Url           string
	ContentType   string
	Authorization string
	Header        map[string]string
	Body          interface{}
	Params        map[string]interface{}
	StatusCode    int
}

func (r *Request) Post() ([]byte, error) {
	if r.ContentType == "" {
		r.ContentType = "application/json"
	}

	payload, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", r.Url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", r.ContentType)
	r.customHeaders(req)

	if r.Authorization != "" {
		req.Header.Add("Authorization", r.Authorization)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	decoded, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r.StatusCode = resp.StatusCode

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := string(decoded)
		log.Println(message)
		return nil, hError.New(resp.Status)
	}

	return decoded, nil
}

func (r *Request) Get() ([]byte, error) {
	if r.ContentType == "" {
		r.ContentType = "application/json"
	}

	client := &http.Client{}

	url := params(r.Url, r.Params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", r.ContentType)
	r.customHeaders(req)

	if r.Authorization != "" {
		req.Header.Add("Authorization", r.Authorization)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	decoded, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r.StatusCode = resp.StatusCode

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := string(decoded)
		log.Println(message)
		return nil, hError.New(resp.Status)
	}

	return decoded, nil
}

func (r *Request) Put() ([]byte, error) {
	if r.ContentType == "" {
		r.ContentType = "application/json"
	}

	payload, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	url := params(r.Url, r.Params)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", r.ContentType)
	r.customHeaders(req)

	if r.Authorization != "" {
		req.Header.Add("Authorization", r.Authorization)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	decoded, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r.StatusCode = resp.StatusCode

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := string(decoded)
		log.Println(message)
		return nil, hError.New(resp.Status)
	}

	return decoded, nil
}

func params(baseURL string, params map[string]interface{}) string {
	if len(params) == 0 {
		return baseURL
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return baseURL
	}

	q := u.Query()
	for key, value := range params {
		switch v := value.(type) {
		case int:
			q.Set(key, fmt.Sprintf("%d", v))
		case uint:
			q.Set(key, fmt.Sprintf("%d", v))
		default:
			q.Set(key, fmt.Sprintf("%v", v)) // String
		}
	}
	u.RawQuery = q.Encode()

	return u.String()
}

func (r *Request) customHeaders(req *http.Request) {
	for key, value := range r.Header {
		req.Header.Add(key, value)
	}
}
