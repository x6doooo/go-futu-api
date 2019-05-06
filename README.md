
### Example

```go
package main

import (
	"fmt"
	gf "github.com/x6doooo/go-futu-api"
	"github.com/x6doooo/go-futu-api/pb/Qot_Common"
	"log"
	"time"
)

func main() {
	// 初始化opend的配置
	cli := gf.NewClient("0.0.0.0:31415")

	// 启动
	err := cli.Run()
	if err != nil {
		panic(err)
	}

	// 建立tcp连接
	resp, err := cli.InitConnect()
	if err != nil {
		fmt.Println(err)
	}

	// 循环发送keepalive请求
	keepAliveInterval := resp.S2C.KeepAliveInterval
	go func() {
		for {
			d := time.Duration(int64(*keepAliveInterval)) * time.Second
			time.Sleep(d)
			cli.KeepAlive()
		}
	}()

	// 获取美股所有板块信息
	respGetPlateSet, err := cli.GetPlateSet(
		Qot_Common.QotMarket_QotMarket_US_Security,
		Qot_Common.PlateSetType_PlateSetType_All,
	)

	for idx, item := range respGetPlateSet.S2C.PlateInfoList {
		log.Println("plate:", idx, *item.Plate.Market, *item.Plate.Code, *item.Name)
	}

	// 获取nasdaq所有股票列表
	market := int32(11)
	code := "NASDAQ"
	respGetPlateSecurity, err := cli.GetPlateSecurity(&Qot_Common.Security{
		Market: &market,
		Code:   &code,
	})

	for idx, item := range respGetPlateSecurity.S2C.StaticInfoList {
		log.Println("security:", idx, *item.Basic.Name, *item.Basic.Security.Code)
	}


}

```

* Credit
This package is heavily inspired by [[https://github.com/FutunnOpen/py-futu-api][py-futu-api]] for Python & [[https://github.com/yongxiongwei/futu][futu]]