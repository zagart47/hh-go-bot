![Static Badge](https://img.shields.io/badge/%D1%81%D1%82%D0%B0%D1%82%D1%83%D1%81-%D0%B2_%D1%80%D0%B0%D0%B7%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%BA%D0%B5-blue)
![Static Badge](https://img.shields.io/badge/GO-1.21+-blue)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/zagart47/hh-go-bot)
![GitHub last commit (by committer)](https://img.shields.io/github/last-commit/zagart47/hh-go-bot)
![GitHub forks](https://img.shields.io/github/forks/zagart47/hh-go-bot)

# HH-Go-Bot
Проект планирую использовать как помощник по поиску вакансий и отклику на них

## Содержание
- [Технологии](#технологии)
- [Начало работы](#начало-работы)
- [Тестирование](#тестирование)
- [Deploy и CI/CD](#deploy-и-ci/cd)
- [Contributing](#contributing)
- [To do](#to-do)
- [Команда проекта](#команда-проекта)

## Технологии
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [SQLite](https://www.sqlite.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [Docker](https://www.docker.com/)

## Использование
Необходимо зарегистрировать telegam-бота и получить его токен.
Токен прописать в конфиги.
Также необходимо иметь ключи от api.hh.ru для работы с откликами

Собрать проект в docker-образ, запустить контейнер
```powershell
docker build -t hh-go-bot .
docker run --name hh-go-bot -dit --restart unless-stopped -d hh-go-bot
```

Пока поддерживает команды бота
```
/jobs
/similar
/resume
```


## Разработка

### Требования
Для установки и запуска проекта необходим... golang .

## Contributing
Если у вас есть предложения или идеи по дополнению проекта или вышли нашли ошибку, то пишите мне в tg: @zagart47

## FAQ
### Зачем вы разработали этот проект?
Это мой пет-проект.

## To do
- [ ] Добавить возможность работать с вакансиями через http
- [ ] Добавить тесты

## Команда проекта
- [Артур Загиров](t.me/zagart47) — Golang Developer

