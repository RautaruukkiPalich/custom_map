# Кастомная мапа

## Описание
Аналог sync.Map с конкурентным доступом

## Установка
```sh
go get github.com/rautaruukkipalich/custom_map/@latest
```


## Использование

### Создание мапы
```
newMap := custommap.NewMap()
```
### Методы

- ```SET(string, any)``` добавление пары ключ/значение в мапу
- ```GET(string) (any, bool)``` получение значения из мапы по ключу
- ```LEN() int``` получение количества элементов мапы

## Тестирование
### Сравнение с syncMap
```
make test
```
#### без детектора гонок
![alt test1](https://github.com/rautaruukkipalich/custom_map/blob/main/img/test_result_new.PNG?raw=true)
#### c детектором гонок (флаг -race)
![alt test_race](https://github.com/rautaruukkipalich/custom_map/blob/main/img/test_result_race_new.PNG?raw=true)


## Бенчмарки
### Сравнение с syncMap
```
make bench
```
![alt bench](https://github.com/rautaruukkipalich/custom_map/blob/main/img/bench_new.PNG?raw=true)
