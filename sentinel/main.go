package main

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/util"
	"log"
	"math/rand"
	"time"

	"github.com/alibaba/sentinel-golang/core/flow"

	sentinel "github.com/alibaba/sentinel-golang/api"
)

func main() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Error to init sentinel, err=%v", err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{ // 基于QPS对某个资源限流
			Resource:               "some-test1",
			TokenCalculateStrategy: flow.Direct, // 直接模式
			ControlBehavior:        flow.Reject, // 直接拒绝
			Threshold:              500,         // 允许多少个
			StatIntervalInMs:       1000,        // 多长时间内
		},
		{ // 基于一定统计间隔时间来控制总的请求数
			Resource:               "some-test2",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              10000,
			StatIntervalInMs:       10000,
		},
		{
			Resource:               "some-test3",
			TokenCalculateStrategy: flow.WarmUp, // 冷启动/预热方式 缓慢增长访问量
			ControlBehavior:        flow.Reject, // 直接拒绝
			Threshold:              1000,        // 允许多少个
			WarmUpPeriodSec:        10,          // 多长时间内达到顶峰
			WarmUpColdFactor:       3,           // 预热因子，默认3
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 埋点逻辑，埋点资源名为 some-test
				e, b := sentinel.Entry("some-test")
				if b != nil {
					// 请求被拒绝，在此处进行处理
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					// 请求允许通过，此处编写业务逻辑
					fmt.Println(util.CurrentTimeMillis(), "Passed")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)

					// 务必保证业务结束后调用 Exit
					e.Exit()
				}

			}
		}()
	}
	<-ch
}
