package main

import (
	db "hottub/db"
	router "hottub/routes"
)

func main() {
	db.Init()
	router.Init()
}
