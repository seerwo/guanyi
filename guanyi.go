package guanyi

import (
	"github.com/seerwo/guanyi/cache"
	"github.com/seerwo/guanyi/wms"
	wmsconfig "github.com/seerwo/guanyi/wms/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	//Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	//Output to stdout instead of the default stderr
	//Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	//Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

//Guanyi struct
type Guanyi struct {
	cache cache.Cache
}

//NewGuanyi init
func NewGuanyi() *Guanyi {
	return &Guanyi{}
}

//SetCache set cache
func (gn *Guanyi) SetCache(cache cache.Cache){
	gn.cache = cache
}

//GetWms get Wms instance
func (gn *Guanyi) GetWms(cfg *wmsconfig.Config) *wms.Wms {
	if cfg.Cache == nil {
		cfg.Cache = gn.cache
	}
	return wms.NewWms(cfg)
}