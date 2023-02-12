package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
)

func newHash(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func upload(c echo.Context) error {
	maxSizeMb := 5
	maxSizeBytes := int64(1024 * 1024 * maxSizeMb)
	maxSizeHuman := fmt.Sprintf("%dMb", maxSizeMb)
	htmlFooter := "<p><a href=\"/\">Return to the main page</a></p>"
	nowNano := time.Now().UnixNano()
	uploadDir := fmt.Sprintf("%s/%d", uploadDir, nowNano)
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return c.HTML(http.StatusBadRequest, fmt.Sprintf("<p>Error: %s</p>%s", err, htmlFooter))
	}
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// check size
	if c.Request().ContentLength > maxSizeBytes {
		return c.HTML(http.StatusBadRequest,
			fmt.Sprintf("<p>Sorry, the biggest file you can upload is %s, ask david to make it bigger</p>%s", maxSizeHuman, htmlFooter))
	}

	// check extension
	fileExt := filepath.Ext(file.Filename)
	if fileExt != ".txt" {
		return c.HTML(http.StatusBadRequest,
			fmt.Sprintf(
				"<p>Sorry, you uploaded a file called: %s (<b>%s</b>)<br> <br> only .txt is supported</p>%s", file.Filename, fileExt, htmlFooter))
	}
	uploadedFileName := fmt.Sprintf("%s/%d%s", uploadDir, nowNano, fileExt)
	dst, err := os.Create(uploadedFileName)
	if err != nil {
		return c.HTML(http.StatusBadRequest, fmt.Sprintf("<p>Error: %s</p>%s", err, htmlFooter))
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// see if I can parse the file OK
	data, err := parseFile(uploadedFileName)
	if err != nil {
		msg := fmt.Sprintf("you uploaded the file OK but I didn't understand it<br>Error: %v", err)
		return c.HTML(http.StatusBadRequest, fmt.Sprintf("%s%s", msg, htmlFooter))
	}

	hash := newHash(32)
	newXlsFile := fmt.Sprintf("%s/%s.xlsx", downloadDir, hash)

	err = renderXls(data, newXlsFile)

	downloadLink := fmt.Sprintf("<a href=\"/downloads/%s.xlsx\">download the results here</a>", hash)

	msg := fmt.Sprintf("<p>File %s uploaded and parsed OK successfully (%s)</p><br>%s<br>",
		file.Filename,
		uploadedFileName,
		downloadLink,
	)
	return c.HTML(http.StatusOK, fmt.Sprintf("%s%s", msg, htmlFooter))
}
