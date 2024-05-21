package stock

const (
	CREATE_PURCHASE_ASN_URL = "/WebApi/Api/ASN/CreatePurchaseAsn"
	CANCEL_PURCHASE_ASN_URL = "/WebApi/Api/ASN/CancelPurchaseAsn"
	TERMINATION_PURCHASE_ASN_URL = "/WebApi/Api/ASN/TerminationPurchaseAsn"
	CREATE_SALE_RETURN_ASN_URL = "/WebApi/Api/ASN/CreateSaleReturnAsn"
	CANCEL_SALES_RETURN_ASN_URL  = "/WebApi/Api/ASN/CancelSalesReturnAsn"
	CREATE_QUALITY_TEST_ORDER_URL = "/WebApi/Api/Quality/CreateQualityTestOrder"
	SEARCH_RECEIPTS_URL = "/WebApi/Api/Receipt/SearchReceipts"
	CREATE_RECEIPT_URL = "/WebApi/Api/Receipt/CreateReceipt"
)

type ReqCreatePurchaseAsn struct {
	Data []*CreatePurchaseAsn `json:"data"`
}

type CreatePurchaseAsn struct {
	OrderNo string `json:"orderNo"`				 //采购单号
	CarrierReferNo string `json:"CarrierReferNo"`
	Warehouse string `json:"warehouse"`			 //仓库编号
	SupplierId string `json:"supplierId"`		 //供应商编号
	ClientNo string `json:"clientNo"`
	Apuser string `json:"apuser"`				//采办人
	Aptime string `json:"aptime"`
	FormNo string `json:"formNo"`
	CreateUserName string `json:"createUserName"`
	CreateDate string `json:"createDate"`
	AuditUserName string `json:"auditUserName"`
	AuditDate string `json:"auditDate"`
	Tocode string `json:"tocode"`				//采购入库单标识
	Reason string `json:"reason"`				//入库原因(Inbound=上新,Replenish=补货,Other=其他，非必填)
	Memo string `json:"memo"`					//备注
	Remark string `json:"remark"`
	PurchaseTypeCode string `json:"purchaseTypeCode"`
	DepartmentName string `json:"DepartmentName"`
	PredictArrivalDate string `json:"PredictArrivalDate"`
	Detail []*PurchaseAsnDetail `json:"detail"`
}

type PurchaseAsnDetail struct {
	LineNo string `json:"LineNo"`
	FromLineNo string `json:"FromLineNo"`
	Lotattribute01 string `json:"lotattribute01"`	//批次属性一
	Lotattribute02 string `json:"lotattribute02"`	//批次属性二
	Lotattribute03 string `json:"lotattribute03"`
	Lotattribute04 string `json:"lotattribute04"`
	Lotattribute05 string `json:"lotattribute05"`
	Lotattribute06 string `json:"lotattribute06"`
	Lotattribute07 string `json:"lotattribute07"`
	Lotattribute08 string `json:"lotattribute08"`
	Lotattribute09 string `json:"lotattribute09"`
	Lotattribute010 string `json:"lotattribute010"`
	Lotkey string `json:"lotkey"`
	Isgift bool `json:"isgift"`
	Sku string `json:"sku"`							//skucode
	Qty int `json:"qty"`
	Storage string `json:"storage"`
	Price int `json:"price"`
	Userdef10 string `json:"userdef10"`
	Expectedweight string `json:"expectedweight"`
	Memo string `json:"memo"`
	InvertoryType string `json:"InvertoryType"`
	Zonecode string `json:"zonecode"`
	Locationcode string `json:"locationcode"`
}

type ReqCancelAsnList struct {
	Data []*CancelAsnList `json:"data"`
}

type CancelAsnList struct{
	Ordeno string `json:"ordeno"`
	Reason string `json:"reason"`
	Orderno string `json:"orderno"`
}

type ReqTerminationAsnList struct {
	Data []*TerminationAsnList `json:"data"`
}

type TerminationAsnList struct {
	Ordeno string `json:"ordeno"`
	Reason string `json:"reason"`
	Orderno string `json:"orderno"`
}

type ReqCreateSaleReturnAsn struct {
	Data []*CreateSaleReturnAsn `json:"data"`
}

type CreateSaleReturnAsn struct {
	Tradeid string `json:"tradeid"`				//来源单号
	Retono string `json:"retono"`				//退货单号
	Orderno string `json:"orderno"`				//原订单号
	Type string `json:"type"`					//入库单类型
	Whcode string `json:"whcode"`				//仓库编码
	Expresscode string `json:"expresscode"`		//快递公司编号
	LogisticsCode string `json:"logisticsCode"`
	Expressnumber string `json:"expressnumber"`	//快递单号
	ReturnReason string `json:"returnReason"`
	Shopcode string `json:"shopcode"`			//店铺编码
	Shopname string `json:"shopname"`			//店铺名称
	Tocode string `json:"tocode"`				//第三方业务系统标识
	Crtime string `json:"crtime"`
	Cruser string `json:"cruser"`				//创建人编码
	Crusno string `json:"crusno"`				//创建人名称
	Aptime string `json:"aptime"`
	Apuser string `json:"apuser"`				//审核人
	Consigneename string `json:"consigneename"`	//发货人姓名
	Mobile string `json:"mobile"`				//发货人手机
	Consigneeaddress string `json:"consigneeaddress"` //发货人地址
	Note string `json:"note"`					//备注
	Clientno string `json:"clientno"`			//货主编号
	Vipcode string `json:"vipcode"`
	Customercode string `json:"customercode"`
	Buyernick string `json:"buyernick"`			//会员昵称
	Detail []*SaleReturnAsnDetail `json:"detail"`
}

type SaleReturnAsnDetail struct {
	Sno string `json:"sno"`
	Sku string `json:"sku"`				//sku
	Qty int `json:"qty"`
	Price int `json:"price"`
	Amount int `json:"amount"`
	Lineno string `json:"lineno"`
	Userdef10 string `json:"userdef10"`
	Lotkey string `json:"lotkey"`
	Lotattribute01 string `json:"lotattribute01"`	//批次属性一
	Lotattribute02 string `json:"lotattribute02"`	//批次属性一
	Lotattribute07 string `json:"lotattribute07"`
	Lotattribute08 string `json:"lotattribute08"`
	Isgift bool `json:"isgift"`
	Clientno string `json:"clientno"`
	FromLineNo string `json:"FromLineNo"`
	Shijituikuanje int `json:"shijituikuanje"`
	OffsetAmount int `json:"OffsetAmount"`
	ActualAmount int `json:"ActualAmount"`
	Remark string `json:"Remark"`
	ZoneNo string `json:"ZoneNo"`
	Tradeid string `json:"tradeid"`
	IsCredited bool `json:"isCredited"`
	InventoryStatusCode string `json:"InventoryStatusCode"`
}

type ReqCancelSalesReturnAsn struct {
	Data []CancelSalesReturnAsn `json:"data"`
}

type CancelSalesReturnAsn struct {
	Ordeno string `json:"ordeno"`
	Reason string `json:"reason"`
	Orderno string `json:"orderno"`
}

type ReqCreateQualityTestOrder struct {
	Data []*CreateQualityTestOrder `json:"data"`
}

type CreateQualityTestOrder struct {
	Aptime string `json:"aptime"`
	Apuser string `json:"apuser"`
	Descri string `json:"descri"`			//质检描述
	Fromno string `json:"fromno"`
	FWmsOrderNo string `json:"fWmsOrderNo"`
	Qty int `json:"qty"`
	Sku string `json:"sku"`
	Sn string `json:"sn"`
	IsAccept bool `json:"isAccept"`
}

type ReqSearchReceipts struct {
	Data []string `json:"data"`
}

type ReqCreateReceipt struct {
	Data []*CreateReceipt `json:"data"`
}

type CreateReceipt struct {
	OrderNo string `json:"orderNo"`					//采购单号
	CarrierReferNo string `json:"CarrierReferNo"`
	Warehouse string `json:"warehouse"`				//仓库编号
	SupplierId string `json:"supplierId"`			//供应商编号
	ClientNo string `json:"clientNo"`
	Apuser string `json:"apuser"`					//采办人
	Aptime string `json:"aptime"`
	FormNo string `json:"formNo"`
	CreateUserName string `json:"createUserName"`
	CreateDate string `json:"createDate"`
	AuditUserName string `json:"auditUserName"`
	AuditDate string `json:"auditDate"`
	Tocode string `json:"tocode"`					//采购入库单标识
	Reason string `json:"reason"`					//入库原因(Inbound=上新,Replenish=补货,Other=其他，非必填)
	Memo string `json:"memo"`						//备注
	Remark string `json:"remark"`
	PurchaseTypeCode string `json:"purchaseTypeCode"`
	DepartmentName string `json:"DepartmentName"`
	PredictArrivalDate string `json:"PredictArrivalDate"`
	Detail []*CreateDeceiptDetail `json:"detail"`
}

type CreateDeceiptDetail struct {
	LineNo string `json:"LineNo"`
	FromLineNo string `json:"FromLineNo"`
	Lotattribute01 string `json:"lotattribute01"`	//批次属性一
	Lotattribute02 string `json:"lotattribute02"`	//批次属性二
	Lotattribute03 string `json:"lotattribute03"`
	Lotattribute04 string `json:"lotattribute04"`
	Lotattribute05 string `json:"lotattribute05"`
	Lotattribute06 string `json:"lotattribute06"`
	Lotattribute07 string `json:"lotattribute07"`
	Lotattribute08 string `json:"lotattribute08"`
	Lotattribute09 string `json:"lotattribute09"`
	Lotattribute10 string `json:"lotattribute10"`
	Lotkey string `json:"lotkey"`
	Isgift bool `json:"isgift"`
	Sku string `json:"sku"`
	Qty int `json:"qty"`
	Storage string `json:"storage"`
	Price int `json:"price"`
	Userdef10 string `json:"userdef10"`
	Expectedweight string `json:"expectedweight"`
	Memo string `json:"memo"`
	InvertoryType string `json:"InvertoryType"`
	Zonecode string `json:"zonecode"`
	Locationcode string `json:"locationcode"`
}

//createPurchaseAsn

//GuanyiSaleReturnAsn