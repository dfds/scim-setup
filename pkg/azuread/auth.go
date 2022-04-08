package azuread

import (
	"encoding/json"
	"fmt"
	"github.com/dfds/scim-setup/pkg/config"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

const (
	// GraphScope The OAuth2 scope used for dealing with MS Graph API
	GraphScope string = "https://graph.microsoft.com/.default"
)

// GetBearerToken returns a token from Azure AD
func GetBearerToken() string {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	sugar := logger.Sugar()
	address := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", config.TenantId())

	data := url.Values{
		"client_id":     {config.ClientId()},
		"client_secret": {config.ClientSecret()},
		"grant_type":    {"client_credentials"},
		"scope":         {GraphScope},
	}

	resp, err := http.PostForm(address, data)

	if err != nil {
		sugar.Fatal(err.Error())
	}

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		sugar.Warn("The bearer the token is missing. Returning empty string.")
		sugar.Warn(err.Error())
		return ""
	}

	return fmt.Sprintf("Bearer %v", res["access_token"])
}
