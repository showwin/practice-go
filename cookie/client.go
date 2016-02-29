package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	var cookies []*http.Cookie
	v := url.Values{"email": {"ishocon@isho.con"}}
	//v.Add("email", "ishocon@isho.con")
	v.Add("password", "ishoconpass")
	req, _ := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBufferString(v.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	jar, _ := cookiejar.New(nil)
	CookieURL, _ := url.Parse("http://localhost:8080/login")
	jar.SetCookies(CookieURL, cookies)
	client := http.Client{Jar: jar}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Print("responce")
	log.Print(resp)
	log.Print(jar.Cookies(CookieURL))
}
