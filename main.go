package main

//go:generate gotext -srclang=en update -out=catalog.go -lang=en,el,ru

/*
# Установка
go get -u golang.org/x/text/cmd/gotext
# Каждый раз при изменение текстов в коде *.go запускать "go generate", она внесет измения в файлы языков
# потом откорректировать нужные файлы языков "messages.gotext.json"
# и потом заново запустить "go generate", будет пересоздан файл "catalog.go", в котором вся локаль в коде!
go generate
*/

import (
	"fmt"
	"os"
)

func main() {
	yazik := "en"
	if len(os.Args) == 2 {
		yazik = os.Args[1]
	}

	fmt.Println(lng(yazik, "Hello World"))
}
