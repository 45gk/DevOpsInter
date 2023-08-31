package main

import (
    "fmt"
    "net/http"
    "github.com/gomodule/redigo/redis"
)

func main() {
    // Подключиться к Redis
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Создать HTTP-сервер
    http.HandleFunc("/set_key", func(w http.ResponseWriter, r *http.Request) {
        // Обработать запрос
        var key, val string
        if err := json.NewDecoder(r.Body).Decode(&key, &val); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        // Сохранить пару ключ-значение в Redis
        _, err := c.Do("SET", key, val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        // Отправить ответ
        w.WriteHeader(http.StatusOK)
    })

    http.HandleFunc("/get_key", func(w http.ResponseWriter, r *http.Request) {
        // Обработать запрос
        key := r.URL.Query().Get("key")

        // Получить значение по ключу из Redis
        val, err := c.Do("GET", key)
        if err != nil {
            if err == redis.ErrNil {
                w.WriteHeader(http.StatusNotFound)
                return
            }

            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        // Отправить ответ
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(val))
    })

    http.HandleFunc("/del_key", func(w http.ResponseWriter, r *http.Request) {
        // Обработать запрос
        key := r.URL.Query().Get("key")

        // Удалить пару ключ-значение из Redis
        _, err := c.Do("DEL", key)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        // Отправить ответ
        w.WriteHeader(http.StatusOK)
    })

    http.ListenAndServe(":8089", nil)
}
