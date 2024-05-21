package stock

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/guanyi/util"
	"github.com/seerwo/guanyi/wms/context"
)

const(
	CREATE_ADJUSTMENT_URL = "/WebApi/Api/Adjustment/CreateAdjustment"
	CANCEL_ADJUSTMENT_URL = "/WebApi/Api/Adjustment/CancelAdjustment"
	INSERT_CYCLE_COUNT_BY_SKU_URL = "/WebApi/Api/CycleCount/InsertCycleCountBySku"
	SEARCH_INVENTORY_URL = "/WebApi/Api/Inventory/SearchInventory"
	INVENTORY_CHANGED_URL = "/WebApi/Api/Inventory/InventoryChanged"
	SEARCH_SKU_ARRAIL_INFO_2_URL = "/WebApi/Api/Inventory/SearchSkuArrailInfo2"
	SEARCH_INVENTORY_2_URL = "/WebApi/Api/Inventory/SearchInventory2"
	CREATE_TRANSFER_URL = "/WebApi/Api/Transfer/CreateTransfer"
	CANCEL_TRANSFER_URL = "/WebApi/Api/Transfer/CancelTransfer"

)

type ReqCreateAdjustment struct {
	Data []*ReqCreateAdjustment `json:"data"`
}

type CreateAdjustment struct {
	OrderNo string `json:"OrderNo"`
	StorerNo string `json:"StorerNo"`
	WarehouseNo string `json:"WarehouseNo"`
	Tocode string `json:"tocode"`
	Memo string `json:"Memo"`				//备注
	ResonCode string `json:"ResonCode"`
	Details []*CreateAdjustmentDetail `json:"Details"`
}

type CreateAdjustmentDetail struct {
	LocationNo string `json:"LocationNo"`
	SKU string `json:"SKU"`
	Qty int `json:"Qty"`
	Description string `json:"Description"`
	FromLotAttribute01 string `json:"FromLotAttribute01"`
	FromLotAttribute02 string `json:"FromLotAttribute02"`
	FromLotAttribute07 string `json:"FromLotAttribute07"`
	FromLotAttribute08 string `json:"FromLotAttribute08"`
}

type ReqCancelAdjustment struct {
	Data []CancelAdjustment `json:"data"`
}

type CancelAdjustment struct {
	OrderNo string `json:"OrderNo"`
}

type ReqInsertCycleCountBySku struct {
	WarehouseNo string `json:"WarehouseNo"`
	StorerNo string `json:"StorerNo"`
	ZoneType string `json:"ZoneType"`			//Pick
	Skus string `json:"Skus"`					//Sku01,Sku02
	ReferNo string `json:"ReferNo"`				//参考单号
}

type ReqSearchInventory struct {
	Data []*SearchInventory `json:"data"`
}

type ResSearchInventory struct {
	util.CommonError
	Type string `json:"type"`
	Result struct {
		Code string `json:"code"`
		Reason string `json:"reason"`
		Orders struct {} `json:"orders"`
	} `json:"Result"`
	Response []*SearchInventoryData `json:"response"`
	ExData struct{} `json:"ExData"`
}

type SearchInventoryData struct {
	Sku string `json:"sku"`
	Qty int `json:"qty"`
	Avlqty int `json:"avlqty"`
	GoodStatus string `json:"GoodStatus"`
	WarehouseNo string `json:"WarehouseNo"`
	LotAttribute02 string `json:"LotAttribute02"`
	LotAttribute07 string `json:"LotAttribute07"`
	LotAttribute08 string `json:"LotAttribute08"`
}

type SearchInventory struct {
	Whcode string `json:"whcode"`			//仓库编号
	Clientno string `json:"clientno"`		//
	Skucode string `json:"skucode"`			//SKU01,SKU02
}

type ReqInventoryChanged struct {

}

type ReqSearchSkuArrailInfo2 struct {
	Data []*SearchSkuArrailInfo2 `json:"data"`
}

type SearchSkuArrailInfo2 struct {
	Whcode string `json:"whcode"`
	Sku string `json:"sku"`
}

type ReqSearchInventory2 struct {
	PageIndex int `json:"PageIndex"`
	PageSize int `json:"PageSize"`
}

type ReqCreateTransfer struct {
	Data []*CreateTransfer `json:"data"`
}

type CreateTransfer struct {
	asnorder string `json:"asnorder"`			//调拨单号
	fromorder string `json:"fromorder"`			//来源单号
	headeruser string `json:"headeruser"`		//经手人编码
	ClientNo string `json:"ClientNo"`
	deptcode string `json:"deptcode"`			//部门编码
	deptname string `json:"deptname"`			//部门名称
	fromwhcode string `json:"fromwhcode"`		//调出仓库编码
	targetwhcode string `json:"targetwhcode"`	//调入仓库编号
	operuserid string `json:"operuserid"`		//创建人编码
	operusername string `json:"operusername"`	//创建人名称
	operdate string `json:"operdate"`			//创建日期(yyyy/MM/dd HH:mm:ss)
	ordertype string `json:"ordertype"`			//业务类型(1:借出2:归还3:核销)
	goodstype string `json:"goodstype"`			//商品类型(0:普通,1：损耗 2：残品)
	remark string `json:"remark"`				//备注
	tocode string `json:"tocode"`
	MasterName string `json:"MasterName"`
	PlanOutput int `json:"PlanOutput"`
	Detail []*CreateTransferDetail `json:"detail"`
}

type CreateTransferDetail struct {
	Sku string `json:"sku"`
	Qty int `json:"qty"`
	ClientNo string `json:"ClientNo"`
}

type ReqCancelTransfer struct {
	Data []*CancelTransfer `json:"data"`
}

type CancelTransfer struct {
	OrderNo string `json:"orderNo"`				//原单号
	Reason string `json:"reason"`				//取消原因
}

//入库
type QualityData struct {
	ApTime string `json:"aptime"`
	ApUser string `json:"apuser"`
	Descri string `json:"descri"`
	FWmsOrderNo string `json:"fWmsOrderNo"`
	FromNo string `json:"fromno"`
	FromTy string `json:"fromty"`
	GoodsStatus string `json:"goodsstatus"`
	Id string `json:"id"`
	IsGift bool `json:"isgift"`
	IsOnSale bool `json:"isonsale"`
	ReceiptNo string `json:"receiptNo"`
	WhCode string `json:"whcode"`
	Qty int `json:"qty"`
	Sku string `json:"sku"`
	Sn string `json:"sn"`
	Weight int `json:"weight"`
	IsAccept bool `json:"isaccept"`
	CreateTime string `json:"createTime"`
	FromLotKey string `json:"fromlotkey"`
	ToLotKey string `json:"tolotkey"`
}

//Stock struct
type Stock struct {
	*context.Context
}

//NewOut instance
func NewStock(context *context.Context) *Stock {
	stock := new(Stock)
	stock.Context = context
	return stock
}

//SerachInventory
func (out *Out) GetSearchInventory(req ReqSearchInventory) (res ResSearchInventory, err error){
	accessParam := ""
	if accessParam, err = out.GetAccessParam(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s", util.WMS_WEB_URL + SEARCH_INVENTORY_URL , accessParam)
	jsonBytes, err := json.Marshal(req.Data)
	if err != nil {
		return
	}
	var response []byte
	if response, err = util.NewHTTPPost(uri, string(jsonBytes)); err != nil {
		return
	}
	if err = json.Unmarshal(response,&res); err != nil {
		return
	}

	if !res.IsSuccess {
		err = fmt.Errorf("GetCreateDN Error , errcode=%d , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}