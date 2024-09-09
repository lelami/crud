package authclient

import (
	"encoding/json"
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

type Response struct {
	Success bool     `json:"success"`
	Data    UserData `json:"data"`
}

type UserData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

func ValidateToken(token string) (bool, UserData) {
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
		return false, UserData{}
	}
	log.Println(resp)
	if resp.StatusCode() != http.StatusOK {
		return false, UserData{}
	}

	var userInfo Response
	err = json.Unmarshal(resp.Body(), &userInfo)
	if err != nil {
		log.Println("Error unmarshalling response body:", err)
		return false, UserData{}
	}

	return true, userInfo.Data
}
