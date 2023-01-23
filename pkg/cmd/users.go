package cmd

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/dfds/scim-setup/pkg/azuread"
	"github.com/dfds/scim-setup/pkg/config"
	"github.com/dfds/scim-setup/pkg/filedata"
	"github.com/dfds/scim-setup/pkg/logging"
	msazureauth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

// LoadUsersIntoGroups reads user email addresses from a text file, then loads each user into an Azure AD Group
func LoadUsersIntoGroups(fileName string) {
	log := logging.GetLogger()
	cred, err := azidentity.NewClientSecretCredential(config.TenantId(), config.ClientId(),
		config.ClientSecret(), &azidentity.ClientSecretCredentialOptions{})

	if err != nil {
		log.Fatalf("Error creating credentials for MS Graph: %v", err.Error())
	}

	auth, err := msazureauth.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{azuread.GraphScope})
	if err != nil {
		log.Fatalf("Error creation authentication provider for MS Graph: %v", err.Error())
		return
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		log.Fatalf("Error creating adapter for MS Graph: %v", err.Error())
		return
	}

	client := msgraphsdk.NewGraphServiceClient(adapter)

	users := filedata.GetUsers(fileName)

	for _, user := range users {
		userId, err := azuread.GetUserId(client, user)
		if err != nil {
			log.Infof("User %s wasn't found in Azure AD", user)
		} else {
			azuread.AddMemberToGroup(client, config.GroupId(), userId)
		}
	}
}
