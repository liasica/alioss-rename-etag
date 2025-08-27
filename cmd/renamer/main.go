// Copyright (C) alioss-rename-etag. 2025-present.
//
// Created at 2025-08-27, by liasica

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// 上下文信息
// https://help.aliyun.com/zh/functioncompute/fc-3-0/user-guide/context-and-log-format-1?spm=5176.fcnext.console-base_help.dexternal.1aa32f03mxVDCk
func main() {
	e := echo.New()
	e.POST("/invoke", func(c echo.Context) error {
		for s, ss := range c.Request().Header {
			fmt.Printf("%s: %s\n", s, strings.Join(ss, " | "))
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	log.Fatal(e.Start(":8080"))
}
