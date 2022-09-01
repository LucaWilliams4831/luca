<<<<<<< HEAD
package shuttle
=======
package stardust
>>>>>>> ba7fba0 (built with no error)

var (
	ControllerDomain string
	ControllerPort   string
	HTTPProxyPort    string
)

type IConfigValue interface {
	GetControllerDomain() string
	GetControllerPort() string
	GetHTTPPort() string
}

func InitConfigValue(conf IConfigValue) {
	ControllerDomain = conf.GetControllerDomain()
	ControllerPort = conf.GetControllerPort()
	HTTPProxyPort = conf.GetHTTPPort()
}
