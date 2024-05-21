package context

import (
	"github.com/seerwo/guanyi/credential"
	"github.com/seerwo/guanyi/wms/config"
)

//Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}