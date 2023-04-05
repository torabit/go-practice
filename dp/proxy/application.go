package proxy

type Application struct{}

func (a *Application) handleRequest(url, method string) Response {
	response := &Response{
		url:      url,
		httpCode: 404,
		body:     "Not Ok",
	}

	if url == "/app/status" && method == "GET" {
		response.url = url
		response.httpCode = 200
		response.body = "Ok"
		return *response
	}

	if url == "/create/user" && method == "POST" {
		response.url = url
		response.httpCode = 201
		response.body = "User Created"
		return *response
	}

	return *response
}
