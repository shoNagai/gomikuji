// gotry project main.go
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := showOmikuji()
	fmt.Fprint(w, msg)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func showOmikuji() string {

	t := time.Now().UnixNano()
	rand.Seed(t)

	yobi := map[int]string{0: "月", 1: "火", 2: "水", 3: "木", 4: "金", 5: "土", 6: "日"}

	buf := make([]byte, 0)

	for i := 0; i < 7; i++ {

		msg := goOmikuji()

		str := fmt.Sprintf("%s曜日の運勢：%s\n", yobi[i], msg)
		buf = append(buf, str...)
	}

	return string(buf)

}

func goOmikuji() string {

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

// ランダムに１〜６までの数字を出す関数
// 6：大吉、
// 5,4：中吉
// 3,2：吉
// 1：凶
//	omikuji := map[int]string{1: "凶", 2: "吉", 3: "吉", 4: "中吉", 5: "中吉", 6: "大吉"}
