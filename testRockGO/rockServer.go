package main

import (
	"github.com/zllangct/RockGO"
	"github.com/zllangct/RockGO/component"
	"github.com/zllangct/RockGO/gate"
	"github.com/zllangct/RockGO/logger"
)

var Server *RockGO.Server

func main()  {
	//初始化服务节点
	Server = RockGO.DefaultServer()

	/*
	   添加组件组
	   添加网关组件（DefaultGateComponent）后，此服务节点拥有网关的服务能力。
	   同理，添加其他组件，如登录组件（LoginComponent）后，拥有登录的服务内容。
	*/
	Server.AddComponentGroup("gate",[]Component.IComponent{&gate.DefaultGateComponent{}})

	//开始服务
	Server.Serve()
}