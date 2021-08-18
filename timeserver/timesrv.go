package main

import (
	"io"
	"log"
	"net"
	"time"
)
//Серверная часть
func main() {
	println("Time server started...") //пишем строку, для того чтобы видеть что сервер запустился
	listener, err := net.Listen("tcp", "localhost:8000") //для прослушивания и приема входящих запросов
	// в пакете net определена функция net.Listen, которая принимает параметры network, address - локальный сервер, по которому будет запускаться сервер.
	if err != nil { //в случае неудачи выдает ошибку
		log.Fatal(err)
	}
	for { //запускается бесконечный цикл
		conn, err := listener.Accept() // метод Accept принимает входящее подключение и ожидает соединения к серверу
		// если ошибки нет, то передается информация, иначе происходит аварийный разрыв
		if err != nil {
			log.Print(err)
			continue //возвращаемся в начало цикла
		}
		go handleConn(conn) //однопоточность, если добавим go перед функцией, то многопоточность
	}
}

func handleConn(c net.Conn) {
	defer c.Close() //закрытие (откладывается до исполнения самой функции)
	for { //запускается бесконечный цикл
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r")) // передается строка,
		// где записывается текущее время в определенном формате
		if err != nil {
			return // если ошибка, то выходим из цикла и срабатывает defer c.Close()
		}
		time.Sleep(1 * time.Second) //время обновляется каждую секунду
	}
}
