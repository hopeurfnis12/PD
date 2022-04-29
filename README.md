# Дневник

## Описание
Веб-приложение где составляешь и следишь за успеваемостью по предметам

### Требовании:
* [Git](https://docs.github.com/en/desktop/installing-and-configuring-github-desktop/installing-and-authenticating-to-github-desktop/installing-github-desktop)
* [Golang](https://go.dev/doc/install)
* [MySQL](https://dev.mysql.com/downloads/mysql/)

### Запуск:
* `git clone https://github.com/hopeurfnis12/PD.git`
* `go get github.com/go-sql-driver/mysql` - скачать **MySQL Driver**
  > если выходит ошибка:  `go.mod file not found in current directory or any parent directory.` то попробуйте ввести `go env -w GO111MODULE=off`
* `go get -u github.com/gorilla/mux` - скачать
* Настроить phpMyAdmin
* Запустить сервер `go run main.go` 	
* В браузере перейти по адресу `http://localhost:7272/`

## Roadmap:
### MVP
* Шаблон страницы с успеваемостью ✅
* Логика показа успеваемости ✅
* Шаблон страницы с ту-ду листом ✅
* Логика показа ту-ду листа ✅

### Авторизация ❎
* Шаблон страницы регистрации
* Логика регистрации
* Шаблон страницы входа
* Логика авторизации

### Добавление
* Шаблон страницы добавления ✅❎
* Логика создания ✅❎

### Поиск ❎
* Шаблон страницы поиска
* Логика поиска

### Редактирование ❎
* Шаблон страницы редактирования
* Логика редактирования

### Удаление/Зачеркивания ❎
* Шаблон страницы удаления
* Логика удаления
