package main

import (
	"flag"
	"log"

	"github.com/kapustkin/GoHomeWork/Lesson_11/src/internal"
)

func main() {
	argFrom := flag.String("from", "", "Путь к исходному файлу")
	argTo := flag.String("to", "", "Путь к новому файлу")
	argBuff := flag.Int("buffer", 1024, "Размер буфера, в кБ")
	flag.Parse()
	err := internal.CopyFileToFile(*argFrom, *argTo, *argBuff)
	if err != nil {
		log.Fatalf("Ошибка при выполнении операции! %s", err)
	}
	log.Printf("Копирование выполнено успешно\n")
}
