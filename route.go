package main
import(
	// "log"
	"net/http"
	// "net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
func GetImages(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	vars := r.URL.Query()
	tag := vars["tag"]
	name := vars["name"]
	if( len(tag)!=1 && len(name)!=1 ){
		fmt.Println("GetImages url worng, vars:",vars)
		return
	}
	w.Write( Images(tag[0],name[0] ) )
	return
}

func GetHomePageArtical(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Access-Control-Allow-Headers","Content-Type");
	body, _ := ioutil.ReadAll(r.Body)
	if(len(body) == 0){
		return
	}
	var postbody map[string]int64
	json.Unmarshal(body, &postbody)
	index := postbody["index"]	//get artical index from which
	fmt.Println(index)
	var date = GetHPADate(index)
	date_json,_ := json.Marshal(date)
	w.Write(date_json)
}	

func GetHomePageHotnews(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Access-Control-Allow-Headers","Content-Type")
	date := GetHomePageHotnewDate()
	date_json,_ := json.Marshal(date)
	w.Write(date_json)
}

//use for test net connection
func Testnet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("net connect scuess!..."))
}