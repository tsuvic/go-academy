package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	makeTemplateFolder()
}

func makeTemplateFolder() {
	flag.Parse()
	if flg := flag.Args(); len(flg) == 2 { //引数は2つのみ。第一引数は対象フォルダ。第二引数は対象フォルダ配下に作成するフォルダ数。

		pwd, err := os.Getwd() //①カレントディレクトリ取得。必ずrootで本プログラムを実行する。
		if err != nil {
			fmt.Println(err)
			return
		}

		path := flg[0]       //②第一引数を取得。
		folderName := "task" //③フォルダ名を決定。

		num, err := strconv.Atoi(flg[1]) //④第二引数の値を取得。
		if err != nil {
			fmt.Println(err)
			return
		}

		fullpath := pwd + "/" + path + "/" + folderName //①＋②＋③＋④でフォルダ作成

		for i := 1; i <= num; i++ {
			if err := os.Mkdir(fullpath+strconv.Itoa(i), 0755); err != nil {
				panic(err)
			}

			fmt.Println(pwd + "/tools/template/main.go")
			src, err := os.Open(pwd + "/tools/template/main.go") //コピー元
			if err != nil {
				panic(err)
			}

			fmt.Println(fullpath + strconv.Itoa(i) + "/" + "main.go")
			dst, err := os.Create(fullpath + strconv.Itoa(i) + "/" + "main.go") //コピー先
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(dst, src)
			if err != nil {
				panic(err)
			}

			//https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
			//https://mattn.kaoriya.net/software/lang/go/20151212021608.htm
			//for文の中に関数スコープ（無名関数）でopen, create, defer closeをするのが良い。
			//deferは見通しを良くするためだから、この程度のコードならこれで良い？
			src.Close()
			dst.Close()
		}
	}

}
