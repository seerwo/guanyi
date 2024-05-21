package credential

type AccessParam struct {
	AppId string `json:"app_id"`
	Timestamp int64 `json:"timestamp"`
	NonceStr string `json:"nonce_str"`
	Url string `json:"url"`
	Signature string `json:"signature"`
}

func NewAccessParam() AccessParam {
	accessParam := *new(AccessParam)
	return accessParam
}