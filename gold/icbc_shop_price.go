package gold

import (
	"encoding/json"
	"fmt"
	"github.com/sreio/gold/tools"
	"strconv"
)

type IcbcShopPrice struct {
	Url    string
	Method string
	Price  float64
}

type ICBCShopResponse struct {
	Pronoinfo []struct {
		Buyprice string `json:"buyprice"`
	} `json:"pronoinfo"`
}

const ShopTpl = `工商如意金：%.2f`

func (i *IcbcShopPrice) GetGoldPrice() GoldType {
	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	response, code, err := tools.HTTPRequest(i.Method, i.Url, headers, []byte("tranCode=A00505"))
	//fmt.Println("response:", string(response))

	if err != nil {
		_ = fmt.Errorf("错误码：%d HTTP请求错误: %v", code, err)
		return nil
	}

	var icbcResponse ICBCShopResponse
	if err := json.Unmarshal(response, &icbcResponse); err != nil {
		_ = fmt.Errorf("解析JSON错误: %v", err)
		return nil
	}

	var price float64
	// 提取并转换金价
	if len(icbcResponse.Pronoinfo) > 0 {
		buyPrice := icbcResponse.Pronoinfo[0].Buyprice
		price, _ = strconv.ParseFloat(buyPrice, 64)
		price /= 100
	}
	return &IcbcShopPrice{
		Price: price,
	}
}

func (i *IcbcShopPrice) SetTemplate() string {
	return fmt.Sprintf(ShopTpl, i.Price)
}
