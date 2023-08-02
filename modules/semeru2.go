package modules

import (
	"encoding/json"
	"errors"
	"github.com/purawaktra/raung1-go/utils"
	"net/http"
	"strings"
)

type Semeru2Instance struct {
	Url          string
	AuthUsername string
	AuthPassword string
}

func CreateSemeru2Instance(url string, authUsername string, authPassword string) Semeru2Instance {
	return Semeru2Instance{
		Url:          url,
		AuthUsername: authUsername,
		AuthPassword: authPassword,
	}
}

func (s2 Semeru2Instance) InsertCredential(accountId int, email string, password string) (error, string) {
	// create data for body request
	data := struct {
		AccountId    int    `json:"account_id"`
		EmailAddress string `json:"email_address"`
		Password     string `json:"password"`
	}{
		AccountId:    accountId,
		EmailAddress: email,
		Password:     password,
	}

	// convert data to json structure and check err
	jsonData, err := json.Marshal(data)
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		return err, "FJ"
	}

	// create io reader
	reader := strings.NewReader(string(jsonData))

	// call http to make request
	url := utils.Semeru2Url + "/insert/credential"
	request, err, code := utils.CreateHTPPRequest(
		url,
		"POST",
		utils.Semeru2AuthUsername,
		utils.Semeru2AuthPassword,
		reader)
	if err != nil {
		utils.Error(err, "InsertCredential", "")
		return err, code
	}

	if request.StatusCode != http.StatusOK {
		utils.Warn("InsertCredential", "response status code not OK")
		return errors.New("response status code not OK"), "FC"
	}

	// create return
	return nil, "00"
}
