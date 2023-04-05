package proxy

type IServer interface {
	handleRequest(string, string) Response
}

type Response struct {
	url      string
	httpCode int
	body     string
}
