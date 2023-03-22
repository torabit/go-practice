package templatemethod

import "fmt"

type Amazon struct {
	OAuth
}

func (a *Amazon) genAccessToken() string {
	token := "amazon-access_token"
	return token
}

func (a *Amazon) login(token string) (string, error) {
	expectToken := "amazon-access_token"

	if token != expectToken {
		return "", fmt.Errorf("this is not amazon access token")
	}

	return "logged in", nil
}
