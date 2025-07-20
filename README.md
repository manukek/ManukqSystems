```go
package main

import (
    "fmt"
)

func main() {
    fmt.Printf(`Simple Website on Go — Manukq Systems

Использование:
  1. git clone https://github.com/manukek/ManukqSystems.git
  2. cd ManukqSystems
  3. go build -o site
  4. ./site
  5. Открой http://localhost:8080 , либо сбилдите бинарник, разверните на хостинге, настройте NGINX конфиги, и будет вам сайтик с красивым доменом (за вопросами ТГ:@manukqq)

  Для онли теста можно запустить одной командой:
  |-                                          -|
  go run github.com/manukek/ManukqSystems@latest
  |-                                          -|

Структура:
  /static/       — статические файлы (css, js, изображения)
  /templates/    — HTML-шаблоны с функцией url_for для статических ссылок

Сайт:
  - Подключите стили в шаблонах:
      <link href="{{ url_for "static" }}css/style.css" rel="stylesheet">
  - Изменяйте index.html и static/css/style.css на ваш вкус
  - Добавляйте роуты в main.go по аналогии (если нужно будет.)

Авторство:
  Меняйте содержимое под себя
  Буду рад, если оставите честь: "by Manukq with ❤️"

Удачи!
`)
}
```