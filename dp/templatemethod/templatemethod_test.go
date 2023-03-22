package templatemethod

import "testing"

func TestTemplateMethod(t *testing.T) {
	t.Run("case amazon token", func(t *testing.T) {

		amazonOAuth := &Amazon{}
		o := OAuth{
			iOAuth: amazonOAuth,
		}

		msg, err := o.genAndLogin()

		if msg != "logged in" {
			t.Fatal("amazon access token is mismatched\n", "error: ", err)
		}
	})

	t.Run("case google token", func(t *testing.T) {

		googleOAuth := &Google{}
		o := OAuth{
			iOAuth: googleOAuth,
		}

		msg, err := o.genAndLogin()

		if msg != "logged in" {
			t.Fatal("google access token is mismatched\n", "error: ", err)
		}
	})
}
