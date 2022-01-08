package handler

import (
	"mockgoproxy/constants"
	"net/http"
)

func CopyHeaders(dst,src http.Header)  {
	for key,values:=range src{
		for _,value :=range values{
			dst.Add(key,value)
		}
	}
}

func ClearHeaders(headers http.Header)  {
	for key, _ := range headers {
		headers.Del(key)
	}
}

func RemoveProxyHeaders(req *http.Request)  {
	req.RequestURI=""
	req.Header.Del(constants.PROXYCONNECTION)
	req.Header.Del(constants.CONNECTION)
	req.Header.Del(constants.KEEPALIVE)
	req.Header.Del(constants.PROXYAUTHENTICATE)
	req.Header.Del(constants.PROXYAUTHORIZATION)
	req.Header.Del(constants.TE)
	req.Header.Del(constants.TRAILERS)
	req.Header.Del(constants.TRANSFERENCODING)
	req.Header.Del(constants.UPGRADE)
}