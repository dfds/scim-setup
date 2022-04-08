package azuread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dfds/scim-setup/pkg/config"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"go.uber.org/zap"
	"net/http"
)

// GetGroupName takes a GraphServiceClient and a group object id, and return the name of an Azure AD Group
func GetGroupName(client *msgraphsdk.GraphServiceClient, groupId string) string {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	sugar := logger.Sugar()
	grp, err := client.GroupsById(groupId).Get(nil)
	if err != nil {
		sugar.Fatal(err.Error())
	}
	return *grp.GetDisplayName()
}

// AddMemberToGroup takes a group's object id and a user's (or service principal's) object id
func AddMemberToGroup(client *msgraphsdk.GraphServiceClient, groupId string, userId string) {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	sugar := logger.Sugar()
	groupUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/groups/%s/members/$ref", groupId)
	memberUrl := fmt.Sprintf("https://graph.microsoft.com/v1.0/directoryObjects/%s", userId)

	values := map[string]string{"@odata.id": memberUrl}
	payload, err := json.Marshal(values)
	if err != nil {
		sugar.Fatal(err.Error())
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", groupUrl, bytes.NewBuffer(payload))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", GetBearerToken())
	resp, err := httpClient.Do(req)

	if err != nil {
		sugar.Fatal(err.Error())
	}

	if resp.StatusCode == 401 {
		sugar.Fatal(resp.Status)
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		sugar.Errorf("The body response was: \n%v", resp.Body)
		sugar.Fatal(err.Error())
	}

	userName, err := GetUserName(client, userId)
	if err != nil {
		sugar.Error(err.Error())
	} else {
		sugar.Infof("Added %s to %s group", userName, GetGroupName(client, config.GroupId()))
	}
}
