package main

import (
	"NoteAssistant/common"
	"NoteAssistant/router"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDataBase()
	r := gin.Default()
	r = router.Collector(r)
	r.Run(":8080")
}
