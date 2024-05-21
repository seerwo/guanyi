package invoice

const(
	CREATE_INVOICE_URL = "/WebApi/Api/Invoice/CreateInvoice"
	SEARCH_INVOICE_URL = "/WebApi/Api/Invoice/SearchInvoice"
	PUSH_STATUS_FINISHED_URL = "/WebApi/Api/Invoice/PushStatusFinished"
	SET_INVOICE_NO_URL = "/WebApi/Api/Invoice/PushStatusFinished"
	RESET_PUSH_STATUS_URL = "/WebApi/Api/Invoice/ResetPushStatus"
)

type ReqCreateInvoice struct {
	Data []CreateInvoice `json:"data"`
}

type CreateInvoice struct {
	Invoiceid string `json:"invoiceid"`
	Orderid string `json:"orderid"`
	Orderkey string `json:"orderkey"`				//来源单据的单号
	Invoicetitle string `json:"invoicetitle"`		//发票抬头
	Invoicecontent string `json:"invoicecontent"`	//发票内容
	Invoicetotal int `json:"invoicetotal"`
	Postfee int `json:"postfee"`
	Invoicetypecode string `json:"invoicetypecode"` //发票类型
	Invoiceno string `json:"invoiceno"`				//发票号
	Accountnumber string `json:"accountnumber"`		//银行账号
	Bank string `json:"bank"`						//开户银行
	Memo string `json:"memo"`						//备注
	FIsPrintInvoice bool `json:"fIsPrintInvoice"`
	Invoicedetail []*Invoicedetail `json:"invoicedetail"`
}

type Invoicedetail struct {
	Invoicedetailid string `json:"invoicedetailid"`
	Invoiceid string `json:"invoiceid"`
	Skuid string `json:"skuid"`
	Sku string `json:"sku"`				//商品编码
	Skuname string `json:"skuname"`		//商品名称
	Amount int `json:"amount"`
	Productno string `json:"productno"`	//产品货号
	Model string `json:"model"`			//款式型号
	Orderedqty int `json:"orderedqty"`
	Shippedqty string `json:"shippedqty"` //已出库数量
	Allocatedqty string `json:"allocatedqty"` //已分配数量
	Pickedqty string `json:"pickedqty"`	//已拣货数量
	Unitname string `json:"unitname"`	//单位名称
	Price int `json:"price"`
	Discountamount int `json:"discountamount"`
	Promotionamount int `json:"promotionamount"`
	Barcode string `json:"barcode"`		//条码
	Colorname string `json:"colorname"`	//颜色
	Sizename string `json:"sizename"`	//尺码名称
	Isgift bool `json:"isgift"`
	Grossweight int `json:"grossweight"`
	Netweight int `json:"netweight"`
	Cube int `json:"cube"`
	Description string `json:"description"` //说明
}

type ReqSearchInvoice struct {
	OrderKeys []string `json:"orderKeys"`
	PageNo string `json:"page_no"`
	PageSize string `json:"page_size"`
	StartCreate string `json:"start_create"`
	EndCreate string `json:"end_create"`
	Code string `json:"code"`
	NoticeCode string `json:"notice_code"`
	OuterCode string `json:"outer_code"`
	WarehouseCode string `json:"warehouse_code"`
	ShopCode string `json:"shop_code"`
}

type ReqPushStatusFinished struct {
	Data []string `json:"data"`
}

type ReqSetInvoiceNo struct {
	Orders []SetInvoiceNo `json:"Orders"`
}

type SetInvoiceNo struct {
	OrderKey string `json:"OrderKey"`
	InvoiceNo string `json:"InvoiceNo"`
}

type ReqResetPushStatus struct {
	Data []string `json:"data"`
}
