package proxy

import (
	"reflect"
	"testing"
)

func TestProxy(t *testing.T) {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"

	t.Run("case: first time get request", func(t *testing.T) {
		response := nginxServer.handleRequest(appStatusURL, "GET")
		expectedResponse := Response{
			url:      appStatusURL,
			httpCode: 200,
			body:     "Ok",
		}
		if !(reflect.DeepEqual(expectedResponse, response)) {
			t.Fatalf("want: %+v\ngot: %+v", expectedResponse, response)
		}
	})

	t.Run("case: second time get request", func(t *testing.T) {
		response := nginxServer.handleRequest(appStatusURL, "GET")
		expectedResponse := Response{
			url:      appStatusURL,
			httpCode: 403,
			body:     "Not Allowed",
		}
		if !(reflect.DeepEqual(expectedResponse, response)) {
			t.Fatalf("want: %+v\ngot: %+v", expectedResponse, response)
		}
	})
}
