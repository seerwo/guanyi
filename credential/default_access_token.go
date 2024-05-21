package credential

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/guanyi/cache"
	"github.com/seerwo/guanyi/util"
	"strconv"
	"sync"
	"time"
)

const(
	//AccessTokenURL get access_token interface
	accessTokenURL = ""
	//CacheKeyWmsPrefix set cache key prefix
	CacheKeyWmsPrefix = "goguanyi_wms_"
)

//DefaultAccessToken default AccessToken to get
type DefaultAccessToken struct {
	appID string
	appSecret string
	cacheKeyPrefix string
	cache cache.Cache
	accessTokenLock *sync.Mutex
}

//NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is inneed")
	}
	return &DefaultAccessToken{
		appID: appID,
		appSecret: appSecret,
		cache:cache,
		cacheKeyPrefix: cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

//ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}

//GetAccessToken get access_token, get cache
func (ak *DefaultAccessToken) GetAccessToken()(accessToken string, err error){
	//lock
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	//cache invalid
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(ak.appID, ak.appSecret)
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	if err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second); err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

func (ak *DefaultAccessToken) GetAccessParam()(accessParam string, err error){

	m := make(map[string]string)
	m["noncestr"] = util.RandomStr(16)
	m["timestamp"] = strconv.FormatInt(util.GetCurrTS(), 10)
	m["app-key"] = ak.appID

	m["signed"], err = util.ParamSign(m, ak.appSecret)

	accessParam = util.UrlParam(m, "")
	return
}

//GetTokenFromServer reforce from server
func GetTokenFromServer(appID,appSecret string) (resAccessToken ResAccessToken, err error){
	url := fmt.Sprintf("%s?grant_type=client_redential&appid-%s&secret=%s", accessTokenURL, appID, appSecret)
	var body []byte
	if body, err = util.HTTPGet(url); err != nil {
		return
	}
	if err = json.Unmarshal(body, &resAccessToken); err != nil {
		return
	}
	if resAccessToken.ErrMsg != ""{
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
	}
	return
}