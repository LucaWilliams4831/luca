package conf

import (
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/sipt/panda141035"
	"github.com/sipt/panda141035/shuttle_config"
=======
	"github.com/sipt/shuttle"
	"github.com/sipt/shuttle/config"
>>>>>>> ba7fba0 (built with no error)
)

func GetHttpMap(ctx *gin.Context) {
	conf := config.CurrentConfig()
	ctx.JSON(200, &Response{
		Data: conf.GetHTTPMap(),
	})
}

func SetHttpMap(ctx *gin.Context) {
	conf := config.CurrentConfig()
	newConf := &config.Config{}
	*newConf = *conf
	var data = &config.HttpMap{}
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(500, &Response{Code: 1, Message: err.Error()})
		return
	}
	newConf.SetHTTPMap(data)
	err = shuttle.ApplyHTTPModifyConfig(newConf)
	if err != nil {
		ctx.JSON(500, &Response{Code: 1, Message: err.Error()})
		return
	}
	conf.SetHTTPMap(data)
	config.SaveConfig(config.CurrentConfigFile(), conf)
	if err != nil {
		ctx.JSON(500, &Response{
			Code:    1,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, &Response{})
	return
}
