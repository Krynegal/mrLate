package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (h *Handler) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := h.Oauth.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) CallBackFromGoogle(w http.ResponseWriter, r *http.Request) {
	log.Println("Callback-gl..")

	state := r.FormValue("state")
	log.Println(state)

	code := r.FormValue("code")
	log.Println(code)

	if code == "" {
		log.Println("Code not found..")
		w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User has denied Permission.."))
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := h.Oauth.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Println("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		log.Println("TOKEN>> AccessToken>> " + token.AccessToken)
		log.Println("TOKEN>> Expiration Time>> " + token.Expiry.String())
		log.Println("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://www.googleapis.com/calendar/v3/users/me/calendarList?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			log.Println("Get: " + err.Error() + "\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ReadAll: " + err.Error() + "\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		log.Println("parseResponseBody: " + string(response) + "\n")

		r, err := fastjson.ParseBytes(response)
		if err != nil {
			http.Error(w, "can't parse", http.StatusBadRequest)
		}
		fmt.Println(r.Get("items"))

		w.Write([]byte("Hello, I'm protected\n"))
		w.Write([]byte(string(response)))
		return
	}
}

func getResponseMap(responseBody []byte) map[string]interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		fmt.Println(err)
	}
	return result
}
