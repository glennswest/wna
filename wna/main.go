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
	"crypto/tls"
        "os"
        "fmt"
)

func main() {
    fmt.Println("WNA - Windows Node Adder")
    ignb64 := os.Getenv("workerign")
    if (len(ignb64) == 0){
        fmt.Printf("wna: No workerign environment variable supplied\n")
        os.Exit(-3)
        }
    fmt.Printf("Decoding base64 ignition file from workerign environment varialbe\n")
    ignbytes, _ := base64.StdEncoding.DecodeString(ignb64)
    ign := string(ignbytes)
    fmt.Printf("ign = %v\n",ign)
    // "ignition":{"config":{"append":[{"source":"
    append := gjson.Get(ign,"ignition.config.append").String()
    fmt.Printf("append = %v\n",append)
    nodeignsrc := gjson.Get(append,"0.source").String()
    fmt.Printf("src = %v\n",nodeignsrc)
    downloadfile("/k/compute.ign",nodeignsrc)
    fmt.Printf("Starting Processing of 2nd stage ignition file\n")
    ignition.Parse_ignition_file("/k/compute.ign","")   
    fmt.Printf("All Processing Complete\n")
    os.Exit(0)
}

func downloadfile(filepath string, url string) error {

    // Get the data
    fmt.Printf("Downloading %v\n",url)
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    resp, err := http.Get(url)
    if err != nil {
	fmt.Printf("Error getting %s, %v\n",url,err)
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


