package main

import n "awesomeProject/lcffilemanager/net"
func main()  {
	//m.ListDir(".")
	r := n.InitConnection()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static/","./static")
	if err:=r.Run("0.0.0.0:80");err!=nil{}else{panic(err)}
}
//rjx宁也来辣？(　^ω^)
//那就来一起学学吧。