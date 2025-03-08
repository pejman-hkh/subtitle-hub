package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func proxyToApi(c *gin.Context) {
	remote, err := url.Parse("http://api:8083")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("site")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func proxyToSite(c *gin.Context) {
	remote, err := url.Parse("http://site:3000")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("site")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func main() {

	static := gin.New()
	static.Static("/.well-known/", "/app/.well-known/")

	r := gin.Default()
	r.Use(Cors())
	r.Any("/*site", func(ctx *gin.Context) {
		param := ctx.Param("site")
		paramSplit := strings.Split(param, "/")

		if paramSplit[1] == ".well-known" {
			static.HandleContext(ctx)
		} else if paramSplit[1] == "api" || paramSplit[1] == "docs" {
			proxyToApi(ctx)
		} else {
			proxyToSite(ctx)
		}
	})

	r1 := gin.Default()
	r1.GET("/*path", func(c *gin.Context) {
		c.Redirect(302, "https://"+c.Request.Host+c.Request.URL.Path)
	})

	go r.RunTLS(":443", "/app/tls/localhost.crt", "/app/tls/localhost.key")
	r1.Run(":80")
}
