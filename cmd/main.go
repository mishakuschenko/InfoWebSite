package main

import (
    "log"
    "net/http"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "home page" как тело ответа.

func home(page http.ResponseWriter, request *http.Request) {
    page.Write([]byte("home page"))
}

func main() {
      
    // Используется функция http.NewServeMux() для инициализации нового рутера, затем
    // функцию "home" регистрируется как обработчик для URL-шаблона "/", то есть главной страницы.

    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    // http.ListenAndServe(":4000", mux) запускает сервер
    // Мы передаем два параметра: TCP-адрес сети для прослушивания (в данном случае это "localhost:4000")
    // и созданный рутер. Если вызов http.ListenAndServe() возвращает ошибку
    // мы используем функцию log.Fatal() для логирования ошибок. 

    log.Println("Запуск веб-сервера на http://127.0.0.1:4000") 

  // ^ перед запуском указываем то, что все норм

    err := http.ListenAndServe(":4000", mux)
    if err != nil {
    log.Fatal(err)
    //обрабатываем ошибки запуска (вся if конструкция)
}        
}
