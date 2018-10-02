////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var PlayerBot = Bot{}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	/* make sure there is a body to the request */
	if r.Body == nil {
		return
	}

	/* read all the request into a buffer */
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	/* decode the url-encoded buffer */
	data, err := url.QueryUnescape(buf.String())
	if err != nil {
		fmt.Println("Couldn't decode URL encoded data.")
	}

	/* remove `data=` from the start of the buffer */
	data = data[5:]

	/* decode the JSON from the buffer */
	var jsonGame JSONGameInfo
	err = json.Unmarshal([]byte(data), &jsonGame)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* get the next action*/
	var game = jsonGame.GameInfo()
	PlayerBot.ExecuteAction(&game.Player, &game.Map)
}

func main() {
	/* add the routes of the API */
	http.HandleFunc("/", handleRequest)

	/* start the HTTP server */
	if err := http.ListenAndServe(":5555", nil); err != nil {
		panic(err)
	}
}
