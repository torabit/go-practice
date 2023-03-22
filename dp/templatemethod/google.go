package templatemethod

import "fmt"

type Google struct {
	OAuth
}

func (g *Google) genAccessToken() string {
	token := "google-access_token"
	return token
}

func (g *Google) login(token string) (string, error) {
	expectToken := "google-access_token"

	if token != expectToken {
		return "", fmt.Errorf("this is not google access token")
	}

	return "logged in", nil
}
