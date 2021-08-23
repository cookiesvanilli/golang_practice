package main

import (
	"io"
	"log"
	"net"
	"os"
)
//Клиентская сторона
func main() {
	conn, err := net.Dial("tcp", "localhost:8000") // Для отправки запросов к ресурсам в сети применяется функция net.Dial().
	//Эта функция принимает 2 параметра: network, address
	if err != nil { //выдает ошибку
		log.Fatal(err) //прерывание (аварийное) и логирование ошибки
	}
	defer conn.Close() //закрывает обе половины сетевого подключения
	io.Copy(os.Stdout, conn) // NOTE: ignoring errors
}
