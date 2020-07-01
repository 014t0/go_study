package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//http.ResponseWriterはレスポンスに関する構造体→情報を返すときに利用
	//http.Requestはリクエストに関する構造体→リクエストに含まれている情報を取得するときに利用
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)   //webページのルーティングを登録、アクセスした時に呼び出す関数
	http.ListenAndServe(":80", nil) //どこのポートで公開するのかの指定
}
