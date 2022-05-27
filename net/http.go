package net

import (
	"encoding/json"
	"fraise/config"
	"fraise/server"
	"github.com/unrolled/secure"
	"log"
	"net/http"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// 负载均衡器控制台请求 Set 更新负载均衡器本地存储的节点列表
	if r.RequestURI == config.Conf.Server.ConsoleApi.Set {
		_, err := w.Write([]byte(server.SetNodeList([]byte(r.FormValue("nodeList")))))
		if err != nil {
			log.Println(err)
		}
		return
	}

	// 负载均衡器控制台请求 Get 获取负载均衡器的本地节点列表
	if r.RequestURI == config.Conf.Server.ConsoleApi.Get {

		data, err := json.Marshal(server.GetNodeList())
		if err != nil {
			log.Println(err)
			return
		}

		_, err = w.Write(data)
		if err != nil {
			log.Println(err)
		}
		return
	}

	//负载均衡转发请求
	server.Forward(w, r)
})

// InitHttp 启动http服务
func InitHttp() {

	//如果不开启HTTPS
	if len(config.Conf.Server.HttpsPort) == 0 {
		http.HandleFunc("/", server.Forward)
		err := http.ListenAndServe("127.0.0.1:"+config.Conf.Server.HttpPort, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	//如果开启HTTPS
	secureMid := secure.New(secure.Options{
		SSLRedirect: true,
		SSLHost:     config.Conf.Server.HttpsAddr + ":" + config.Conf.Server.HttpsPort, //http重定向到https
	})

	h := secureMid.Handler(handler)

	go func() {
		err := http.ListenAndServe("127.0.0.1:"+config.Conf.Server.HttpPort, h)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	log.Fatal(http.ListenAndServeTLS(":"+config.Conf.Server.HttpsPort, config.Conf.Server.CertFile, config.Conf.Server.KeyFile, h))
}
