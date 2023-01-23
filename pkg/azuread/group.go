package azuread

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dfds/scim-setup/pkg/config"
	"github.com/dfds/scim-setup/pkg/logging"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

// GetGroupName takes a GraphServiceClient and a group object id, and return the name of an Azure AD Group
func GetGroupName(client *msgraphsdk.GraphServiceClient, groupId string) string {
	log := logging.GetLogger()
	grp, err := client.GroupsById(groupId).Get(context.Background(), nil)
	if err != nil {
		log.Error(err.Error())
	}
	return *grp.GetDisplayName()
}

// AddMemberToGroup takes a group's object id and a user's (or service principal's) object id
func AddMemberToGroup(client *msgraphsdk.GraphServiceClient, groupId string, userId string) {
	log := logging.GetLogger()
	groupUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/groups/%s/members/$ref", groupId)
	memberUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/directoryObjects/%s", userId)

	values := map[string]string{"@odata.id": memberUrl}
	payload, err := json.Marshal(values)
	if err != nil {
		log.Error(err.Error())
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", groupUrl, bytes.NewBuffer(payload))
	if err != nil {
		log.Error(err)
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", GetBearerToken())
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Error(err.Error())
	}

	if resp.StatusCode == 401 {
		log.Error(resp.Status)
	}

	userName, err := GetUserName(client, userId)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("Added %s to %s group", userName, GetGroupName(client, config.GroupId()))
	}
}
