package handler

import "github.com/zituocn/gow"

func ErrorHandler(c *gow.Context) {
	ServerJSON(c, 404, 404, "page not found")
}
