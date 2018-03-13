package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"github.com/buaazp/fasthttprouter"
)

var (
	logTopic = "log"
	metricTopic = "metric"
)

func consumeLogRequest(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(201)
}

func consumeMetricRequest(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(201)
}

func main() {
	r := fasthttprouter.New()
	r.PUT("/log", consumeLogRequest)
	r.PUT("/metric", consumeMetricRequest)

	fasthttp.ListenAndServe(":80", r.Handler)

	log.Println("Launched server")
}
