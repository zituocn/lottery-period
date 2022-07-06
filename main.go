package main

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/lottery-period/router"
)

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	router.APIRouter(r)
	r.Run()
}
