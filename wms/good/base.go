package good

const(
	IMPORT_CODES_URL = "/WebApi/Api/Code/ImportCodes"
	CREATE_LOT_URL = "/WebApi/Api/Lot/CreateLot"
	CREATE_SHOP_URL = "/WebApi/Api/Shop/CreateShop"
	LOAD_SHOP_URL = "/WebApi/Api/Shop/LoadShop"
	CREATE_SUPPLIER = "/WebApi/Api/Supplier/CreateSupplier"
	SELECT_OPERATION_LOG_URL = "/WebApi/Api/OperationLog/SelectOperationLog"
	PUSH_OPERATION_LOG_URL = "/WebApi/Api/OperationLog/PushOperationLog"
	PURCHASE_ENTRY_ORDER_URL = "/WebApi/Api/Base/PurchaseEntryOrder"
	RETRUN_ENTRY_ORDER_URL = "/WebApi/Api/Base/ReturnEntryOrder"
	TRANSFER_ENTRY_ORDER_URL = "/WebApi/Api/Base/TransferEntryOrder"
	SALE_DELIVERY_ORDER_URL = "/WebApi/Api/Base/SaleDeliveryOrder"
	RETRUN_DELIVERY_ORDER = "/WebApi/Api/Base/ReturnDeliveryOrder"
	TRANSFER_DELIVERY_ORDER = "/WebApi/Api/Base/TransferDeliveryOrder"
	ADJUST_ORDER_URL = "/WebApi/Api/Base/AdjustOrder"
	INVENTORY_ORDER_URL = "/WebApi/Api/Base/InventoryOrder"
)

type ReqImportCodes struct {
	Data []*ImportCode `json:"data"`
}

type ImportCode struct {
	FListName string `json:"fListName"`				//Country
	FDescription string `json:"fDescription"`		//国家
	FIsSystem bool `json:"fIsSystem"`
	FIsReadOnly bool `json:"fIsReadOnly"`
	Details []*Details `json:"Details"`
}

type Details struct {
	FCodeValue string `json:"fCodeValue"`
	FCodeName string `json:"fCodeName"`				//美国
	FIsDefault bool `json:"fIsDefault"`
	FIsActive bool `json:"fIsActive"`
	FIsSystem bool `json:"fIsSystem"`
	FDescription string `json:"fDescription"`		//美国
}

type ReqCreateLot struct {
	Data []*CreateLot `json:"data"`
}

type CreateLot struct {
	FLotID string `json:"fLotID"`
	FSKUID string `json:"fSKUID"`
	FLotKey string `json:"fLotKey"`
	FLotAttribute01 string `json:"fLotAttribute01"`
	FLotAttribute02 string `json:"fLotAttribute02"`
	FLotAttribute03 string `json:"fLotAttribute03"`
	FLotAttribute04 string `json:"fLotAttribute04"`
	FLotAttribute05 string `json:"fLotAttribute05"`
	FLotAttribute06 string `json:"fLotAttribute06"`
	FLotAttribute07 string `json:"fLotAttribute07"`
	FLotAttribute08 string `json:"fLotAttribute08"`
	FLotAttribute09 string `json:"fLotAttribute09"`
	FLotAttribute10 string `json:"fLotAttribute10"`
	FIsActive string `json:"fIsActive"`
	FIsSource string `json:"fIsSource"`
	FSKU string `json:"fSKU"`
	FCreateDate string `json:"fCreateDate"`
	FCreateUser string `json:"fCreateUser"`
	FUpdateDate string `json:"fUpdateDate"`
	FUpdateUser string `json:"fUpdateUser"`
}

type ReqCreateShop struct {
	Data []*CreateShop `json:"data"`
}

type CreateShop struct {
	Clientno string `json:"clientno"`				//货主编码
	Code string `json:"code"`
	Name string `json:"name"`						//店铺一
	Status string `json:"status"`
	Platform string `json:"platform"`				//淘宝
	Qrurl string `json:"qrurl"`						//
	Qrimg string `json:"qrimg"`
	Vendorid int `json:"vendorid"`
}

type ResLoadShop struct {

}

type ReqCreateSupplier struct {
	Data []*CreateSupplier `json:"data"`
}

type CreateSupplier struct {
	Clientno string `json:"clientno"`				//
	Code string `json:"code"`						//供应商代码
	Name string `json:"name"`						//供应商名称
	Linkuser string `json:"linkuser"`				//联系人
	Mobile string `json:"mobile"`					//手机号码
	Telephone string `json:"telephone"`				//电话号码
	Fax string `json:"fax"`							//传真号码
	Zip string `json:"zip"`							//邮编号码
	Email string `json:"email"`						//邮箱
	Website string `json:"website"`					//公司网站
	Bankaccount string `json:"bankaccount"`			//银行账号
	Alipay string `json:"alipay"`					//支付宝帐号
	Suppliertypecode string `json:"suppliertypecode"` //供应商类型代码
	Address string `json:"address"`					//地址
	Note string `json:"note"`						//备注
	IsActive string `json:"isActive"`
	Cusattr1 string `json:"cusattr1"`
	Cusattr2 string `json:"cusattr2"`
	Cusattr3 string `json:"cusattr3"`
	Cusattr4 string `json:"cusattr4"`
	Cusattr5 string `json:"cusattr5"`
	Cusattr6 string `json:"cusattr6"`
	Cusattr7 string `json:"cusattr7"`
	Cusattr8 string `json:"cusattr8"`
	Cusattr9 string `json:"cusattr9"`
	Cusattr10 string `json:"cusattr10"`
}

type ReqSelectOperationLog struct {
	Data []*SelectOperationLog `json:"data"`
}

type SelectOperationLog struct {
	OrderNo string `json:"orderNo"`
}

type ResSelectOperationLog struct {
	Body struct {
		Data []*OperationLog `json:"data"`
		ExData struct {} `json:"ExData"`
	} `json:"Body"`
}

type OperationLog struct {
	Ordeno string `json:"ordeno"`		//出库通知单号
	Currno string `json:"currno"`		//当前单据号
	Currtp string `json:"currtp"`		//当前单据类型
	Fromno string `json:"fromno"`		//当前单据的来源单号
	Fromtp string `json:"fromtp"`		//来源单据的单据类型
	Status string `json:"status"`		//单据状态
	Descri string `json:"descri"`		//状态描述
	Aptime string `json:"aptime"`		//确认时间
	Apuser string `json:"apuser"`		//确认人
	Systyp bool `json:"systyp"`
}

type ReqPushOperationLog struct {
	Data []*PushOperationLog `json:"data"`
}

type PushOperationLog struct {
	Shipcode string `json:"shipcode"`
	Dpocode string `json:"dpocode"`
	Opman string `json:"opman"`
	Optime string `json:"optime"`
	Status string `json:"status"`
}

//CreateSupplier