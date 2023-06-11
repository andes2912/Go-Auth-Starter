package main

import (
	"starterkit-go-auth/initializer"
	"starterkit-go-auth/routes"
)

func init()  {
	initializer.LoadEnvVariables()
	initializer.KoneksiDB()
	initializer.Migration()
}

func main()  {
	routes.Route()
}