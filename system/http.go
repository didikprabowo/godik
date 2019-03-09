package system

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

	server := new(http.Server)
	server.Addr = address
	server.ListenAndServe()
}
