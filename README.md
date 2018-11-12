# Optimal process planning

An algorithm aimed at optimal processes planning is presented to your attention. Formally, we describe the problem.
Let there be ![1](https://latex.codecogs.com/svg.latex?n) files that need to be processed by preloading into
memory. For memory, the limit is put forward in the form of a total, simultaneously used memory equal to
![2](https://latex.codecogs.com/svg.latex?M). Then, at a time, there may be files in memory, the total volume
which does not exceed ![3](https://latex.codecogs.com/svg.latex?M). I.e.

![4](https://latex.codecogs.com/svg.latex?\sum\limits_{i%20\in%20\text{Memory}}%20m_i%20\leq%20M)

where![5](https://latex.codecogs.com/svg.latex?m_i) is the maximum amount of memory that will be used for processing
![7](https://latex.codecogs.com/svg.latex?i) -th file.

In addition, at the same time there can be no more memory in ![6](https://latex.codecogs.com/svg.latex?C) files, i.e.
![7](https://latex.codecogs.com/svg.latex?|\text{Memory}%20|%20\leq%20C).

Our task is to minimize the total execution time, given that the set of files is updated after complete the processing 
of one of the previous files. The algorithm for selecting the following file is as follows:

* Select all possible files in decreasing order of processing time
* If the next file does not fit into memory, then the file of the maximum size from the possible
* If no files can be added, it is proposed to process the files

# Оптимальное планирование процессов

Вашему вниманию представляется алгоритм, нацелнный на оптимальное планирование процессов. Формально опишем задачу. 
Пусть есть ![1](https://latex.codecogs.com/svg.latex?n) файлов, которые требуется обработать, предварительно загрузив в
память. Для памяти выдвигается ограничение в виде суммарно, одновременно используемой памяти, равной 
![2](https://latex.codecogs.com/svg.latex?M). Тогда, единовременно, в памяти могут находиться файлы, суммарный объём 
которых не превышает ![3](https://latex.codecogs.com/svg.latex?M). Т.е. 

![4](https://latex.codecogs.com/svg.latex?\sum\limits_{i%20\in%20\text{Memory}}m_i\leq%20M)

где ![5](https://latex.codecogs.com/svg.latex?m_i) - максимальный объём памяти, который будет использоваться для обработки
![7](https://latex.codecogs.com/svg.latex?i)-го файла. 

Кроме того, одновременно в памяти не может находиться более ![6](https://latex.codecogs.com/svg.latex?C) файлов, т.е.
![7](https://latex.codecogs.com/svg.latex?|\text{Memory}|\leq C). 

В нашей задаче требуется минимизировать суммарное время выполнения, учитывая, что набор файлов обновляется после 
завершения обработки одного из предыдущих файлов. Алгоритм выбора следующего файла состоит в следующем:

* Выбрать все возможные файлы в порядке убывания времени их обработки
* Если следующий файл не умещается в память, то выбирается файл максимального объёма из возможных
* Если ни одни файл добавить не удаётся, то предлагается обработать файлы 

# Literature
1. Job shop scheduling, [Wikipedia](https://en.wikipedia.org/wiki/Job_shop_scheduling)
2. Planning, [Wikipedia](https://en.wikipedia.org/wiki/Planning)