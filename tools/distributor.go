package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	makeTemplateFolder()
}

func makeTemplateFolder() {
	flag.Parse()
	flg := flag.Args()

	//引数は2つ。第一引数は対象フォルダ。第二引数は対象フォルダ配下に作成するフォルダ数。
	if len(flg) != 2 {
		fmt.Println("fewer args than expected")
	}

	//カレントディレクトリを取得
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	//引数を取得
	parentFolder := flg[0]
	num, err := strconv.Atoi(flg[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	//作成予定のディレクトリパスを宣言
	//https://mattn.kaoriya.net/software/lang/go/20171024130616.htm
	path := filepath.Join(pwd, parentFolder)

	//ディレクトリ（サブディレクトリを含む場合はMkdirAll）を作成
	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
	}

	//forループの中でdeferする場合、キューはLIFOのためにメモリを大量使用し、爆発する
	//この程度のコードの場合、問題にならないので、関数スコープ（func()）を使わず、forループの中の最後でdeferしても良い
	//https://mattn.kaoriya.net/software/lang/go/20151212021608.htm
	//https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
	for i := 1; i <= num; i++ {
		func() {
			childFolder := "task" + strconv.Itoa(i)

			//ディレクトリを作成
			if err := os.Mkdir(filepath.Join(path, childFolder), 0755); err != nil {
				panic(err)
			}

			//コピー元のファイルを宣言
			src, err := os.Open(filepath.Join(pwd, "/tools/template/main.go"))
			if err != nil {
				panic(err)
			}
			defer src.Close()

			//コピー先のファイルを宣言
			dst, err := os.Create(filepath.Join(path, childFolder, "main.go"))
			if err != nil {
				panic(err)
			}
			defer dst.Close()

			//コピーを実施
			_, err = io.Copy(dst, src)
			if err != nil {
				panic(err)
			}
		}()

	}
}
