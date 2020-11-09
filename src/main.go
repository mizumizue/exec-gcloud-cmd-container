package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Body struct {
	Args string `json:"args"`
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		bytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Errorf("%v", err)
			return
		}

		var body Body
		if err := json.Unmarshal(bytes, &body); err != nil {
			fmt.Println(err)
		}
		res := strings.Split(body.Args, " ")
		out, err := exec.Command("gcloud", res...).Output()
		if err != nil {
			fmt.Fprintf(writer, string(out), err)
			return
		}
		fmt.Fprintf(writer, string(out))
	})
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("server start")
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
