package authentication

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/entity"
)

func meFromRemote(token string) (*entity.User, error) {
	authUrl := config.AuthUrl()
	jsonData := map[string]string{
		"query": `
				{ 
					me {
						id
						name
						loginName
						roles {
							id
							name
						}
					}
				}
		`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", authUrl, bytes.NewBuffer(jsonValue))
	request.Header.Set(consts.AUTHORIZATION, consts.BEARER+token)
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, errors.New("Can't access authentication service: " + authUrl)
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	var user entity.User
	var userJson map[string]interface{}
	json.Unmarshal(data, &userJson)
	fmt.Println(userJson)
	if userJson["data"] != nil {
		meJson := userJson["data"].(map[string]interface{})["me"]
		if meJson != nil {
			err = mapstructure.Decode(meJson, &user)
			if err != nil {
				panic(err.Error())
			}
		}
		return &user, nil
	}
	return nil, nil
}
