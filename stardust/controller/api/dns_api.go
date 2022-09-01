package api

import (
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
	"github.com/sipt/panda141035/dns"
=======
	"github.com/sipt/shuttle/dns"
>>>>>>> ba7fba0 (built with no error)
)

func DNSCacheList(ctx *gin.Context) {
	ctx.JSON(200, &Response{
		Data: dns.DNSCacheList(),
	})
}
func ClearDNSCache(ctx *gin.Context) {
	dns.ClearDNSCache()
	ctx.JSON(200, &Response{})
}

