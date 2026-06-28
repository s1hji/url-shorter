package main

import (
	"flag"
	"log"
	"net/http"
	"url-shorter/api/handlers"
	"url-shorter/postgres"
	"url-shorter/storage"
	"url-shorter/storage/memory"
)

func main() {
	storageType := flag.String("storage", "memory", "тип хранилища: moemory || postgres")
	addr := flag.String("addr", ":8080", "адрес сервера")
	connect := flag.String("connect", "", "строка подключения к postgres")
	flag.Parse()

	var storage storage.Storage

	switch *storageType {
	case "memory":
		storage = memory.NewMemoryStorage()
		log.Println("используется memory хранилище")
	case "postgres":
		if *connect == "" {
			log.Fatal("для postgres нужно передать флаг -connect")
		}
		postgres, err := postgres.NewPostgresStorage(*connect)
		if err != nil {
			log.Fatalf("не удалось подключиться к postgres %v", err)
		}
		defer postgres.Close()
		storage = postgres
		log.Println("сейчас используется хранилище postgres")

	default:
		log.Fatalf("непонятное хранилище %s", *storageType)
	}

	shorterHandler := handlers.ShortenHandler{
		Store: storage,
	}
	redirectHandler := handlers.RedirectHandler{
		Storage: storage,
	}

	http.HandleFunc("/shorter", shorterHandler.ServeHttp)
	http.HandleFunc("/", redirectHandler.RedirectServeHttp)
	http.ListenAndServe(*addr, nil)

}
