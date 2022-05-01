# Дневник

## Описание
Веб-приложение где составляешь и следишь за успеваемостью по предметам

### Требовании:
* [Git](https://docs.github.com/en/desktop/installing-and-configuring-github-desktop/installing-and-authenticating-to-github-desktop/installing-github-desktop)
* [Golang](https://go.dev/doc/install)
* [MySQL](https://dev.mysql.com/downloads/mysql/)
* [phpMyAdmin](https://www.phpmyadmin.net/) - для удобства пользовался [Open Server Panel](https://ospanel.io/) 

### Go пакеты:
* `go get github.com/go-sql-driver/mysql`
* `go get -u github.com/gorilla/mux`

> если выходит ошибка:  `go.mod file not found in current directory or any parent directory.` то попробуйте ввести `go env -w GO111MODULE=off`

### Запуск:
* `git clone https://github.com/hopeurfnis12/PD.git`
* Поменять в **main.go** `user` и `pswr` для авторизации в **phpMyAdmin**
* Запустить сервер `go run main.go` 	
* В браузере перейти по адресу `http://localhost:7272/`

## Roadmap:
### MVP
* Шаблон страницы с предметами ✅
* Логика показа предметов ✅
* Шаблон страницы с ту-ду листом ✅
* Логика показа ту-ду листа ✅

### Добавление
* Шаблон формы добавления ✅✅
* Логика создания ✅✅

### Поиск ❎
* Шаблон страницы поиска
* Логика поиска

### Редактирование ❎
* Шаблон страницы редактирования
* Логика редактирования

### Удаление/Зачеркивания
* Логика удаления ✅
* Логика зачеркивания ✅
