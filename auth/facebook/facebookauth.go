package fauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/VladRomanciuc/CSProject/auth/structs"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// ConfigGoogle to set config of oauth
func ConfigFacebook() *oauth2.Config {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()
	conf := &oauth2.Config{
		ClientID:     viper.GetString("facebook_client_id"),
		ClientSecret: viper.GetString("facebook_client_secret"),
		RedirectURL:  "https://127.0.0.1:8080/auth/facebook/callback",	
		Endpoint: facebook.Endpoint,
		Scopes: []string{"public_profile","email"}, // you can use other scopes to get more data
	}
	return conf
}

// GetUser of user
func GetFacebookUser(token string) structs.CallbackResponse {
	reqURL, err := url.Parse("https://graph.facebook.com/me?access_token=" + token)
	if err != nil {
		panic(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Accept": {"application/json"},
			"Authorization": {ptoken}},
	}

	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data structs.CallbackResponse
	errr := json.Unmarshal(response, &data)
	if errr != nil {
		panic(errr)
	}
	return data
}