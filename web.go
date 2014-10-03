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
    <script type="text/javascript" src="http://code.jquery.com/jquery-git.js"></script>
    <style type="text/css">
      body {
        overflow: hidden;
      }
    </style>

    <script type="text/javascript">//<![CDATA[
      $(window).load(function(){
        $(function(){
          $('#video').css({ width: $(window).innerWidth() + 'px', height: $(window).innerHeight() + 'px' });

          $(window).resize(function(){
            $('#video').css({ width: $(window).innerWidth() + 'px', height: $(window).innerHeight() + 'px' });
          });
        });

      });//]]>
    </script>


  </head>
  <body>
    <iframe id="video" src="//www.youtube.com/embed/qjOZtWZ56lc?rel=0&autoplay=1" frameborder="0" allowfullscreen></iframe>
  </body>
</html>
`

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, page)
}
