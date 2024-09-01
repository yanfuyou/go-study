package main

import (
    "os"
    "io"
    "log"
    "net/http"
	"html/template"
)


const (
    UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter,r *http.Request){
    if r.Method == "GET" {
        // io.WriteString(w,"<html><header><header/><body><form method=\"POST\" action=\"/upload\"" +
	    // " enctype=\"multipart/form-data\">"+
	    // "Choose an image to upload:<input name=\"image\" type=\"file\" /> "+
	    // "<input type=\"submit\" value=\"Upload\"/>"+
	    // "</form></body></html>")
		t,err := template.ParseFiles("./html/upload.html")
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		t.Execute(w,nil)
	    return
    }
    if r.Method == "POST" {
	f,h,err := r.FormFile("image")
	if err!= nil {
	    http.Error(w,err.Error(),
	    http.StatusInternalServerError)
	    return
	}
	filename := h.Filename
	defer f.Close()
	t,err := os.Create(UPLOAD_DIR + "/" + filename)
	if err != nil {
	    http.Error(w,err.Error(),http.StatusInternalServerError)
	    return
	}
	defer t.Close()
	if _,err := io.Copy(t,f);err != nil {
 	    http.Error(w,err.Error(),http.StatusInternalServerError)
	    return
	}
	http.Redirect(w,r,"/view?id="+filename, http.StatusFound)
    }
}


func viewHandler(w http.ResponseWriter,r *http.Request){
	fileId := r.FormValue("id")
	fpath := UPLOAD_DIR +"/" + fileId
	w.Header().Set("Content-Type","image")
	http.ServeFile(w,r,fpath)
}

func main(){
	http.HandleFunc("/view",viewHandler)
    http.HandleFunc("/upload",uploadHandler)
    err := http.ListenAndServe(":8080",nil)
    if err != nil{
	log.Fatal("ListenAndServer: ",err.Error())
    }
}
