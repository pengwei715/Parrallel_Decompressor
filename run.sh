#!bin/bash

time ./decompressor <./data/txts/1.txt >>small.txt

for i in 1 2 4 6 8 10 12
do
    time ./decompressor -p $i <./data/txts/1.txt >>small.txt
done

time ./decompressor<./data/txts/2.txt >>meduim.txt

for i in 1 2 4 6 8 10 12
do
    time ./decompressor -p $i <./data/txts/2.txt >>meduim.txt
done


time ./decompressor <./data/txts/3.txt >>large.txt

for i in 1 2 4 6 8 10 12
do
    time ./decompressor -p $i <./data/txts/3.txt 2>>large.txt
done


time ./decompressor <./test_data/test.txt >>test.txt

for i in 1 2 4 6 8 10 12
do
    time ./decompressor -p $i <./test_data/test.txt >>test.txt
done

