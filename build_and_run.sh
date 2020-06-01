#!bin/bash
go build mpcs52060/proj3/src/decompressor

for h in 4 3 2 1 0
do
    time ./decompressor <./test_data/test.txt
done



for i in 1 2 4 6 8
do
    for j in 4 3 2 1 0
    do
	time ./decompressor -p $i < ./test_data/test.txt
    done
done
