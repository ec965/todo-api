package main

import (
	"github.com/ec965/todo-api/routers"
	"github.com/ec965/todo-api/models"
)

func main() {
	models.InitDB();
	r := routers.InitRouters();
	r.Run();
}
