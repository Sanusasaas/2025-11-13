## Пример запроса на проверку ссылок: 
```bash
$ curl -X POST http://localhost:8080/check -H "Content-Type: application/json" -d '{"links":["https://google.com", "https://yandex.ru"]}'
```
## Пример запроса на получение пдф файла:
```bash
$ curl -X POST http://localhost:8080/report -H "Content-Type: application/json" -d '{"links_num": [1]}' --output report.pdf
```
