package auth

import (
	"encoding/json"
	"github.com/VladRomanciuc/CSProject/auth/structs"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/spf13/viper"
)

	
// ConfigGoogle to set config of oauth
func ConfigGoogle() *oauth2.Config {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()
	conf := &oauth2.Config{
		ClientID:     viper.GetString("ClientID"),
		ClientSecret: viper.GetString("ClientSecret"),
		RedirectURL:  "http://127.0.0.1:8080/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email"}, // you can use other scopes to get more data
		Endpoint: google.Endpoint,
	}
	return conf
}

// GetEmail of user
func GetEmail(token string) string {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken}},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)

	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var data structs.GoogleResponse
	errorz := json.Unmarshal(body, &data)
	if errorz != nil {

		panic(errorz)
	}
	return data.Email
}