package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/petronetto/echo-sample/route"
)

func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	router := route.Init()
	router.Run(fasthttp.New(":8888"))
}
