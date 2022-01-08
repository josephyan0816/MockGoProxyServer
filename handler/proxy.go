package handler

import "net/http"

type ProxyServer struct {
	Travel  *http.Transport
	Browser string
}

func NewProxyServer() *http.Server {

	return &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &ProxyServer{
			Travel: &http.Transport{
				Proxy:http.ProxyFromEnvironment,
				DisableKeepAlives: false,
			},
		},
	}

}

func (ps *ProxyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

}