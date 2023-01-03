package fetch

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Fetch struct {
	Cookie []*http.Cookie
	client *http.Client
}

func (f *Fetch) InitUnSafe() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       time.Second * 15,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func (f *Fetch) Init() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       time.Second * 15,
	}
}
func (f *Fetch) SkipTlsCheck() {
	f.client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func (f *Fetch) Get(url string) ([]byte, error) {
	response, err := f.GetRaw(url)
	if err != nil {
		return nil, err
	}
	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (f *Fetch) GetRaw(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	response, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return response, err
}

func (f *Fetch) PostFormRaw(url string, requestData url.Values) (*http.Response, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	return f.client.Do(request)
}

func (f *Fetch) PostForm(url string, requestData url.Values) ([]byte, error) {
	response, err := f.PostFormRaw(url, requestData)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return ioutil.ReadAll(response.Body)
}

func (f *Fetch) PostJsonFormRaw(url string, requestData map[string]string) (*http.Response, error) {
	bytesData, _ := json.Marshal(requestData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	request.Header.Set("Content-Type", "application/json")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	return f.client.Do(request)
}

func (f *Fetch) PostJsonForm(url string, requestData map[string]string) ([]byte, error) {
	response, err := f.PostJsonFormRaw(url, requestData)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return ioutil.ReadAll(response.Body)
}

func cookieMerge(cookieA []*http.Cookie, cookieB []*http.Cookie) []*http.Cookie {

	for _, v := range cookieB {
		for k, v2 := range cookieA {
			if v.Name == v2.Name {
				cookieA = append(cookieA[:k], cookieA[k+1:]...)
				break
			}
		}
	}
	cookieA = append(cookieA, cookieB...)
	return cookieA
}
