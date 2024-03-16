package main

import (
    "github.com/AaronO/gogo-proxy"
    "net/http"
    "os"
)

func main() {
    p, _ := proxy.New(proxy.ProxyOptions{
        Balancer: func(req *http.Request) (string, error) {
            return "http://pan.fewdrive.eu.org", nil
        },
    })
    http.ListenAndServe(":"+os.Getenv("5244"), p)
    
}
