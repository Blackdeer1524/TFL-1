# Лабораторные работы по курсу "Теория Формальных Языков"

**Решения лежат в соответствующих git ветках**

Выполнил: Слесарев Данила, ИУ9-51Б

# ВАРИАНТ 2, 4, 5

**ВСЕ ВХОДНЫЕ ДАННЫЕ ПОДАЮТСЯ В `input.txt`. В нём также лежит пример входа**

Формат задания информации о  грамматике:

Сначала через запятую указываются терминалы (их длина только 1),
затем указываются правила переписывания (нетерминаты могут иметь длину больше 1). 
Задача правил происходит до встречи пустой строки.

после подаются длина в строкак и сама w_0 (аналагично с w_1)

# Установка библиотек

```sh
go mod tidy
```

# Запуск
```
STRAT=greedy go run cmd/cli/main.go
```
