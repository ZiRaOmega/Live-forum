package main

func main() {
	db := GetDB()
	sqlMaker(db)
	defer CloseDB()

	StartServerHandler()
}
