package main

import (
	"fmt"

	"github.com/derotune/tenant-auth-service/config"
	"github.com/derotune/tenant-auth-service/oauth"
	"github.com/derotune/tenant-auth-service/secretManager"
	"golang.org/x/oauth2"

	"net/http"
)

var oauthConfig *oauth2.Config

func handleAuthRequest(w http.ResponseWriter, r *http.Request) {

	// returnUrl is what we need to get the provider based on the host and
	// redirect the user back to its original page
	returnUrl := r.URL.Query().Get("returnUrl")
	if returnUrl == "" {
		// throw 404
		http.Error(w, "No returnUrl given!", http.StatusUnprocessableEntity)
		return
	}
	// oauthConfig, err :=
	// try to find config for given url
	tenantConfig := config.FindMatchingConfiguration(returnUrl)

	if tenantConfig == nil {
		http.Error(w, "No Configuration found for the given returnUrl. ErrorCode: 50590", http.StatusMethodNotAllowed)
		return
	}

	// lets try to find ClientId and ClientSecret for the matchedConfiguration
	clientId := secretManager.GetSecret(tenantConfig.ClientId)
	clientSecret := secretManager.GetSecret(tenantConfig.ClientSecret)

	if clientId == "" {
		http.Error(w, "No Configuration found for the given returnUrl. ErrorCode: 50591", http.StatusNotFound)
		return
	}

	if clientSecret == "" {
		http.Error(w, "No Configuration found for the given returnUrl. ErrorCode: 50592", http.StatusNotFound)
		return
	}

	// get the oauth config object
	oauthConfig = oauth.BuildProviderConfig(oauth.ProviderOptions{
		Provider:     tenantConfig.Provider.Name,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Issuer:       tenantConfig.Provider.Issuer,
	})

	if oauthConfig == nil {
		http.Error(w, "Oauth Config object could not be created. ErrorCode: 50593", http.StatusInternalServerError)
	}

	// redirect to the provider AuthCodeUrl using the original url as state
	http.Redirect(w, r, oauthConfig.AuthCodeURL(returnUrl), http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	// time to make the token exchange with the provider and
	token, err := oauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))

	if err != nil {
		http.Error(w, fmt.Sprintf("Code exchange failed: %s. ErrorCode: 50594", err.Error()), http.StatusUnauthorized)
		return
	}

	// we got the token. So lets print it here.
	// In a real environment we wouldn't print it here of course, we would redirect to the initial called url from the user.
	fmt.Fprintln(w, "Access Token: "+token.AccessToken)
	fmt.Fprintln(w, "Original url: "+r.FormValue("state"))
}

func main() {

	http.HandleFunc("/", handleAuthRequest)
	http.HandleFunc("/callback", handleCallback)
	// http.HandleFunc("/exchange-token", handleExchangeToken)
	fmt.Println(http.ListenAndServe(":80", nil))

}
