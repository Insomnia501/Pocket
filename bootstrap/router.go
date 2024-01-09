package bootstrap

import "pocket-serv/routes"

// RunServer 启动服务器
func RunServer() {
	r := routes.InitRouter()
	r.Run(":8080")
}
