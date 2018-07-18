package main

import (
    "fmt"
    "bytes"
    "os/exec"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    //"os/exec"
)

type osWindows struct {
    command string
    w http.ResponseWriter
}

func (osW *osWindows) execCommand(command string) {
    
    if bytes.EqualFold([]byte(command), []byte("vol_plus")) {
        fmt.Fprintf(osW.w,command)
        out, err := exec.Command("nircmd.exe changesysvolume 2000").Output()
        if err != nil {
		    log.Fatal(err)
    	}
    	fmt.Printf("The date is %s\n", out)
        //cmd.Run()
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w, r.URL.Path)
}

func main() {
    doStartWebServer()
    doStartSerialPortScaner();
    
    //var input string
    //fmt.Scanln(&input)
}

func commandsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Command List: vol_plus, vol_minus, shutdown, hibernation, mute")
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    name := params["name"]
    fmt.Fprintf(w, name)
    
    osW := new(osWindows)
    osW.w = w
    osW.execCommand(name);
}

func doStartWebServer() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/command/{name}/", commandHandler)
    rtr.HandleFunc("/commands/", commandsHandler)

    http.Handle("/", rtr)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func doStartSerialPortScaner() {
    fmt.Println("Hello World")
}