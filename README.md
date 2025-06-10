# Clicker-Go

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/NKtym/Clicker-Go)
[![Go Report Card](https://goreportcard.com/badge/github.com/NKtym/Clicker-Go)](https://goreportcard.com/report/github.com/NKtym/Clicker-Go)
[![Code Size](https://img.shields.io/github/languages/code-size/NKtym/Clicker-Go?style=flat-square)](https://github.com/NKtym/Clicker-Go)
![GitHub Release](https://img.shields.io/github/v/release/NKtym/Clicker-Go?display_name=release)
___
**Clicker-Go** - Игра в которой вам предстоит кликать!!! В ней вы сможете прокачиваться и убивать боссов. Отличная игра чтобы скоротать время.

## Содержание
- [Clicker-Go](#clicker-go)
  - [Содержание](#содержание)
  - [Использование](#использование)
  - [Кнопки](#кнопки)
  - [TO DO](#to-do)

## Использование
С официального сайта установите библиотеку для создания 2D игр https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://ebitengine.org/&ved=2ahUKEwiSiIatheSNAxW0DRAIHW-oH3wQFnoECBgQAQ&usg=AOvVaw1gGtMYv0YB60w-maL3lNeu (также у вас уже должен быть установлен go версии не ниже 1.22)

После склонируйте репозиторий и соберите проект:
```shell
git clone https://github.com/NKtym/Clicker-Go.git
go mod tidy
go run .
```

## Кнопки
- Space - получение очков/нанесение урона боссам(урон по боссам складывается с уровнем вашего клика)
- S - открыть магазин
- F - меню выора боссов
- K - сохранение прогресса в один из трех сейвов
- L - загрузка прогресса из одного из трех сейвов
- Y - подтверждение действия
- N - отказ от действия
- 1 - выбор босса 1/прокачка кликов/сохранение/загрузка сохранения
- 2 - выбор босса 2/прокачка автокликов/сохранение/загрузка сохранения
- 3 - выбор босса 3/получение подарков/сохранение/загрузка сохранения
- Tab - статистика игры
- Esc - закрыть приложение

## TO DO
- ✅ Readme file
- ✅ Основной функционал кликов
- ✅ Магазин
- ✅ Битвы с боссами
- ✅ Статистика игрока
- ✅ Кнопка выхода из игры
- ❌ Сохранение прогресса
- ❌ Достижения
- ❌ Подарки
- ❌ Скины