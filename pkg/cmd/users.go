package cmd

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/dfds/scim-setup/pkg/azuread"
	"github.com/dfds/scim-setup/pkg/config"
	"github.com/dfds/scim-setup/pkg/filedata"
	msazureauth "github.com/microsoft/kiota/authentication/go/azure"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"log"
)

// LoadUsersIntoGroups reads user email addresses from a text file, then loads each user into an Azure AD Group
func LoadUsersIntoGroups(fileName string) {
	cred, err := azidentity.NewClientSecretCredential(config.TenantId(), config.ClientId(),
		config.ClientSecret(), &azidentity.ClientSecretCredentialOptions{})

	if err != nil {
		log.Fatalf("Error creating credentials for MS Graph: %v\n", err.Error())
	}

	auth, err := msazureauth.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{azuread.GraphScope})
	if err != nil {
		log.Fatalf("Error creation authentication provider for MS Graph: %v\n", err.Error())
		return
	}

	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		log.Fatalf("Error creating adapter for MS Graph: %v\n", err.Error())
		return
	}

	client := msgraphsdk.NewGraphServiceClient(adapter)

	users := filedata.GetUsers(fileName)

	for _, user := range users {
		azuread.AddMemberToGroup(client, config.GroupId(), azuread.GetUserId(client, user))
	}
}
