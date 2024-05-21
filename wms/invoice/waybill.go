package invoice

const(
	GET_EWAYBILL_ORDER_ZD_URL = "/WebApi/Api/EWaybill/GetEWaybillOrderZD"
	GET_EWAYBILL_ORDER_KYSY_URL = "/WebApi/Api/EWaybill/GetEWaybillOrderKYSY"
	SEARCH_EWAYBILL_NO = "/WebApi/Api/EWaybill/SearchEWaybillNo"
	UPDATE_EWAYBILL_URL = "/WebApi/Api/EWaybill/UpdateEWaybill"
	UPDATE_EWAYBILLS_URL = "/WebApi/Api/EWaybill/UpdateEWaybill"
	CANCEL_EWAYBILL_URL = "/WebApi/Api/EWaybill/CancelEWaybill"
	SEARCH_EWAYBILL_LOGISTIC_AREA_URL = "/WebApi/Api/EWaybill/SearchEWaybillLogisticArea"
)

type ReqEWaybillOrderZD struct {
	Orderid string `json:"orderid"`				//客户订单号
	ParcelQuantity int `json:"parcel_quantity"`
	Head string `json:"Head"`
}

type ReqEWaybillOrderKYSY struct {
	ReqEWaybillOrderZD
}

type ReqSearchEWaybillNo struct {
	OrderInfos []OrderInfos `json:"OrderInfos"`
	RequestCount int `json:"RequestCount"`
	LogisticsCompanyID string `json:"LogisticsCompanyID"`
	CarrierCode string `json:"CarrierCode"`
	BusinessType string `json:"BusinessType"`
	ShopNo string `json:"ShopNo"`				//店铺名称
	WarehouseNo string `json:"WarehouseNo"`
	StorerID string `json:"StorerID"`
}

type OrderInfos struct {
	OrderID string `json:"OrderID"`
	ParterID string `json:"ParterID"`
	Payer int `json:"Payer"`
	PayMethod int `json:"PayMethod"`
	Sender struct {
		Name string `json:"Name"`
		Company string `json:"Company"`
		Province string `json:"Province"`
		City string `json:"City"`
		District string `json:"District"`
		Town string `json:"Town"`
		Address string `json:"Address"`
		Zip string `json:"Zip"`
		TelPhone string `json:"TelPhone"`
		MobilePhone string `json:"MobilePhone"`
	} `json:"Sender"`
	Receiver struct {
		Name string `json:"Name"`
		Company string `json:"Company"`
		Province string `json:"Province"`
		City string `json:"City"`
		District string `json:"District"`
		Town string `json:"Town"`
		Address string `json:"Address"`
		Zip string `json:"Zip"`
		TelPhone string `json:"TelPhone"`
		MobilePhone string `json:"MobilePhone"`
	} `json:"Receiver"`
	WaveKey string `json:"WaveKey"`
	ProductList []*ProductList `json:"ProductList"`
	CollectingAmount int `json:"CollectingAmount"`
	TotalWeight int `json:"TotalWeight"`
	NetWeight int `json:"NetWeight"`
	SendStartTime string `json:"SendStartTime"`
	SendEndTime string `json:"SendEndTime"`
	Remark string `json:"Remark"`
	ReferOrderKey string `json:"ReferOrderKey"`
	IsCOD string `json:"IsCOD"`
	DeclareValue string `json:"DeclareValue"`
	CarrierReferNo string `json:"CarrierReferNo"`
	VendorId int `json:"vendor_id"`
	Is3PL bool `json:"Is3PL"`
	PackageQuantity int `json:"PackageQuantity"`
	TicketType string `json:"TicketType"`
}

type ProductList struct {
	Name string `json:"Name"`				//商品名称
	Description string `json:"Description"` //详细信息
	Price int `json:"Price"`
	Qty int `json:"Qty"`
	Remark string `json:"Remark"`			//备注
}

type ReqUpdateEWaybill struct {
	Payer int `json:"Payer"`
	PayMethod int `json:"PayMethod"`
	Sender struct {
		Name string `json:"Name"`
		Company string `json:"Company"`
		Province string `json:"Province"`
		City string `json:"City"`
		District string `json:"District"`
		Town string `json:"Town"`
		Address string `json:"Address"`
		Zip string `json:"Zip"`
		TelPhone string `json:"TelPhone"`
		MobilePhone string `json:"MobilePhone"`
	} `json:"Sender"`
	Receiver struct {
		Name string `json:"Name"`
		Company string `json:"Company"`
		Province string `json:"Province"`
		City string `json:"City"`
		District string `json:"District"`
		Town string `json:"Town"`
		Address string `json:"Address"`
		Zip string `json:"Zip"`
		TelPhone string `json:"TelPhone"`
		MobilePhone string `json:"MobilePhone"`
	} `json:"Receiver"`
	WaveKey string `json:"WaveKey"`
	ProductList []*ProductList `json:"ProductList"`
	CollectingAmount int `json:"CollectingAmount"`
	TotalWeight int `json:"TotalWeight"`
	NetWeight int `json:"NetWeight"`
	SendStartTime string `json:"SendStartTime"`
	SendEndTime string `json:"SendEndTime"`
	Remark string `json:"Remark"`
	ReferOrderKey string `json:"ReferOrderKey"`
	IsCOD bool `json:"IsCOD"`
	EWaybillID string `json:"EWaybillID"`
	OrderID string `json:"OrderID"`
	ParterID string `json:"ParterID"`
	LogisticArea string `json:"LogisticArea"`
	LogisticsCompanyID string `json:"LogisticsCompanyID"`
	CarrierCode string `json:"CarrierCode"`
	BusinessType string `json:"BusinessType"`
	ShopNo string `json:"ShopNo"`
	WarehouseNo string `json:"WarehouseNo"`
	StorerID string `json:"StorerID"`
}

type ReqUpdateEWaybills struct {
	ReqUpdateEWaybill
}

type ReqCancelEWaybill struct {
	EWaybillID string `json:"EWaybillID"`				//物流单号
	CancelReason string `json:"CancelReason"`			//取消原因
	OrderID string `json:"OrderID"`						//出库单号
	ParterID string `json:"ParterID"`
	LogisticsCompanyID string `json:"LogisticsCompanyID"`
	CarrierCode string `json:"CarrierCode"`
	BusinessType string `json:"BusinessType"`
	ShopNo string `json:"ShopNo"`						//店铺编号
	WarehouseNo string `json:"WarehouseNo"`
	StorerID string `json:"StorerID"`
}

type ReqSearchEWaybillLogisticArea struct {
	ReqSearchEWaybillNo
}