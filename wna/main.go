package main
import (
        "github.com/tidwall/gjson"
        "github.com/glennswest/libignition/ignition"
        "ioutil"
        "io"
        "net/http"
	"crypto/tls"
        "os"
        "fmt"
)

func main() {
    fmt.Println("WNA - Windows Node Adder")
    ign := ReadFile("/Program Files/WindowsNodeManager/settings/env/settings/workerign")
    // "ignition":{"config":{"append":[{"source":"
    append := gjson.Get(ign,"ignition.config.append").String()
    nodeignsrc := gjson.Get(append,"0.source").String()
    fmt.Printf("src = %v\n",nodeignsrc)
    downloadfile("/k/compute.ign",nodeignsrc)
    ignition.Parse_ignition_file("/k/compute.ign","")   
    fmt.Printf("All Processing Complete\n")
    os.Exit(0)
}

func ReadFile(thepath string) string {
    b, err := ioutil.ReadFile(thepath) // just pass the file name
    if err != nil {
        //log.Print(err)
        return ""
    }
    str := string(b)
   return str
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


