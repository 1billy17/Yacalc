# Yacalc

**Yacalc** — это HTTP-сервис, предназначенный для вычисления математических выражений, которые передаются в формате JSON. Этот проект позволяет отправлять выражения с помощью POST-запросов и получать либо результаты вычислений, либо сообщения об ошибках.

## Возможности

- **Корректный ввод**: Выражение вычислено успешно.
- **Ошибка 422**: Входные данные не соответствуют требованиям приложения.
- **Ошибка 500**: Случай какой-либо иной ошибки («Что-то пошло не так»).



## Примеры использования

### 1. **Успешное выполнение**

Вычисление выражения `2+2*2`:

```
curl -i -X POST http://localhost:8080/api/v1/calculate \
--header "Content-Type: application/json" \
--data '{"expression": "2+2*2"}'
```

**Ответ**:
```
HTTP/1.1 200 OK
Date: Sun, 22 Dec 2024 03:10:57 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

{"result":6}%
```

### 2. **Ошибка 422**

Ввод некорректного выражения (`fsvs`):

```
curl -i -X POST http://localhost:8080/api/v1/calculate \
--header "Content-Type: application/json" \
--data '{"expression": "fsvs"}'
```

**Ответ**:
```
HTTP/1.1 422 Unprocessable Entity
Date: Sun, 22 Dec 2024 03:12:40 GMT
Content-Length: 35
Content-Type: text/plain; charset=utf-8

{"error":"Expression is not valid"}%
```

### 3. **Ошибка 500**

Передача пустого выражения:

```
curl -i -X POST http://localhost:8080/api/v1/calculate \
--header "Content-Type: application/json" \
--data '{"expression": ""}'
```

**Ответ**:
```
HTTP/1.1 500 Internal Server Error
Date: Sun, 22 Dec 2024 03:16:09 GMT
Content-Length: 33
Content-Type: text/plain; charset=utf-8

{"error":"Internal server error"}% 
```

## Инструкция по запуску проекта

1. **Клонируете репозиторий**:
   ```
   git clone https://github.com/1billy17/Yacalc.git
   cd Yacalc
   ```

2. **Запускаете проект**:
   ```
   go run main.go
   ```

3. **Сервер запустится на порту 8080**.

Теперь можно отправлять запросы на эндпоинт:  
http://localhost:8080/api/v1/calculate


## Инструкция по запуску тестов

1. **Введите команду**:
   ```
    go test -v ./pkg/calculation/tests
   ```

2. **Получите результат**:
   ```
   === RUN   TestYacalc
    === RUN   TestYacalc/Valid_Expression
    === RUN   TestYacalc/Invalid_Expression
    === RUN   TestYacalc/Internal_server_error
    --- PASS: TestYacalc (0.00s)
    --- PASS: TestYacalc/Valid_Expression (0.00s)
    --- PASS: TestYacalc/Invalid_Expression (0.00s)
    --- PASS: TestYacalc/Internal_server_error (0.00s)
    PASS
    ok      Yacalc/pkg/calculation/tests    (cached)
   ```