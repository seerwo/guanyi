package stock

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/guanyi/util"
	"github.com/seerwo/guanyi/wms/context"
)

const(
	CREATE_POR_URL = "/WebApi/Api/POR/CreatePOR"
	CREATE_PORN_URL = "/WebApi/Api/PORN/CreatePORN"
	CANCEL_PORN_URL = "/WebApi/Api/PORN/CancelPORN"
	SHIPMENT_FINISHED_URL = "/WebApi/Api/Shipment/ShipmentFinished"
	SHIPMENT_REJECTED_URL = "/WebApi/Api/Shipment/ShipmentRejected"
	SEARCH_SHIPMENTS_URL = "/WebApi/Api/Shipment/SearchShipments"
	UPDATE_SHIPMENT_URL = "/WebApi/Api/Shipment/UpdateShipment"
	SHIPMENT_FINISHED_TO_OVERSEAS_URL = "/WebApi/Api/Shipment/ShipmentFinishedToOverseas"
	CREATE_DN_URL = "/WebApi/Api/DN/CreateDN"
	CREATE_DN_PARALLEL = "/WebApi/Api/DN/CreateDNParallel"
	CANCEL_DN_URL = "/WebApi/Api/DN/CancelDN"
)

type ReqCreatePOR struct {
	Data []*CreatePOR `json:"data"`
}

type CreatePOR struct {
	OrderNo string `json:"orderNo"`
	POReturnKey string `json:"POReturnKey"`
	TotalCategoryQty int `json:"TotalCategoryQty"`
	SupplierId string `json:"supplierId"`
	Warehouse string `json:"warehouse"`
	CreateUserId string `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
	CarrierName string `json:"CarrierName"`
	PredictDeliveryDate string `json:"PredictDeliveryDate"`
	CarrierReferNo string `json:"CarrierReferNo"`
	CreateDate string `json:"createDate"`
	AuditUserId string `json:"auditUserId"`
	AuditUserName string `json:"auditUserName"`
	OperateUser string `json:"OperateUser"`
	Remark string `json:"remark"`
	Tocode string `json:"tocode"`						//采购退货入库标识
	Reason string `json:"reason"`						//退货入库原因：（EndofSeason）换季，（Reject）次品
	ReceiveCompany string `json:"ReceiveCompany"`
	ReceiveName string `json:"ReceiveName"`
	ReceiveZipCode string `json:"ReceiveZipCode"`
	ReceiveEmail string `json:"ReceiveEmail"`
	ReceiveTel string `json:"ReceiveTel"`
	ReceiveMobile string `json:"ReceiveMobile"`
	ReceiveCountryCode string `json:"ReceiveCountryCode"`
	ReceiveProvince string `json:"ReceiveProvince"`
	ReceiveCity string `json:"ReceiveCity"`
	ReceiveArea string `json:"ReceiveArea"`
	ReceiveTown string `json:"ReceiveTown"`
	ReceiveDetailaddress string `json:"ReceiveDetailaddress"`
	OrderType string `json:"OrderType"`
	Detail []CreatePORDetail `json:"detail"`
	FromOrderNo string `json:"FromOrderNo"`
}

type CreatePORDetail struct {
	Sku string `json:"sku"`
	Qty int `json:"qty"`
	Storage string `json:"storage"`
	Price int `json:"price"`
	Clientno string `json:"clientno"`
	Remark string `json:"Remark"`
	Lotkey string `json:"lotkey"`
	Lotattribute01 string `json:"lotattribute01"`
	Lotattribute02 string `json:"lotattribute02"`
	Lotattribute03 string `json:"lotattribute03"`
	Lotattribute06 string `json:"lotattribute06"`
	Lotattribute07 string `json:"lotattribute07"`
	Lotattribute08 string `json:"lotattribute08"`
	Fromdetaillineno string `json:"fromdetaillineno"`
	FReferNo string `json:"fReferNo"`
	FInventoryType string `json:"fInventoryType"`
	Zonecode string `json:"zonecode"`
	Locationcode string `json:"locationcode"`
}

type ReqCreatePORN struct {
	Data []*CreatePORN `json:"data"`
}

type CreatePORN struct {
	OrderNo string `json:"orderNo"`
	POReturnKey string `json:"POReturnKey"`
	TotalCategoryQty int `json:"TotalCategoryQty"`
	SupplierId string `json:"supplierId"`
	Warehouse string `json:"warehouse"`
	CreateUserId string `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
	CarrierName string `json:"CarrierName"`
	PredictDeliveryDate string `json:"PredictDeliveryDate"`
	CarrierReferNo string `json:"CarrierReferNo"`
	CreateDate string `json:"createDate"`
	AuditUserId string `json:"auditUserId"`
	AuditUserName string `json:"auditUserName"`
	OperateUser string `json:"OperateUser"`
	Remark string `json:"remark"`
	Tocode string `json:"tocode"`					//采购退货入库标识
	Reason string `json:"reason"`					//退货入库原因：（EndofSeason）换季，（Reject）次品
	ReceiveCompany string `json:"ReceiveCompany"`
	ReceiveName string `json:"ReceiveName"`
	ReceiveZipCode string `json:"ReceiveZipCode"`
	ReceiveEmail string `json:"ReceiveEmail"`
	ReceiveTel string `json:"ReceiveTel"`
	ReceiveMobile string `json:"ReceiveMobile"`
	ReceiveCountryCode string `json:"ReceiveCountryCode"`
	ReceiveProvince string `json:"ReceiveProvince"`
	ReceiveCity string `json:"ReceiveCity"`
	ReceiveArea string `jsong:"ReceiveArea"`
	ReceiveTown string `json:"ReceiveTown"`
	ReceiveDetailaddress string `json:"ReceiveDetailaddress"`
	OrderType string `json:"OrderType"`
	Detail []*CreatePORNDetail `json:"detail"`
	FromOrderNo string `json:"FromOrderNo"`
}

type CreatePORNDetail struct {
	Sku string `json:"sku"`
	Qty int `json:"qty"`
	Storage string `json:"storage"`
	Price int `json:"price"`
	Clientno string `json:"clientno"`
	Remark string `json:"Remark"`
	Lotkey string `json:"lotkey"`
	Lotattribute01 string `json:"lotattribute01"`
	Lotattribute02 string `json:"lotattribute02"`
	Lotattribute03 string `json:"lotattribute03"`
	Lotattribute06 string `json:"lotattribute06"`
	Lotattribute07 string `json:"lotattribute07"`
	Lotattribute08 string `json:"lotattribute08"`
	Fromdetaillineno string `json:"fromdetaillineno"`
	FReferNo string `json:"fReferNo"`
	FInventoryType string `json:"fInventoryType"`
	Zonecode string `json:"zonecode"`
	Locationcode string `json:"locationcode"`
}

type ReqCancelPORN struct {

}

type ReqShipmentFinished struct {

}

type ReqShipmentRejected struct {
	Data []*ShipmentRejected `json:"data"`
}

type ShipmentRejected struct {
	Shipmentid string `json:"Shipmentid"`		//出库单id
	RejectReason string `json:"RejectReason"`	//拒单原因
}

type ReqSearchShipments struct {
	Type string `json:"type"`				//查询类型：All、None
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	Tocode string `json:"tocode"`
	OrderNos []string `json:"OrderNos"`
}

type ReqUpdateShipment struct {
	Referno string `json:"referno"`				//参考单号
	Buyernick string `json:"buyernick"`			//会员昵称
	Buyername string `json:"buyername"`			//买家姓名
	Countryname string `json:"countryname"`		//国家
	Statename string `json:"statename"`			//省区
	Cityname string `json:"cityname"`			//城市名
	Districtname string `json:"districtname"`	//区/县
	Zip string `json:"zip"`						//邮编
	Address string `json:"address"`				//地址
	Mobile string `json:"mobile"`				//手机
	Telephone string `json:"telephone"`			//固定电话
	Email string `json:"email"`					//邮箱
	Buyermemo string `json:"buyermemo"`			//买家留言
	Receivername string `json:"receivername"`	//收货人姓名
}

type ReqShipmentFinishedToOverseas struct {

}

type ReqCreateDN struct {
	Data []*CreateDN `json:"data"`
}

type ResCreateDN struct {
	util.CommonError
	Result string `json:"Result"`
	CancelReason string `json:"CancelReason"`
	Body struct {
		FailedRecords map[string]string `json:"FailedRecords"`
	} `json:"Body"`
}

type CreateDN struct {
	Whcode string `json:"whcode,omitempty"`				//仓库编码
	Clientno string `json:"clientno,omitempty"`			//
	Ordeno string `json:"ordeno,omitempty"`				//出库通知单号
	Fromty string `json:"fromty,omitempty"`				//来源订单类型
	Ordertype string `json:"ordertype,omitempty"`
	Fromno string `json:"fromno,omitempty"`				//来源单号
	Opform string `json:"opform,omitempty"`				//订单来源平台
	Priori string `json:"priori,omitempty"`				//配货优先级
	Crtime string `json:"crtime,omitempty"`				//下单时间
	Patime string `json:"patime,omitempty"`				//支付时间
	Paymentorderno string `json:"paymentorderno,omitempty"`
	Flightno string `json:"flightno,omitempty"`
	Adtime string `json:"adtime,omitempty"`				//审核时间
	Expdate string `json:"expdate,omitempty"`
	Tocode string `json:"tocode,omitempty"`				//出库单标识
	Aduser string `json:"aduser,omitempty"`				//审核人
	Cruser string `json:"cruser,omitempty"`
	Caarea string `json:"caarea,omitempty"`
	Cacode string `json:"cacode,omitempty"`				//物流公司编码
	Caname string `json:"caname,omitempty"`				//物流公司名称
	Posfee int `json:"posfee,omitempty"`
	Piscod bool `json:"piscod,omitempty"`
	Isdelv bool `json:"isdelv,omitempty"`
	Isgift bool `json:"isgift,omitempty"`
	IsNeedRepackFee bool `json:"IsNeedRepackFee,omitempty"`
	Shopna string `json:"shopna,omitempty"`				//来源店铺名
	Shopno string `json:"shopno,omitempty"`				//来源店铺编号
	Bunick string `json:"bunick,omitempty"`				//会员昵称
	Custna string `json:"custna,omitempty"`				//会员名称
	Recena string `json:"recena,omitempty"`				//收件人姓名
	Postco string `json:"postco,omitempty"`				//邮政编码
	Provco string `json:"provco,omitempty"`				//省编码
	Provna string `json:"provna,omitempty"`				//省名称
	Cityco string `json:"cityco,omitempty"`				//市编码
	Cityna string `json:"cityna,omitempty"`				//市名称
	Distco string `json:"distco,omitempty"`				//区编码
	Distna string `json:"distna,omitempty"`				//区名称
	Townco string `json:"townco,omitempty"`
	Townna string `json:"townna,omitempty"`
	Addres string `json:"addres,omitempty"`				//收货地址
	Mobile string `json:"mobile,omitempty"`				//移动电话
	Teleph string `json:"teleph,omitempty"`				//固定电话
	Sememo string `json:"sememo,omitempty"`				//卖家留言
	Bumemo string `json:"bumemo,omitempty"`				//买家留言
	Shmemo string `json:"shmemo,omitempty"`				//商家留言
	Gymemo string `json:"gymemo,omitempty"`				//订单发票
	Remark string `json:"remark,omitempty"`				//oms用备注
	IsInvoice bool `json:"IsInvoice,omitempty"`
	Invoicedetail struct {} `json:"invoicedetail,omitempty"`
	IsTopay string `json:"IsTopay,omitempty"`				//是否到付
	Ordeca int `json:"ordeca,omitempty"`
	RepackFee string `json:"RepackFee,omitempty"`
	ActuallyPaid int `json:"ActuallyPaid,omitempty"`
	Pretype string `json:"pretype,omitempty"`				//预售类型
	Isspec string `json:"isspec,omitempty"`				//是否个性化包裹
	Specmk string `json:"specmk,omitempty"`				//个性化包裹内容
	Codnum string `json:"codnum,omitempty"`
	Busimode string `json:"busimode,omitempty"`			//业务模式
	Pickmode string `json:"pickmode,omitempty"`			//捡货类型
	IdCardNumber string `json:"IdCardNumber,omitempty"`	//身份证号
	IspickUpOffLine string `json:"ispickUpOffLine,omitempty"` //否把单据给到ssp
	Scheduletype string `json:"scheduletype,omitempty"`
	ScheduleDay string `json:"scheduleDay,omitempty"`
	ScheduleStart string `json:"scheduleStart,omitempty"`
	ScheduleEnd string `json:"scheduleEnd,omitempty"`
	Custom_status string `json:"custom_status,omitempty"`
	Portcode string `json:"portcode,omitempty"`
	IsInsured string `json:"IsInsured,omitempty"`
	FDeclareValue string `json:"fDeclareValue,omitempty"`
	CodServiceFee string `json:"CodServiceFee,omitempty"`
	Note string `json:"note,omitempty"`
	ReservoirArea string `json:"reservoirArea,omitempty"`
	CounterUnit string `json:"counterUnit,omitempty"`
	MasterName string `json:"MasterName,omitempty"`
	PlanOutput string `json:"PlanOutput,omitempty"`
	IsSingle string `json:"IsSingle,omitempty"`
	Skulst []*DNSkuList `json:"skulst,omitempty"`
	SetPackageCode string `json:"setPackageCode,omitempty"`
	SetPackageName string `json:"setPackageName,omitempty"`
	Salesman string `json:"Salesman,omitempty"`
	CombinationProductDetail string `json:"combinationProductDetail,omitempty"`
	OrderTag string `json:"OrderTag,omitempty"`
	Country string `json:"Country,omitempty"`
	ServiceItems string `json:"ServiceItems,omitempty"`
	PrintInfo string `json:"PrintInfo,omitempty"`
	PrintDeliveryInfo string `json:"PrintDeliveryInfo,omitempty"`
	IsNeedTerminateDelivery string `json:"IsNeedTerminateDelivery,omitempty"`
	TerminateDeliveryFee string `json:"TerminateDeliveryFee,omitempty"`
	IsNeedTerminateSetup string `json:"IsNeedTerminateSetup,omitempty"`
	TerminateSetupFee string `json:"TerminateSetupFee,omitempty"`
	TotalDispatchQuantity string `json:"TotalDispatchQuantity,omitempty"`
	LastDeliveryTime string `json:"LastDeliveryTime,omitempty"`
	ThreePLTiming string `json:"ThreePLTiming,omitempty"`
	CnOrderCode string `json:"cnOrderCode,omitempty"`
	FIsJITX string `json:"fIsJITX,omitempty"`
	LogisticsAreaCode string `json:"logisticsAreaCode,omitempty"`
	DistributionItinerary string `json:"distributionItinerary,omitempty"`
	Iscompulsory bool `json:"iscompulsory,omitempty"`
	FVehicleNo string `json:"fVehicleNo,omitempty"`
	fVehicleIndex string `json:"fVehicleIndex,omitempty"`
	JitxOrderType string `json:"jitxOrderType,omitempty"`
	DeliveryKPIStartTime string `json:"deliveryKPIStartTime,omitempty"`
	CodAmount string `json:"codAmount,omitempty"`
}

type DNSkuList struct {
	TradeId string `json:"TradeId,omitempty"`
	Lineno string `json:"lineno,omitempty"`			//行号
	Referlineno string `json:"referlineno,omitempty"`
	Clientno string `json:"clientno,omitempty"`
	Sku string `json:"sku"`					//编码
	ItemId string `json:"itemId,omitempty"`
	Numsku string `json:"numsku,omitempty"`			//平台SKU编码
	Qty int `json:"qty,omitempty"`
	Pri int `json:"pri,omitempty"`
	Amo int `json:"amo,omitempty"`
	Pro string `json:"pro,omitempty"`					//优惠金额
	Pay int `json:"pay,omitempty"`
	Dis int `json:"dis,omitempty"`
	LogisticsCost string `json:"LogisticsCost,omitempty"`
	Cost int `json:"cost,omitempty"`
	Com string `json:"com,omitempty"`					//是否为组合商品
	Nam string `json:"nam,omitempty"`					//商品名称（针对组合商品）
	Pre string `json:"pre,omitempty"`					//是否预售
	Isgift bool `json:"isgift,omitempty"`
	Codets string `json:"codets,omitempty"`			//行邮税号
	LotKey string `json:"lotKey,omitempty"`			//批次号
	LotAttribute01 string `json:"lotAttribute01,omitempty"`
	LotAttribute02 string `json:"lotAttribute02,omitempty"`
	LotAttribute03 string `json:"lotAttribute03,omitempty"`
	LotAttribute04 string `json:"lotAttribute04,omitempty"`
	LotAttribute05 string `json:"lotAttribute05,omitempty"`
	LotAttribute06 string `json:"lotAttribute06,omitempty"`
	LotAttribute07 string `json:"lotAttribute07,omitempty"`
	LotAttribute08 string `json:"lotAttribute08,omitempty"`
	LotAttribute09 string `json:"lotAttribute09,omitempty"`
	LotAttribute10 string `json:"lotAttribute10,omitempty"`
	LocationNo string `json:"locationNo,omitempty"`	//货位
	DispatchProductId string `json:"dispatchProductId,omitempty"`
	StrPlanUse int `json:"strPlanUse,omitempty"`
	Colorcode string `json:"colorcode,omitempty"`
	Colorname string `json:"colorname,omitempty"`
	Desc string `json:"desc,omitempty"`				//备注
	Remarks string `json:"Remarks,omitempty"`
	Stationno string `json:"stationno,omitempty"`
	Ext1 string `json:"ext1,omitempty"`
	Ext2 string `json:"ext2,omitempty"`
	Ext3 string `json:"ext3,omitempty"`
	OutSkuCode string `json:"outSkuCode,omitempty"`	//外仓商品编码
}

// 出库
type ShipmentData struct {
	ShipmentId string `json:"ShipmentId"`
	Cacode string `json:"cacode"`
	Date string `json:"date"`
	Detail []struct {
		PackageDetail string `json:"Packagedetail"`
		MachineCode string `json:"machineCode"`
		Model string `json:"model"`
		PackageCount int `json:"packageCount"`
		StorageLocation string `json:"storageLocation"`
		Qty int `json:"qty"`
		Sku string `json:"sku"`
		Weight float64 `json:"weight"`
	} `json:"detail"`
	DriverTel string `json:"driverTel"`
	Express []struct {
		Code string `json:"code"`
		Number string `json:"number"`
		Weight string `json:"weight"`
	} `json:"express"`
	InvoiceDetail [] struct {
		Id string `json:"id"`
		SkuCode string `json:"skucode"`
		InvoiceTime string `json:"invoicetime"`
		Invoiceno string `json:"invoiceno"`
	} `json:"invoicedetail"`
	TerminateDeliveryFee string `json:"terminatedeliveryfee"`
	TerminateSetupFee string `json:"terminatesetupfee"`
	ExprNo string `json:"exprno"`
	FtoCode string `json:"ftoCode"`
	FromNo string `json:"fromno"`
	Invoice bool `json:"invoice"`
	Invoiceno string `json:"invoiceno"`
	Note string `json:"note"`
	OrderNo string `json:"orderno"`
	RefeNo string `json:"refeno"`
	ScanCode string `json:"scancode"`
	SerialNo string `json:"serialno"`
	Weight float64 `json:"weight"`
	WhCode string `json:"whcode"`
	Type string `json:"type"`
	OutWareId string `json:"outWareId"`

}


type ReqCreateDNParallel struct {
	ReqCreateDN
}

type ReqCancelDN struct {
	Data []*CancelDN `json:"data"`
}

type CancelDN struct {
	Whcode string `json:"whcode"`			//仓库编码
	OwnerCode string `json:"ownerCode"`
	Ordeno string `json:"ordeno"`			//出库通知单号
	Reason string `json:"reason"`			//取消原因
}

//Out struct
type Out struct {
	*context.Context
}

//NewOut instance
func NewOut(context *context.Context) *Out {
	out := new(Out)
	out.Context = context
	return out
}

//CreateDN
func (out *Out) GetCreateDN(req ReqCreateDN) (res ResCreateDN, err error){
	accessParam := ""
	if accessParam, err = out.GetAccessParam(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s%s", util.WMS_WEB_URL + CREATE_DN_URL , accessParam)

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
	if !res.IsSuccess && res.Result == "false" {
		err = fmt.Errorf("GetCreateDN Error , errcode=%d , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
//cancelDN