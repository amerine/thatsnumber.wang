package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

var page = `
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">
    <title>thats numberwang</title>
    <style type="text/css">
      body {
        overflow: hidden;
	background: black;
      }
      .container {
        position: relative;
        width: 100%;
        height: 0;
        padding-bottom: 56.25%;
      }
      .video {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <iframe class="video" src="//www.youtube.com/embed/qjOZtWZ56lc?rel=0&autoplay=1" frameborder="0" allowfullscreen></iframe>
    </div>
  </body>
</html>
`

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, page)
}
