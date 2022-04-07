package azuread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dfds/scim-setup/pkg/config"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"log"
	"net/http"
)

// GetGroupName takes a GraphServiceClient and a group object id, and return the name of an Azure AD Group
func GetGroupName(client *msgraphsdk.GraphServiceClient, groupId string) string {
	grp, err := client.GroupsById(groupId).Get(nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return *grp.GetDisplayName()
}

// AddMemberToGroup takes a group's object id and a user's (or service principal's) object id
func AddMemberToGroup(client *msgraphsdk.GraphServiceClient, groupId string, userId string) {
	groupUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/groups/%s/members/$ref", groupId)
	memberUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/directoryObjects/%s", userId)

	values := map[string]string{"@odata.id": memberUrl}
	payload, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err.Error())
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", groupUrl, bytes.NewBuffer(payload))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", GetBearerToken())
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}

	if resp.StatusCode == 401 {
		log.Fatal(resp.Status)
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Added %s to %s group.\n", GetUserName(client, userId), GetGroupName(client, config.GroupId()))
}
