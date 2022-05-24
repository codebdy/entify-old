package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"rxdrag.com/entify/config"
	"rxdrag.com/entify/consts"
	"rxdrag.com/entify/entity"
)

func meFromRemote(token string) *entity.User {
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
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

	return nil
}
