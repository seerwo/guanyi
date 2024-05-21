package good

const(
	CREATE_SKU_URL = "/WebApi/Api/SKU/CreateSKU"
)


type ReqCreateSKU struct {
	Data []*CreateSKU `json:"data"`
}

type CreateSKU struct {
	Owcode string `json:"owcode"`			//
	Owname string `json:"owname"`			//货主名称
	Cacode string `json:"cacode"`			//类目代码
	Itemno string `json:"itemno"`			//商品货号
	Itemna string `json:"itemna"`			//商品名称
	Shname string `json:"shname"`			//商品简称
	Unitcode string `json:"unitcode"`		//主单位
	Unitname string `json:"unitname"`		//主单位
	Price int `json:"price"`
	Branid string `json:"branid"`			//品牌代码
	Brandname string `json:"brandname"`		//品牌名称
	Season string `json:"season"`			//季节代码
	Ducode string `json:"ducode"`			//负责人代码
	Duname string `json:"duname"`			//负责人名称
	Hocode string `json:"hocode"`			//库存承担人代码
	Itstat string `json:"itstat"`			//状态
	Intype string `json:"intype"`			//商品属性
	Clientno string `json:"clientno"`		//货主代码
	isprivacy string `json:"isprivacy"`
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
	Cusattr11 string `json:"cusattr11"`
	Cusattr12 string `json:"cusattr12"`
	Cusattr13 string `json:"cusattr13"`
	Cusattr14 string `json:"cusattr14"`
	Cusattr15 string `json:"cusattr15"`
	Cusattr16 string `json:"cusattr16"`
	Cusattr17 string `json:"cusattr17"`
	Cusattr18 string `json:"cusattr18"`
	Cusattr19 string `json:"cusattr19"`
	Cusattr20 string `json:"cusattr20"`
	Imageurl string `json:"imageurl"`			//商品图片地址
	Properties string `json:"properties"`		//附加属性
	Cruser string `json:"cruser"`
	DiscountPolicy string `json:"discountPolicy"` //优惠政策,来自OMS系统优惠政策
	Description string `json:"description"`		//备注,来自OMS系统备注1
	Detail []*SkuDetail `json:"detail"`
}

type SkuDetail struct {
	Itemno string `json:"itemno"`		//商品货号
	Stcode string `json:"stcode"`		//规格代码(sku)
	Stname string `json:"stname"`		//规格名称
	Brcode string `json:"brcode"`		//条形码
	Color string `json:"color"`			//颜色名称
	Size string `json:"size"`			//尺寸大小
	Length int `json:"length"`
	Wide int `json:"wide"`
	High int `json:"high"`
	Volume int `json:"volume"`
	Validdays int `json:"validdays"`
	Weight int `json:"weight"`
	Gweight int `json:"gweight"`
	Instat string `json:"instat"`			//状态
	FirstPrice string `json:"FirstPrice"`
	CostPrice string `json:"CostPrice"`
	ProductModel string `json:"productModel"`
	DosageForm string `json:"dosageForm"`	//剂型
}