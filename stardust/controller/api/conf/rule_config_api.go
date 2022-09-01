package conf

import (
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/sipt/panda141035/shuttle_config"
	"github.com/sipt/panda141035/rule"
=======
	"github.com/sipt/shuttle/config"
	"github.com/sipt/shuttle/rule"
>>>>>>> ba7fba0 (built with no error)
)

const (
	RuleOpAdd    = "add"
	RuleOpRemove = "remove"
	RuleOpUpdate = "update"
)

func GetRule(ctx *gin.Context) {
	conf := config.CurrentConfig()
	ctx.JSON(200, &Response{
		Data: conf.GetRule(),
	})
}

type RuleOp struct {
	Op    string   `json:"op"`
	Index int      `json:"index"`
	Rule  []string `json:"rule"`
}

func SetRule(ctx *gin.Context) {
	conf := config.CurrentConfig()
	newConf := &config.Config{}
	*newConf = *conf
	data := make([][]string, 0, 64)
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(500, &Response{Code: 1, Message: err.Error()})
		return
	}
	newConf.SetRule(data)
	err = rule.ApplyConfig(newConf)
	if err != nil {
		ctx.JSON(500, &Response{Code: 1, Message: err.Error()})
		return
	}
	conf.SetRule(data)
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
