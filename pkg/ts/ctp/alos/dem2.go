package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gocolly/colly"
)

type AuthCookieJar struct {
	Path string
}

func NewAuthCookieJar(path string) *AuthCookieJar {
	return &AuthCookieJar{Path: path}
}

func (jar *AuthCookieJar) SaveCookies(cookies []*http.Cookie) error {
	cookieJarPath := filepath.Join(os.Getenv("HOME"), ".bulk_download_cookiejar.txt")
	cookieFile, err := os.OpenFile(cookieJarPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer cookieFile.Close()

	encodedCookies := ""
	for _, cookie := range cookies {
		encodedCookies += cookie.String() + "\n"
	}

	_, err = cookieFile.WriteString(encodedCookies)
	return err
}

func (jar *AuthCookieJar) GetNewCookie(asfUrs4 map[string]string) error {
	c := colly.NewCollector()
	c.Jar = new(http.Client).Jar

	// Basic auth credentials
	userPass := "mocheer:gulaxy@1107G"
	encodedUserPass := base64.StdEncoding.EncodeToString([]byte(userPass))

	authCookieURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&state=",
		asfUrs4["url"], asfUrs4["client"], asfUrs4["redir"])

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Authorization", "Basic "+encodedUserPass)
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 401 && r.TextContains("Please enter your Earthdata Login credentials") {
			log.Println("Authentication failed")
			return
		}

		if jar.checkCookieIsLoggedIn(c.Jar.Cookies(nil)) {
			if err := jar.SaveCookies(c.Jar.Cookies(nil)); err != nil {
				log.Println("Error saving cookies:", err)
				return
			}
			log.Println("Cookies saved successfully")
			return
		}

		log.Println("Failed to login")
	})

	// Visit the auth URL to trigger the login process
	c.Visit(authCookieURL)

	return nil
}

func (jar *AuthCookieJar) checkCookieIsLoggedIn(cookies []*http.Cookie) bool {
	for _, cookie := range cookies {
		if cookie.Name == "urs_user_already_logged" {
			return true
		}
	}
	return false
}

func main2() {
	asfUrs4 := map[string]string{
		"url":    "https://urs.earthdata.nasa.gov/oauth/authorize",
		"client": "BO_n7nTIlMljdvU6kRRB3g",
		"redir":  "https://auth.asf.alaska.edu/login",
	}

	jar := NewAuthCookieJar(".")
	if err := jar.GetNewCookie(asfUrs4); err != nil {
		log.Fatalf("Error getting new cookie: %v", err)
	}
}
