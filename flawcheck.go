package main

import (
  "fmt"
  "os"
  "github.com/parnurzeal/gorequest"
  "gopkg.in/alecthomas/kingpin.v2"
  "time"
)

var (
  verbose   = kingpin.Flag("verbose", "Run in verbose mode.").Bool()
  host = kingpin.Arg("host", "FlawCheck Host URL").Default("https://api.flawcheck.com/").String()
  api_version = kingpin.Arg("api_version", "FlawCheck API Version").Default("v1").String()

  //timeout = kingpin.Flag("timeout", "Timeout waiting for upload").OverrideDefaultFromEnvar("PING_TIMEOUT").Required().Short('t').Duration()
  api_key     = kingpin.Arg("api_key", "FlawCheck REST API Key").String()
  container   = kingpin.Arg("container", "Path to container to upload").String()
)

func main(){
  kingpin.Version("1.0.0")
  kingpin.Parse()
  if(*api_key == ""){
    fmt.Println("Please provide your API Key!, see --help.")
    os.Exit(1)
  }
  if(*container == ""){
    fmt.Println("Please provide the path to your image file!, see --help.")
    os.Exit(1)
  }

  //fmt.Printf("API KEY %s", *api_key)
  os.Exit(0)
}

func upload_container(){
  request := gorequest.New().Timeout(2 * time.Minute)
  resp, body, errs:= request.Get(*host).End()

}

func is_ready(){
}

func get_report(){
}

func parse_json(){
}

func get_score(){

}
