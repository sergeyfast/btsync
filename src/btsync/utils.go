package btsync

import (
	"net/http"
	"net/url"
	"log"
)

type myjar struct {
	jar map[string] []*http.Cookie
}

var Debug bool = false

func (p* myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	if Debug {
		log.Printf("The URL is : %s\n", u.String())
		log.Printf("The cookie being set is : %s\n", cookies)
	}
	p.jar [u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
	if Debug {
		log.Printf("The URL is : %s\n", u.String())
		log.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
	}
	return p.jar[u.Host]
}
