package main

import "fmt"

//当项目的规模增大后就不太适合继续在项目的 main.go 文件
//中去实现路由注册相关逻辑了，我们会倾向于把路由部分的代码都拆分出来，形成一个单独的文件或包：
func main()  {
	r := setupRouter()
	if err := r.Run(); err != nil{
		fmt.Println("startup service failed, err:%v\n",err)
	}
	
}