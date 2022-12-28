package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jassue-gin/app/models"
	"jassue-gin/global"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/x-funs/go-fun"
)

type messageService struct {
}

var MessageService = new(messageService)

// 获取黑名单缓存 key
func (messageService *messageService) Jijin() string {
	url2 := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=4ff11db0-25fc-4525-a187-2b09afdc17e3"
	meurl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=c762d7c9-f28f-474c-901b-41645a0b0a56"
	var codes = make([]string, 0)
	global.App.DB.Model(models.Data{}).Where("type", 1).Select("code").Scan(&codes)
	var code = strings.Join(codes, ",")
	requrl := "https://api.doctorxiong.club/v1/fund?code=" + code
	data, _ := fun.HttpGet(requrl)
	js := string(data)
	res := gjson.Get(js, "data").String()
	var jsonData = []byte(res)
	var b []models.Jijin
	_ = json.Unmarshal(jsonData, &b)
	for i := 0; i < len(b); i++ {
		message := "{\n" +
			"    \"msgtype\": \"markdown\",\n" +
			"    \"markdown\": {\n" +
			"        \"content\": \"今日基金反馈<font color=\\\"warning\\\"></font>\\n\n" +
			"         >类型:<font color=\\\"comment\\\">基金</font>\n" +
			"         >基金代码:<font color=\\\"comment\\\">" + b[i].Code + "</font>\n" +
			"         >基金名称:<font color=\\\"comment\\\">" + b[i].Name + "</font>\n" +
			"         >今日价格:<font color=\\\"comment\\\">" + b[i].DayGrowth + "元" + "</font>\n" +
			"         >今日涨跌幅:<font color=\\\"comment\\\">" + fun.ToString(b[i].NetWorth) + "%" + "</font>\"\n" +
			"    }\n" +
			"}"
		req, err := http.NewRequest("POST", meurl, bytes.NewBuffer([]byte(message)))
		_, _ = http.NewRequest("POST", url2, bytes.NewBuffer([]byte(message)))
		// req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("status", resp.Status)
		fmt.Println("response:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
	}
	return "基金消息推送成功"
}
