package azuread

import (
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"go.uber.org/zap"
)

// GetUser returns the ObjectId, DisplayName and Email address of an Azure AD User
func GetUser(client *msgraphsdk.GraphServiceClient, user string) (string, string, string) {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	sugar := logger.Sugar()
	usr, err := client.UsersById(user).Get(nil)
	if err != nil {
		sugar.Fatal(err.Error())
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
