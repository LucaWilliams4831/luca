<<<<<<< HEAD
package shuttle
=======
package stardust
>>>>>>> ba7fba0 (built with no error)

import "errors"

var (
	ErrorReadTimeOut  = errors.New("read time out")
	ErrorWriteTimeOut = errors.New("write time out")
	ErrorReject       = errors.New("connection reject")

	ErrorServerNotFound = errors.New("server or server group not found")

	ErrorUnknowType = errors.New("unknow type")
)
