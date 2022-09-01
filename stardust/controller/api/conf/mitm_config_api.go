package conf

import (
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/sipt/panda141035"
=======
	"github.com/sipt/shuttle"
>>>>>>> ba7fba0 (built with no error)
)

func GetMitMRules(ctx *gin.Context) {
	var response Response
	response.Data = shuttle.GetMitMRules()
	ctx.JSON(200, response)
}

func AppendMitMRules(ctx *gin.Context) {
	d := ctx.Query("domain")
	if len(d) > 0 {
		shuttle.AppendMitMRules(d)
	}
	var response Response
	response.Data = shuttle.GetMitMRules()
	ctx.JSON(200, response)
}

func DelMitMRules(ctx *gin.Context) {
	d := ctx.Query("domain")
	if len(d) > 0 {
		shuttle.RemoveMitMRules(d)
	}
	var response Response
	response.Data = shuttle.GetMitMRules()
	ctx.JSON(200, response)
}
