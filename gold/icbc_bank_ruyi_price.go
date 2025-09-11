package gold

import (
	"encoding/json"
	"fmt"
	"github.com/sreio/gold/tools"
	"strconv"
	"time"
)

type IceBankRuyiPrice struct {
	Url       string
	Method    string
	G5Price   string
	G10Price  string
	G20Price  string
	G50Price  string
	G100Price string
	G200Price string
	SysTime   string
}

var KPriceMap = map[int]string{
	5:   "90000000000031379283",
	10:  "90000000000031379284",
	20:  "90000000000031379285",
	50:  "90000000000031379286",
	100: "90000000000031379287",
	200: "90000000000031379288",
}

const RuyiTpl = `
Api同步时间：%s

---

|来源| 类型 | 价格 |
|:---:|:---:|:-----------:|
|ICBC| 5g  |%s￥|
|ICBC| 10g |%s￥|
|ICBC| 20g |%s￥|
|ICBC| 50g |%s￥|
|ICBC| 100g |%s￥|
|ICBC| 200g |%s￥|

`

func (i *IceBankRuyiPrice) GetGoldPrice() GoldType {

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
	}
	response, code, err := tools.HTTPRequest(i.Method, i.Url, headers, nil)
	//fmt.Println("response:", string(response))

	if err != nil {
		_ = fmt.Errorf("错误码：%d HTTP请求错误: %v", code, err)
		return nil
	}

	var outerData struct {
		ProdSkuPriceJson string `json:"prodSkuPriceJson"`
		SysTime          string `json:"sysTime"`
	}
	if err := json.Unmarshal(response, &outerData); err != nil {
		_ = fmt.Errorf("解析JSON错误: %v", err)
		return nil
	}

	// 金价
	var skuPriceResponse map[string]string
	if err := json.Unmarshal([]byte(outerData.ProdSkuPriceJson), &skuPriceResponse); err != nil {
		_ = fmt.Errorf("解析JSON错误: %v", err)
		return nil
	}

	sysTime, _ := strconv.ParseInt(outerData.SysTime, 10, 64)
	if err != nil {
		return nil
	}

	return &IceBankRuyiPrice{
		G5Price:   skuPriceResponse[KPriceMap[5]],
		G10Price:  skuPriceResponse[KPriceMap[10]],
		G20Price:  skuPriceResponse[KPriceMap[20]],
		G50Price:  skuPriceResponse[KPriceMap[50]],
		G100Price: skuPriceResponse[KPriceMap[100]],
		G200Price: skuPriceResponse[KPriceMap[200]],
		SysTime:   time.Unix(sysTime/1000, 0).Format(time.DateTime),
	}
}

func (i *IceBankRuyiPrice) SetTemplate() string {
	return fmt.Sprintf(RuyiTpl, i.SysTime, i.G5Price, i.G10Price, i.G20Price, i.G50Price, i.G100Price, i.G200Price)
}
