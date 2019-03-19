package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Auth struct {
	AccountSid string
	AuthToken  string
	Number     string
}

func CallNumber(number string) string {

	// Initialize twilio client
	_ = getCredentials()
	// client := twilio.NewClient(auth.AccountSid, auth.AuthToken, nil)

	// var callURL, _ = url.Parse("https://kev.inburke.com/zombo/zombocom.mp3")

	// _, err = client.Calls.MakeCall(auth.Number, number,
	// 	callURL)

	// if err != nil {
	// 	return err.Error()
	// }

	return fmt.Sprintf("Calling " + number)

}

func ExtractNumber(msg string) (string, error) {

	for _, x := range strings.Split(msg, " ") {
		if len(x) == 11 || len(x) == 12 {
			if strings.Contains(x, "+3725") {
				return x, nil
			}
		}
	}

	return "", errors.New(fmt.Sprintf("It was not Estonian cell number"))
}

func IsBadNumberInMessage(msg string) bool {

	if strings.Contains(msg, "+3725") {
		return true
	}

	return false
}

func getCredentials() *Auth {
	// Open jsonFile
	jsonFile, err := os.Open("conf.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	var auth Auth

	json.Unmarshal(byteValue, &auth)

	return &auth
}
