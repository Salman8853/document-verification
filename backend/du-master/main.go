package main

import (
	"du-master/config"
	"du-master/router"
	"log"
	"net/http"
)

func main() {

	r := router.Router()

	// fmt.Println(os.Environ())
	log.Printf("Server started at port %q ", config.APP_CONFIG.PORT)
	if err := http.ListenAndServe(":"+config.APP_CONFIG.PORT, r); err != nil {
		log.Fatal(err)
	}
}
