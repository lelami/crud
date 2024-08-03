package authclient

import (
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

var c *fasthttp.HostClient

func Init(host string) {
	c = &fasthttp.HostClient{
		Addr: host,
	}
}

func ValidateToken(token string) bool {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI("http://" + c.Addr + "/get_user_info")
	req.Header.Set(fasthttp.HeaderAuthorization, token)
	req.Header.SetHost(c.Addr)
	req.Header.SetMethod(fasthttp.MethodGet)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := c.Do(req, resp)
	if err != nil {
		return false
	}
	log.Println(resp)
	if resp.StatusCode() != http.StatusOK {
		return false
	}

	return true
}
