package server

import (
	"encoding/json"
	"fmt"
	"fraise/config"
	"io/ioutil"
	"log"
	"sync/atomic"
)

//节点
type node struct {
	Url    string `json:"url"`
	Weight int    `json:"weight"`
}

//节点权重分配列表
var weightList atomic.Value

// InitList 初始化服务节点列表
func InitList() {

	//加载本地数据
	localData, err := ioutil.ReadFile("./" + config.Conf.Server.ListFile)
	if err != nil {
		panic(err)
	}
	//如果本地数据为空
	if localData == nil || len(localData) == 0 {
		//将本地数据设为空数组
		localData = []byte("[]")
		err = ioutil.WriteFile("./"+config.Conf.Server.ListFile, localData, 0766)
		if err != nil {
			panic(err)
		}
	}

	//获取服务节点列表
	var localList []node
	err = json.Unmarshal(localData, &localList)
	if err != nil {
		panic(err)
	}
	fmt.Println("Node list:", localList)

	//将服务节点列表按权重载入内存
	var wl []string
	for _, v := range localList {
		for i := 0; i < v.Weight; i++ {
			wl = append(wl, v.Url)
		}
	}
	weightList.Store(wl)
}

//获取权重分配列表
func getWeightList() []string {
	return weightList.Load().([]string)
}

// GetNodeList 获取服务节点列表
func GetNodeList() []node {
	//加载本地数据
	localData, err := ioutil.ReadFile("./" + config.Conf.Server.ListFile)
	if err != nil {
		log.Println(err)
	}

	//获取服务节点列表
	var localList []node
	err = json.Unmarshal(localData, &localList)
	if err != nil {
		log.Println(err)
	}
	return localList
}

// SetNodeList 更新服务节点列表
func SetNodeList(data []byte) string {

	err := ioutil.WriteFile("./"+config.Conf.Server.ListFile, data, 0766)
	if err != nil {
		log.Println(err)
		return "err"
	}
	//重新载入内存
	InitList()
	return "ok"
}
