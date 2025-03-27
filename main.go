package main

import (
	"github.com/danfelab/youthcongressnepal/connect"
	"github.com/danfelab/youthcongressnepal/server"
)

func main(){
	
	connect.DB();

	server.StartServer();	

}