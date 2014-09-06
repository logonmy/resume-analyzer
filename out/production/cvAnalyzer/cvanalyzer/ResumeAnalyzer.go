package main

import (
	"fmt"
	"analysis"
	"net/http"
	"controller"
	"time"
	"redis"
)

func startWeb() {
	fmt.Println("初始化词库...")
	analysis.InitConf()
	fmt.Println("加载路由信息...")
	http.HandleFunc("/analysis", controller.AnalysisController)
	http.HandleFunc("/submit", controller.SubmitController)
	http.HandleFunc("/resume", controller.ReadResumeController)
	fmt.Println("启动web服务...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("启动错误：", err)
	}
}

func devTest() {
	analysis.InitConf()
	resume := analysis.Analysis(analysis.Read("51.txt"))
	fmt.Println(resume)
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("++++++++++ 3 +++++++++++")
}

func redisSet(key string, value string) {

	spec := redis.DefaultSpec().Host("172.17.0.56").Db(0).Password("");
	client, e := redis.NewSynchClientWithSpec(spec);
	if e != nil { fmt.Println("服务器连接有异常", e); return }
	inva := []byte(value)
	fmt.Println(inva)
	e=client.Set(key, inva);
	if e != nil { fmt.Println("添加数据错误", e); return }
	fmt.Println(key, "添加成功")

}
func redisGet(key string) string {

	spec := redis.DefaultSpec().Host("172.17.0.56").Db(0).Password("");
	client, e := redis.NewSynchClientWithSpec(spec);
	if e != nil { fmt.Println("服务器连接有异常", e); return "" }
	content, err := client.Get(key)
	if err != nil {
		fmt.Println("获取错误", err);
		return ""
	}
	client.Quit()
	fmt.Println(content)
	return string(content)
}

func main() {
	//redisSet("wzg", "fdsfds fjkdsjfs ")
	//redisSet("wzg", "fdsfd 为是老魏123 ")
	//content:=redisGet("wzg")
	//fmt.Println(content)
	//time.Sleep(time.Duration(3) * time.Second)
	//fmt.Println("++++++++++++")
	//redisGet("wzg")
	//startWeb()
	devTest()
//	fmt.Println(string([]rune{26377}))
//	fmt.Println(string([]rune{32447,20256, 36755, 24037, 31243, 24072}))
//	fmt.Println(string([]rune{32447, 20256, 36755, 24037, 31243, 24072}))
	//aaa:=[]rune("  　•/\f\v")
//	aaa := []rune(`智联招聘
//
//公司描述：
//
//公司性质：
//中外合营(合资·合作)`)
//	fmt.Println(aaa)
	fmt.Println([]rune("()（）"))
}



