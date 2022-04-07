package azuread

import (
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"log"
)

// GetUser returns the ObjectId, DisplayName and Email address of an Azure AD User
func GetUser(client *msgraphsdk.GraphServiceClient, user string) (string, string, string) {
	usr, err := client.UsersById(user).Get(nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return *usr.GetId(), *usr.GetDisplayName(), *usr.GetMail()
}

// GetUserId returns the Object Id for an Azure AD user
func GetUserId(client *msgraphsdk.GraphServiceClient, user string) string {
	id, _, _ := GetUser(client, user)
	return id
}

// GetUserName returns the Display Name for an Azure AD user
func GetUserName(client *msgraphsdk.GraphServiceClient, user string) string {
	_, name, _ := GetUser(client, user)
	return name
}

// GetUserEmail returns the mail address for an Azure AD user
func GetUserEmail(client *msgraphsdk.GraphServiceClient, user string) string {
	_, _, email := GetUser(client, user)
	return email
}
