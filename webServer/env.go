package main

import(
	"os"
)

func setEnviroment(){
	// Database Name
	os.Setenv("dbName", "NoTrackWebsiteStats")
	// Database Host
	os.Setenv("dbHost", "localhost")
	// Database Port
	os.Setenv("dbPort", "8889")
	// Datbase Username
	os.Setenv("dbUser", "root")
	// Datbase Password
	os.Setenv("dbPass", "root")
}
