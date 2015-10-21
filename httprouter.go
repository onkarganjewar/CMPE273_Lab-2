package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"

)
type Request struct {
Name string `json:"name"`
Id int `json:"id"`
Address string `json:"address"`
Zip int `json:"zip"`
}

type Response struct {
Greet string `json:"greeting"`
}


func post_h (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  // Stub a request to be populated from the body
    req := Request{}

	  // Populate the request data
    json.NewDecoder(r.Body).Decode(&req)




    // Marshal provided interface into JSON structure
    reqj, _ := json.Marshal(req)
  	w.Header().Set("Content-Type", "application/json")
  	w.WriteHeader(201)

		// Create a new Response struct for storing the unmarshaled JSON struct
		var res Response
		if err := json.Unmarshal(reqj, &res); err != nil {
				fmt.Println(err)
		return
		}
		res.Greet = "Hello, "+req.Name+"!"
		// Marshal provided request into JSON Response
		resj, _ := json.Marshal(res)
		fmt.Fprintf(w,"%s",resj)

}


func hello (rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

  fmt.Fprintf(rw, "Hello, %s!\n",p.ByName("name"))
}




func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Add a handler on /hello
  	r.GET("/hello/:name", hello)
    r.POST("/hello", post_h)

    // Fire up the server
    http.ListenAndServe("localhost:3030", r)
}


/***********************************************OUTPUT*****************


Onkar@onkar-personal MINGW64 /c/Go/src/golang/httpjson
$ go run httprouter.go

Onkar@onkar-personal MINGW64 /c/Go/src/golang/httpjson
$  curl -H "Content-Type: application/json" -X POST -d '{"name":"Foo Bar"}' http://127.0.0.1:3030/hello
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    48  100    30  100    18     30     18  0:00:01 --:--:--  0:00:01 30000{"greeting":"Hello, Foo Bar!"}

**********************************************************************/
