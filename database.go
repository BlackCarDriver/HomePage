package main
import(
	"io/ioutil"
	"fmt"
	_"github.com/lib/pq"
	"database/sql"
	"log"
	"strings"
	"strconv"
)
var  img_path string = "/home/blackcardriver/Documents/date/images/"
var Db *sql.DB	

func init(){
	var err error
	Db,err = sql.Open("postgres","user=postgres dbname=homepage password=include123 ")
	if err != nil {
		log.Fatal("sql open err :",err)
		return
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal("sql ping err :",err)
	}else{
		log.Println("database connect success...")
	}
}

//Get the lastest fith news data
func GetHomePageHotnewDate()(date [5]HomePageNews){
	command := "select imgurl, linkurl, title from homepagenew order by id desc limit 5"
	rows,err := Db.Query(command)
	defer rows.Close()
	if err!= nil{
		log.Println("error at gethomepagenewdate: ",err)
		return
	}
	index := 0
	for rows.Next() {
		rows.Scan(&date[index].Img_url, &date[index].Link_url, &date[index].Title)
		index++ 
	}
	return 
}

//Get HomePageArtical Date fome database index from 'index'
func GetHPADate(index int64)(date []ArticlaBox){
	command := "select imgurl, linkurl, brif, date from homepageartical order by id asc offset $1 limit 10"
	rows,err := Db.Query(command,index)
	var temp ArticlaBox
	defer rows.Close()
	if err!=nil{
		fmt.Println("error at GetHomePageArtical() , err:",err)
		return
	}
	for i := 0;rows.Next();i++ {
		rows.Scan(&temp.Img_url, &temp.Link_url, &temp.Brif, &temp.Date)
		date = append(date,temp)
	}
	return
}



//create some mark date for test
func makedata(){
	templace := `insert into homepageartical(imgurl,linkurl,date)values(
		'http://129.204.193.192:4400/get?tag=&name=$.jpg',
		'http://129.204.193.192:4400/get?tag=&name=$.jpg',
		CURRENT_DATE
	);`
	for i:=12;i<60;i++ {
		command := strings.Replace(templace, "$",strconv.Itoa(i),-1)
		_ , err := Db.Query(command)
    	if err != nil {
        	fmt.Println("worng!err:",err)
    	}
	}
	return;
}












//return images's byte that user needed
func Images(tag string, name string)[]byte{
	filepath := img_path + tag + name
	temp,err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Images() err: ", err)
	}
	return temp
}
