package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type HttpRequest struct {
	Url           string
	ContentType   string
	Authorization string
	Body          interface{}
	StatusCode    int
}

func (h *HttpRequest) Post() ([]byte, error) {
	if h.ContentType == "" {
		h.ContentType = "application/json"
	}

	payload, err := json.Marshal(&h.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", h.Url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", h.ContentType)

	if h.Authorization != "" {
		req.Header.Add("Authorization", h.Authorization)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	decoded, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	h.StatusCode = resp.StatusCode

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := string(decoded)
		log.Println(message)
		return nil, errors.New(resp.Status)

	}

	return decoded, nil
}

func (h *HttpRequest) Get() ([]byte, error) {
	if h.ContentType == "" {
		h.ContentType = "application/json"
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", h.Url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", h.ContentType)

	if h.Authorization != "" {
		req.Header.Add("Authorization", h.Authorization)
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

	h.StatusCode = resp.StatusCode

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := string(decoded)
		log.Println(message)
		return nil, errors.New(resp.Status)

	}

	return decoded, nil
}

func (h *HttpRequest) Put() ([]byte, error) {
	if h.ContentType == "" {
		h.ContentType = "application/json"
	}

	payload, err := json.Marshal(&h.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("PUT", h.Url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", h.ContentType)

	if h.Authorization != "" {
		req.Header.Add("Authorization", h.Authorization)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	decoded, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	h.StatusCode = resp.StatusCode

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := string(decoded)
		log.Println(message)
		return nil, errors.New(resp.Status)

	}

	return decoded, nil
}
