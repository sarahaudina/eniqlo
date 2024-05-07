package main

import (
 "com.eniqlo/app"
)

func main() {
 var a app.App
 a.CreateConnection()
 a.CreateRoutes()
 a.Run()
}