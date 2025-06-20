package routes

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStaticRoutes(embedded fs.FS, engine *gin.Engine) {
	sub, err := fs.Sub(embedded, "web/dist")
	if err != nil {
		log.Fatal("init static routes failed!")
	}

	hfs := http.FS(sub)

	engine.StaticFS("/", hfs)
}
