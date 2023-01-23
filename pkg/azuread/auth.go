package azuread

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dfds/scim-setup/pkg/config"
	"github.com/dfds/scim-setup/pkg/logging"
)

const (
	// GraphScope The OAuth2 scope used for dealing with MS Graph API
	GraphScope string = "https://graph.microsoft.com/.default"
)

// GetBearerToken returns a token from Azure AD
func GetBearerToken() string {
	log := logging.GetLogger()
	address := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", config.TenantId())

	data := url.Values{
		"client_id":     {config.ClientId()},
		"client_secret": {config.ClientSecret()},
		"grant_type":    {"client_credentials"},
		"scope":         {GraphScope},
	}

	resp, err := http.PostForm(address, data)

	if err != nil {
		log.Fatal(err.Error())
	}

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Warn("The bearer the token is missing. Returning empty string.")
		log.Warn(err.Error())
		return ""
	}

	return fmt.Sprintf("Bearer %v", res["access_token"])
}
