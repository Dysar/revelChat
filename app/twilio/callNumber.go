package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"github.com/kevinburke/twilio-go"
)

type Auth struct {
	AccountSid string
	AuthToken  string
	Number     string
}

func CallNumber(msg string) string {

	number, err := extractNumber(msg)

	if err != nil {
		return err.Error()
	}

	// Initialize twilio client
	auth := getCredentials()
	client := twilio.NewClient(auth.AccountSid, auth.AuthToken, nil)

	var callURL, _ = url.Parse("https://kev.inburke.com/zombo/zombocom.mp3")

	_, err = client.Calls.MakeCall(auth.Number, number,
		callURL)

	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("Calling " + number)

}

func extractNumber(msg string) (string, error) {

	for _, x := range strings.Split(msg, " ") {
		if len(x) == 11 || len(x) == 12 {
			if strings.Contains(x, "+3725") {
				return x, nil
			}
		}
	}

	return "", errors.New("I accept only Estonian cell numbers for now")
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
