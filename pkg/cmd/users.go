package cmd

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/dfds/scim-setup/pkg/azuread"
	"github.com/dfds/scim-setup/pkg/config"
	"github.com/dfds/scim-setup/pkg/filedata"
	msazureauth "github.com/microsoft/kiota/authentication/go/azure"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"go.uber.org/zap"
)

// LoadUsersIntoGroups reads user email addresses from a text file, then loads each user into an Azure AD Group
func LoadUsersIntoGroups(fileName string) {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	sugar := logger.Sugar()
	cred, err := azidentity.NewClientSecretCredential(config.TenantId(), config.ClientId(),
		config.ClientSecret(), &azidentity.ClientSecretCredentialOptions{})

	if err != nil {
		sugar.Fatalf("Error creating credentials for MS Graph: %v", err.Error())
	}

	auth, err := msazureauth.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{azuread.GraphScope})
	if err != nil {
		sugar.Fatalf("Error creation authentication provider for MS Graph: %v", err.Error())
		return
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		sugar.Fatalf("Error creating adapter for MS Graph: %v", err.Error())
		return
	}

	client := msgraphsdk.NewGraphServiceClient(adapter)

	users := filedata.GetUsers(fileName)

	for _, user := range users {
		azuread.AddMemberToGroup(client, config.GroupId(), azuread.GetUserId(client, user))
	}
}
