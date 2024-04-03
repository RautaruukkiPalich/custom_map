# Кастомная мапа

## Описание
Аналог sync.Map с конкурентным доступом

## Установка
```sh
go get github.com/RautaruukkiPalich/custom_map
```


## Использование

### Создание мапы
```
newMap := custommap.NewMap()
```
### Методы

- ```SET(string, any)``` созранение пары ключ/значение в мапу
- ```GET(string) (any, bool)``` получение значения из мапы по ключу
- ```LEN() int``` получение количества элемнетов мапы

## Тестирование
### Сравнение с syncMap
```
make test
```
#### без детектора гонок (флага -race)
![alt test1](img/test_result.png)
#### c детектором гонок (флагом -race)
![alt test_race](img/test_result_race.png)


## Бенчмарки
Сравнение с syncMap
```
make bench
```
![alt bench](img/bench.png)
