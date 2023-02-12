package main

import (
	"fmt"
	"os"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var downloadDir = "/tmp/ptt-downloads"
var uploadDir = "/tmp/ptt-uploads"

func main() {
	e := echo.New()

	listen := "8080"
	// downloadDir := "/tmp/ptt-downloads"
	zapLogger, _ := zap.NewProduction()
	log := zap.SugaredLogger(*zapLogger.Sugar())

	e.Use(echozap.ZapLogger(zapLogger))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := os.MkdirAll(downloadDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	e.Static("/downloads", downloadDir)
	e.Static("/", "public")
	e.POST("/upload", upload)

	e.HideBanner = true

	log.Infow("listening",
		"address", fmt.Sprintf("http://localhost:%s", listen),
	)
	e.Logger.Fatal(e.Start(":8080"))

	// data, err := parseFile("sample.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = generateXls(data, "book1.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
