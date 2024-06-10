package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/mbocsi/file-based-router/router"
)

func main() {
	rootRoute := flag.String("root", "routes", "The root route directory")
	port := flag.Int("port", 8080, "Listening port")
	flag.Parse()
	root := path.Clean(*rootRoute)
	addr := fmt.Sprintf(":%v", *port)

	handler := router.NewFBHandler(root)
	router := router.NewFBRouter(addr, handler)
	fmt.Printf("Running on %v\nServing files from %v\n", addr, root)

	err := router.Run()
	if err != nil {
		fmt.Println("Error when running router")
	}
}
