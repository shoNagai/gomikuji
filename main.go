// gotry project main.go
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/*
今週の運勢
URLGetパラメータ
name:占い対象の名前
*/
func handler(w http.ResponseWriter, r *http.Request) {
	var name string = r.FormValue("name")
	if name == "" {
		name = "名無しさん"
	} else {
		name = name + "さん"
	}

	name = fmt.Sprintf("〜%sの今週の運勢〜\n", name)

	omi := showOmikuji()

	buf := make([]byte, 0)
	buf = append(buf, name...)
	buf = append(buf, omi...)

	fmt.Fprint(w, string(buf))
}

/*
	一週間分のgomikuji結果を返却
	戻り値：string
*/
func showOmikuji() string {

	t := time.Now().UnixNano()
	rand.Seed(t)

	yobi := map[int]string{0: "月", 1: "火", 2: "水", 3: "木", 4: "金", 5: "土", 6: "日"}

	buf := make([]byte, 0)

	for i := 0; i < 7; i++ {

		msg := gomikuji()

		str := fmt.Sprintf("%s曜日の運勢：%s\n", yobi[i], msg)
		buf = append(buf, str...)
	}

	return string(buf)

}

/*
	ランダムにおみくじ結果を返却
	戻り値：stirng
    6：大吉
	5,4：中吉
	3,2：吉
	1：凶
*/
func gomikuji() string {

	kuji := rand.Intn(6) + 1
	var msg string = ""

	if kuji == 1 {
		msg = "凶"
	} else if kuji == 2 || kuji == 3 {
		msg = "吉"
	} else if kuji == 4 || kuji == 5 {
		msg = "中吉"
	} else if kuji == 6 {
		msg = "大吉"
	}

	return msg
}
