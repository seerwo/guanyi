package wms

import (
	"github.com/seerwo/guanyi/credential"
	"github.com/seerwo/guanyi/wms/config"
	"github.com/seerwo/guanyi/wms/context"
	"github.com/seerwo/guanyi/wms/oauth"
	"github.com/seerwo/guanyi/wms/stock"
)

//Wms wms relate api
type Wms struct {
	ctx *context.Context
}

//NewWms instance api
func NewWms(cfg *config.Config) *Wms {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyWmsPrefix, cfg.Cache)
	ctx := &context.Context{
		Config: cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Wms{ctx:ctx}
}

//SetAccessTokenHandle custom access_token get method
func(wms *Wms) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle){
	wms.ctx.AccessTokenHandle = accessTokenHandle
}

//GetContext get Context
func (wms *Wms) GetContext() *context.Context {
	return wms.ctx
}

//GetAcessToken get access_token
func (wms *Wms) GetAccessToken()(string, error){
	return wms.ctx.GetAccessToken()
}

//GetOauth oauth2
func (wms *Wms) GetOauth() *oauth.Oauth{
	return oauth.NewOauth(wms.ctx)
}

//GetWms
func (wms *Wms) GetOut() *stock.Out {
	return stock.NewOut(wms.ctx)
}

func (wms *Wms) GetStock() *stock.Stock {
	return stock.NewStock(wms.ctx)
}