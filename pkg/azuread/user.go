package azuread

import (
	"context"

	"github.com/dfds/scim-setup/pkg/logging"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

// GetUser returns the ObjectId, DisplayName and Email address of an Azure AD User
func GetUser(client *msgraphsdk.GraphServiceClient, user string) (string, string, string, error) {
	log := logging.GetLogger()
	usr, err := client.UsersById(user).Get(context.Background(), nil)
	if err != nil {
		log.Error(err.Error())
		return "", "", "", err
	}
	return *usr.GetId(), *usr.GetDisplayName(), *usr.GetMail(), nil
}

// GetUserId returns the Object Id for an Azure AD user
func GetUserId(client *msgraphsdk.GraphServiceClient, user string) (string, error) {
	id, _, _, err := GetUser(client, user)
	if err != nil {
		return "", err
	}
	return id, err
}

// GetUserName returns the Display Name for an Azure AD user
func GetUserName(client *msgraphsdk.GraphServiceClient, user string) (string, error) {
	_, name, _, err := GetUser(client, user)
	if err != nil {
		return "", err
	}
	return name, nil
}

// GetUserEmail returns the mail address for an Azure AD user
func GetUserEmail(client *msgraphsdk.GraphServiceClient, user string) (string, error) {
	_, _, email, err := GetUser(client, user)
	if err != nil {
		return "", err
	}
	return email, nil
}
