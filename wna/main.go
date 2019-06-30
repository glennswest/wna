package main
import (
        // "github.com/tidwall/sjson"
        "github.com/tidwall/gjson"
        "github.com/glennswest/libignition/ignition"
        "encoding/base64"
        // "github.com/glennswest/libpowershell/pshell"
        // "encoding/json"
        "io"
        "net/http"
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
    // "ignition":{"config":{"append":[{"source":"
    append := gjson.Get(ign,"ignition.config.append").String()
    log.Printf("append = %v\n",append)
    nodeignsrc := gjson.Get(append,"0.source").String()
    log.Printf("src = %v\n",nodeignsrc)
    downloadfile(nodeignsrc,"/k/compute.ignition")
    ignition.Parse_ignition_file("/k/compute.ignition","/k/ignition")   

}

func downloadfile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}


