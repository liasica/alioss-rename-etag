// Copyright (C) alioss-rename-etag. 2025-present.
//
// Created at 2025-08-27, by liasica

package controller

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Invoke(c echo.Context) error {
	for s, ss := range c.Request().Header {
		log.Printf("%s: %s\n", s, strings.Join(ss, " | "))
	}

	rc := c.Request().Body
	if rc != nil {
		defer rc.Close()
		b, _ := io.ReadAll(rc)
		log.Println(string(b))
	}

	return c.String(http.StatusOK, "Hello, World!")
}
