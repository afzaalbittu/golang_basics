package main
import (
	"fmt"
	"io/ioutil"

	"net/http"
	//"go/types"
	//"io"
	//"flag"
	"net/url"
	"strings"
	"os"
)

func downloadmp3(urlref string){
	resp, error:= http.Get(urlref)
	if error!= nil{
		fmt.Println("error")

	}
       defer resp.Body.Close()

       dataInBytes,error:= ioutil.ReadAll(resp.Body)
       if error!= nil {

       	fmt.Println("error")


	   }
	   ParsedUrlRef,_ := url.Parse(urlref)

	   filename:= ParsedUrlRef.Path
	   filename = strings.Replace(filename,"/","_",-1)
	   filename = strings.TrimLeft(filename,"_")
	   filename = "/Users/Naseer/Downloads/"+filename +".html"
       filehandle,error:= os.OpenFile(filename,os.O_CREATE|os.O_WRONLY, 0600)
       if error!= nil{
       	fmt.Println("error")
       	return
	   }
	   defer filehandle.Close()
	   filehandle.WriteAt(dataInBytes,19)



}
func main(){
	downloadmp3("https://golang.org/pkg/go/parser/#example_ParseFile")
}

