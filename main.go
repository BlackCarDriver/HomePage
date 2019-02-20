package main
 
import(
	"net/http"
	"fmt"
	"time"
)

func main(){
	fmt.Println("It Program is use for provice image service,")
	fmt.Println("default images path is: /home/blackcardriver/Documents/date/images/")
	mux := http.NewServeMux()
	mux.HandleFunc("/Testnet",Testnet)
	mux.HandleFunc("/images",GetImages)
	mux.HandleFunc("/gethomepageartical",GetHomePageArtical)
	mux.HandleFunc("/gethomepagehotnews",GetHomePageHotnews)
	server := &http.Server{
		Addr: "0.0.0.0:4400",
		Handler: mux,
		ReadTimeout: time.Duration(10 * int64(time.Second)),
	}
	server.ListenAndServe()
}
