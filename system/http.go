package system

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Port struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var port Port

//Http Server
func HttpServer() {
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&port)

	var address = port.Host + ":" + port.Port
	fmt.Printf("server started at %s\n", address)
	routes := RegisterRoutes()
	server := new(http.Server)
	server.Addr = address
	server.Handler = routes
	server.WriteTimeout = time.Second * 15
	server.ReadTimeout = time.Second * 15
	server.IdleTimeout = time.Second * 60
	// server run
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
