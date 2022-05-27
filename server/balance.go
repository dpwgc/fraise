package server

import (
	"fmt"
	"fraise/config"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

//计数器
var cnt atomic.Value

// Forward 转发请求
func Forward(w http.ResponseWriter, r *http.Request) {

	//更改请求url
	u, err := url.Parse(fmt.Sprintf("%s%s", balance(config.Conf.Server.Policy), r.RequestURI))
	//fmt.Println(u)
	//如果出错
	if err != nil {
		//跳过，换下一个url
		log.Println(err)
		return
	}

	proxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL = u
		},
	}

	//转发完毕，跳出循环
	proxy.ServeHTTP(w, r)
}

//负载均衡
func balance(policy uint8) string {
	switch policy {

	//加权随机策略
	case 0:
		return randomLB()

	//加权轮询策略
	case 1:
		return weightLB()

	default:
		return weightLB()
	}
}

//加权随机策略
func randomLB() string {
	list := getWeightList()
	return list[rand.Intn(len(list))]
}

//加权轮询策略
func weightLB() string {

	//如果cnt为空，则进行初始化
	if cnt.Load() == nil {
		cnt.Store(0)
	}

	//获取权重列表
	list := getWeightList()

	//如果当前计数器超过列表上限，则将计数器归0
	i := cnt.Load().(int)
	if i >= len(list) {
		i = 0
		cnt.Store(i)
	}

	//计数器累加
	cnt.Store(i + 1)

	//返回url
	return list[i]
}
