package main
import (
        // "github.com/tidwall/sjson"
        // "github.com/tidwall/gjson"
        // "github.com/glennswest/libignition/ignition"
        "encoding/base64"
        // "github.com/glennswest/libpowershell/pshell"
        // "encoding/json"
        "os"
        "log"
)

func main() {
    log.Println("WNA - Windows Node Adder")
    ignb64 := os.Getenv("workerign")
    if (len(ignb64) == 0){
        log.Printf("wna: No workerign environment variable supplied\n")
        os.Exit(-3)
        }
    ignbytes, _ := base64.StdEncoding.DecodeString(ignb64)
    ign := string(ignbytes)
    log.Printf("ign = %v\n",ign)
}


