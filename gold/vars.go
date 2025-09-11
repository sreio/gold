package gold

import "github.com/sreio/gold/web/model"

type GoldType interface {
	GetGoldPrice() GoldType // 获取金价
	SetTemplate() string    // 设置模板数据
}

var GoldMap = map[string]GoldType{
	model.ICBCSHOPPRICE: &IcbcShopPrice{
		Method: "POST",
		Url:    "https://mybank.icbc.com.cn/servlet/AsynGetDataServlet",
	},
	model.ICBCBANKRUYIPRICE: &IceBankRuyiPrice{
		Method: "GET",
		Url:    "https://mall.icbc.com.cn/products/queryProdPriceAjax.jhtml?productId=9003867817",
	},
}
