package api

import (
<<<<<<< HEAD
	"github.com/sipt/panda141035"
=======
	"github.com/sipt/shuttle"
>>>>>>> ba7fba0 (built with no error)
	"github.com/gin-gonic/gin"
)

func GetRecords(ctx *gin.Context) {
	ctx.JSON(200, &Response{
		Data: shuttle.GetRecords(),
	})
}
func ClearRecords(ctx *gin.Context) {
	shuttle.ClearRecords()
	dump := shuttle.GetDump()
	if dump != nil {
		dump.Clear()
	}
	ctx.JSON(200, &Response{})
}
