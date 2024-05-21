package invoice

const(
	UPDATE_INTERCEPT_ORDER_URL = "/WebApi/Api/InterceptOrder/UpdateInterceptOrder"
	RECEIPT_STATUS_CODE_URL = "/WebApi/Api/ChangeOrdStat/ReceiptStatusCode"
	SHIPMENT_CHECK_URL = "/WebApi/Api/ChangeOrdStat/ShipmentCheck"
)

type ReqUpdateInterceptOrder struct {
	Data []*UpdateInterceptOrder `json:"data"`
}

type UpdateInterceptOrder struct {
	WarehouseNo string `json:"WarehouseNo"`			//仓库代码
	StorerNo string `json:"StorerNo"`				//货主代码
	OrderType string `json:"OrderType"`				//DN
	ReferNo string `json:"ReferNo"`					//参考单号
}

type ReqReceiptStatusCode struct {
	Data []string `json:"data"`
}

type ReqShipmentCheck struct {
	Data []string `json:"data"`
}
