package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jtarte/myservice/utils"

	log "github.com/sirupsen/logrus"
)

// Caller struct
// describe the caller that is doing call to url
type Caller struct {
	servicename string
	version     string
	target      string
}

// NewCaller inits a new Caller instance
func NewCaller() *Caller {

	caller := new(Caller)
	caller.servicename = os.Getenv("NAME")
	if caller.servicename == "" {
		caller.servicename = "DummyService"
	}
	caller.version = os.Getenv("VERSION")
	if caller.version == "" {
		caller.version = "v0"
	}
	caller.target = os.Getenv("TARGET")

	return caller
}

// index handles the processing of an URL
// w the HTTP writer used to send the response
// r the HTTP request
func (c *Caller) invoke(w http.ResponseWriter, r *http.Request) {
	log.Debug("Entering invoke function")
	stack := "[" + c.servicename + " " + c.version
	if c.target == "" {
		log.Info("no externall call")
	} else {
		targetTab := strings.Split(c.target, ",")
		for i, s := range targetTab {
			targetURL := "http://" + s + "/invoke"
			log.Debug("trying to call " + targetURL)
			response, err := http.Get(targetURL)
			returnedMsg := ""
			if err != nil {
				log.Error(err.Error())
				log.Trace(fmt.Sprint(response))
				returnedMsg = "call_error"
			} else {
				defer response.Body.Close()
				log.Debug("response code: " + strconv.Itoa(response.StatusCode))
				var result map[string]string
				json.NewDecoder(response.Body).Decode(&result)
				log.Trace("response : " + fmt.Sprint(result))
				returnedMsg = result["stack"]
			}
			if i == 0 {
				stack += " calls " + string(returnedMsg)
			} else {
				stack += " then " + string(returnedMsg)
			}
		}
	}
	stack += " ]"
	msg := map[string]string{"stack": stack}
	log.Info("returned stack: " + stack)
	utils.RespondJSON(w, r, 200, msg)
	log.Debug("Exiting invoke function")
}
