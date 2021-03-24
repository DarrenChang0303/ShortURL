package controller

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	mongoDB "shortURL/database"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var tpl *template.Template
var hostName string
var localIP string = "http://localhost:8000"

type NotAvailable struct {
	Title string
}

//There has no input to JSON if the value is empty
type RequestData struct {
	OriginalURL string `json:"originalURL, omitempty"`
	Alias       string `json:"alias,omitempty`
}

type ResponseData struct {
	OriginalURL string `json:"originalURL, omitempty"`
	CreateURL   string `json:"createURL, omitempty"`
	ID          string `json:"id, omitempty"`
	IsAlias     bool   `json:"isRequest, omitempty"`
}

//Init function initialize the html file
func Init() {
	tpl = template.Must(template.ParseGlob("view/*.html"))
}

//Showing the home page server

func HomePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Show home page")
	Init()
	err := tpl.ExecuteTemplate(writer, "homPage.html", nil)
	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("error load homPage.html"))
	}

}

/*
判斷資料庫是否有原網址
有->	查詢已有的原網址並提供對立短網址
沒有->	儲存原網址並創建新短網址後儲存
*/

/*
確認隨機數:
1.用時間導出
2.利用原網址內容導出
3.返回程式碼在DB內再確認一次
*/

func CreateURL(w http.ResponseWriter, r *http.Request) {
	//
	var request RequestData
	// str := r.PostFormValue("originalURL")
	// request.OriginalURL = PrefixSlash("originalURL")
	fmt.Println("Input URL:", request.OriginalURL)
	fmt.Println("Alias Id:", request.Alias)

	var res NotAvailable
	// forbiddenInput := "http://localhost:8000"
	if strings.Contains(request.OriginalURL, localIP) {
		res.Title = "URL domain banned"
		tpl.ExecuteTemplate(w, "Not available.html", res)
		return
	}
	if request.Alias == "" {
		fmt.Print("Create shortURL without alias")
		RequestData := CreateWithoutAlias(request)
		tpl.ExecuteTemplate(w, "create.html", RequestData) //TODO
		return
	} else {
		fmt.Print("Create with alias")
		RequestData := CreateWithAlias(request)
		tpl.ExecuteTemplate(w, "Available.html", RequestData) //TODO
	}
	// fmt.Println("Send CreateURL respond successful")
}

func CreateWithoutAlias(request RequestData) ResponseData {
	var response ResponseData
	collection := mongoDB.MongoClient.Database("url_database").Collection("url_table")
	collection.FindOne(context.Background(), bson.M{"OriginalURL": request.OriginalURL, "IsAlias": false}).Decode(&response)
	if (response == ResponseData{}) {
		response.ID = uniqueID()
		response.OriginalURL = request.OriginalURL
		response.CreateURL = localIP + response.ID
		collection.InsertOne(context.TODO(), bson.M{
			"ID":          response.ID,
			"OriginalURL": response.OriginalURL,
			"CreateURL":   response.CreateURL,
			"IsAlias":     false,
		})
		fmt.Print("Insert successful")
	}
	fmt.Print("Create shortURL without alias")
	return response

}

func CreateWithAlias(request RequestData) ResponseData {
	var response ResponseData
	collection := mongoDB.MongoClient.Database("url_database").Collection("url_table")
	collection.FindOne(context.Background(), bson.M{"ID": request.Alias}).Decode(&response)
	if (response == ResponseData{}) {
		response.ID = request.Alias
		response.CreateURL = localIP + response.ID
		response.OriginalURL = request.OriginalURL
		collection.InsertOne(context.TODO(), bson.M{
			"ID":          response.ID,
			"OriginalURL": response.OriginalURL,
			"CreateURL":   response.CreateURL,
			"IsAlias":     true,
		})
	}
	fmt.Print("Create shortURL with alias")

	return response
}

func uniqueID() string {
	ID := strconv.Itoa(int(time.Now().Unix()))

	return ID
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Redirect to original URL")
	params := mux.Vars(r)
	collection := mongoDB.MongoClient.Database("url_database").Collection("url_table")
	var response ResponseData
	collection.FindOne(context.Background(), bson.M{"ID": params["id"]}).Decode(&response)
	if (response != ResponseData{}) {
		http.Redirect(w, r, response.OriginalURL, 303)
		fmt.Println("Redirect successful")
	}

}
