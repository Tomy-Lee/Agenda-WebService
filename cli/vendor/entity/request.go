package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	// ErrInternalServerError ..
	ErrInternalServerError = fmt.Errorf("Internal Server Error")
)

var (
	scheme = "http"
	host   = "localhost:8080"
)

func init() {
	if _, present := os.LookupEnv("MOCK"); present {
		SetServer("https://private-ef9a4e-mensu.apiary-mock.com")
	}
}

// SetServer ..
func SetServer(serverURL string) (err error) {
	server, err := url.Parse(serverURL)
	if err != nil {
		return err
	}
	scheme = server.Scheme
	host = server.Host
	return nil
}

func request(method string, api string, reqBodyPtr interface{}, resBodyPtr interface{}) (code int, err error) {
	var reqBodyReader io.Reader
	if reqBodyPtr != nil {
		logger.Println("[request] preparing request body")
		var byteBody []byte
		if byteBody, err = json.Marshal(reqBodyPtr); err != nil {
			panic(err)
		}
		reqBodyReader = bytes.NewReader(byteBody)
	}

	logger.Println("[request] building full url")
	// build full url
	var fullURL *url.URL
	if fullURL, err = url.Parse(api); err != nil {
		panic(err)
	}
	fullURL.Scheme = scheme
	fullURL.Host = host

	logger.Println("[request] newing request")

	var req *http.Request
	if req, err = http.NewRequest(method, fullURL.String(), reqBodyReader); err != nil {
		panic(err)
	}

	openid := CurSessionModel.GetCurOpenid()
	if len(openid) > 0 {
		logger.Println("[request] adding openid cookie")
		// add openid cookie
		openidCookie := http.Cookie{
			Name:  "openid",
			Value: CurSessionModel.GetCurOpenid(),
		}
		req.AddCookie(&openidCookie)
	}

	logger.Println("[request] performing request with method =", method, ", url =", fullURL.String(), ", cookie =", req.Cookies())

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	code = res.StatusCode
	logger.Println("[request] responded with code =", code)

	if resBodyPtr != nil {
		logger.Println("[request] preparing response body")
		err = json.NewDecoder(res.Body).Decode(resBodyPtr)
	}
	if code >= http.StatusInternalServerError {
		err = ErrInternalServerError
	}
	logger.Println("[request] request done, err =", err)
	return
}
