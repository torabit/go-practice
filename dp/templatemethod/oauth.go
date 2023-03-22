package templatemethod

type IOAuth interface {
	genAccessToken() string
	login(string) (string, error)
}

type OAuth struct {
	iOAuth IOAuth
}

func (o *OAuth) genAndLogin() (string, error) {
	token := o.iOAuth.genAccessToken()
	message, err := o.iOAuth.login(token)

	if err != nil {
		return "", err
	}

	return message, nil
}
