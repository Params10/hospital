package reception

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ReceptionHandler recieves alerts
func ReceptionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		var receivedObject received
		json.Unmarshal(body, &receivedObject)
		resp := alertReceiver(receivedObject)
		fmt.Fprintf(w, resp)
	default:
		fmt.Fprintf(w, "Only post methods supported.")
	}
}