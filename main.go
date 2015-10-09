package main

import (
    "log"
    "net/http"
)



func droneServer(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Built by Drone!"))
}

func main() {
    http.HandleFunc("/", droneServer)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
