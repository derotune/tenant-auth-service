package oauth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type ProviderOptions struct {
	ClientId     string
	ClientSecret string
	Provider     string
	Issuer       string
}

func BuildProviderConfig(providerOptions ProviderOptions) *oauth2.Config {

	switch providerOptions.Provider {
	case "google":
		return &oauth2.Config{
			RedirectURL:  "http://localhost/callback",
			ClientID:     providerOptions.ClientId,
			ClientSecret: providerOptions.ClientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
	case "oidc":
		ctx := context.Background()
		provider, err := oidc.NewProvider(ctx, providerOptions.Issuer)
		if err != nil {
			//todo: better error handling
			fmt.Println(err)
		}

		return &oauth2.Config{
			RedirectURL:  "http://localhost/callback",
			ClientID:     providerOptions.ClientId,
			ClientSecret: providerOptions.ClientSecret,
			Scopes:       []string{"openid"},
			Endpoint:     provider.Endpoint(),
		}
	}
	return nil
}
