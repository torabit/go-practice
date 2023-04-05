package proxy

type Nginx struct {
	appliation        *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		appliation:        &Application{},
		maxAllowedRequest: 1,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) Response {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return Response{
			url:      url,
			httpCode: 403,
			body:     "Not Allowed",
		}
	}
	return n.appliation.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
