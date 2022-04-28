# Дневник

## Описание
Веб-приложение где составляешь и следишь за успеваемостью по предметам

### Требовании:
* [Git](https://docs.github.com/en/desktop/installing-and-configuring-github-desktop/installing-and-authenticating-to-github-desktop/installing-github-desktop)
* [Golang](https://go.dev/doc/install)
* [MySQL](https://dev.mysql.com/downloads/mysql/)

### Запуск:
* `git clone https://github.com/hopeurfnis12/PD.git`
* `go get github.com/go-sql-driver/mysql` - скачать MySQL Driver
	> если выходит ошибка:  `go.mod file not found in current directory or any parent directory.`, то попробуйте ввести `go env -w GO111MODULE=off`
* Запустить исполняемый файл: запустить ...

## Roadmap:
### MVP
* Шаблон страницы с успеваемостью
* Логика показа успеваемости
* Шаблон просмотра успеваемости по определенному предмету
* Логика просмотра успеваемости по определенному предмету
* Шаблон страницы с ту-ду листом
* Логика показа ту-ду листа

### Авторизация
* Шаблон страницы регистрации
* Логика регистрации
* Шаблон страницы входа
* Логика авторизации

### Добавление
* Шаблон страницы добавления
* Логика создания

### Поиск
* Шаблон страницы поиска
* Логика поиска

### Редактирование
* Шаблон страницы редактирования
* Логика редактирования

### Удаление/Зачеркивания
* Шаблон страницы удаления
* Логика удаления
