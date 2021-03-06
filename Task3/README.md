# Task2

Программа для решения задачи коммивояжера. Третье задание по Теории Конечных Графов. Версия 0.7

## Легенда:

В качестве точек следования были выбраны 10 пожарных частей города Уфы.

## Фичи:

* Интерфейс ввода координат реализован посредством xml-файла
* Параллельная реализация вывода
* Реализация алгоритмов "Метод ближайшего соседа" и "Метод муравьиной колонии"
* Алгоритм муравьиной колонии реализован в параллельном варианте
* Отключение вывода в файлы для тестирования
* Автор устал

## Планы:

* Не умереть до конца семестра

## Системные требования:

* OS: Windows x86/x86_64, Linux x86/x86_64
* x86/x86_64 совместимый процессор
* Примерно 10n Free RAM, где n - размер входных файлов
* Для сборки: версия GO не ниже 1.8

## Получение: 

#### Готовый билд: 

[Скачайте](https://drive.google.com/drive/folders/1f592j2D2QWnLc9NYC2LKQIhmo9BGqxw6?usp=sharing) готовый билд для Вашей OS. 

#### Сборка из исходников:

1. Установите [GO](https://golang.org/dl/)
1. Установите [$GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
1. Выполните ```go get github.com/vahriin/BigGraph```
1. Выполните ```go install github.com/vahriin/BigGraph/Task3```
1. В директории ```$GOPATH/bin``` появится бинарный файл ```Task3```, после чего вы можете его запустить из командной строки.

## Использование

Используйте файлы ```adjacency_list.csv``` и ```node_list.csv```, сгенерированные при помощи [Task1](https://github.com/vahriin/BigGraph/tree/master/Task1). Поместите их в директорию ```./input/Task3```. Также в этой директории должны находиться файлы ```travel_points.csv``` и ```point.xml```, представляющие конечные точки и начальную точку соответственно. Примеры файлов вы можете найти [здесь](https://drive.google.com/drive/folders/1etnIJyZGjnMhCC1LwSRabLVASWjRW_Hl?usp=sharing)

Измените аттрибуты ```lat``` и ```lon``` в файле ```point.xml``` в соответствии с широтой и долготой требуемой точки.

Запустите программу из командной строки. Программой будут сгенерированы:

1. Файл ```pathways.csv```. Файл содержит одну строку - путь обхода всех точек, указанных в ```travel_points.csv```.

2. Файл ```road_graph.svg```, в котором на изображение города будет наложен сгенерированный путь. Зеленым кругом обозначена начальная точка. Градиентом от белого к черному - точки, в зависимости от очередности их прохождения.

3. В стандартный вывод будет передана информация о длине сгенерированного пути и времени обхода каждой точки. (Внимание: данный вывод не будет отключен в тестовом режиме).

## Анализ тестирования

Тестировались алгоритмы "Метод ближайшего соседа" и "Метод муравьиной колонии".

### Анализ времени выполнения

Время выполнения оценивалось путем подсчета суммы времени выполнения программы для каждой точки. Несмотря на то, что данную оценку нельзя считать точной, ввиду того, что в данное время входят в том числе и операции ввода-вывода, будем считать, что данная оценка является достаточной для наших целей. При оценке обеспечивалась стабильно низкая нагрузка на процессор со стороны других программ в течении всего времени выполнения.

1. Метод ближайшего соседа: время выполнения алгоритма в среднем составляет 0.8 секунды.

2. Метод муравьиной колонии: время выполнения алгоритма зависит от количества муравьев и времени жизни колонии. Были подобраны минимальные значения данных переменных; время выполнения алгоритма в среднем составило 10-15 секунд.

### Анализ правильности выполнения

Правильность выполнения оценивалась путем сравнения результатов выполнения алгоритмов. Тестирование производилось на пяти равномерно распределенных по городу точках. Сравнивались длины кратчайших путей для каждого алгоритма. Сравнение производилось вручную.

В качестве эталона использовалась длина кратчайшего пути, найденная за несколько десятков итераций муравьиного алгоритма с различными параметрами.

1. Метод ближайшего соседа находил пути с длиной, превышающей эталон на 10-25 тысяч метров, в зависимости от расположения точки старта.

2. Метод муравьиной колонии в среднем (на константных параметрах) находил пути с длиной, превышающей эталон на 5-15 тысяч метров.

Разница длин путей-результатов выполнения обоих алгоритмов составляла 5-15 тысяч метров в пользу "муравьиной колонии".

**Вывод:** Использование метода ближайшего соседа оправдано при необходимости получить быстрый результат. Однако, если время поиска решения не является критичным параметром, лучше использовать метод муравьиной колонии. При этом необходимо учитывать, что для получения оптимального решения методом муравьиной колонии может понадобиться подбирать константы вручную.

## Благодарности
* [Юре](https://github.com/bruce-willis) за найденный "критический" баг.
* Race-detector'у Go за найденный действительно критический баг и все то время, что мы провели вместе.
* [Rob Pike](http://herpolhode.com/rob/) за Round() в Go 1.10
* Традиционно: певице Lily Allen за [песню](https://music.yandex.ru/album/33786/track/322559) для адресата второй благодарности.
