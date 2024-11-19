package main

import "go-restful-blog-api/v2/config"

func main() {
	config.ConnectDB()
	config.TestDBConnection()
}
