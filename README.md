# Decompressor

This software decompresses all the compressed files inside of one repo. Now it support three compressed files: .zlib .bz2 .gz. 

## Get started

There are three bash scripts.

### Step 1: build and run the tests

```
bash build_and_run.sh
```
It will build the decompressor from the source code and run the sytem with test data (very small)


### Step 2: get all the big file from uchiago box

```
bash get_files.sh

```
Then you can find the data in ./data/
Inside the data dictionary, there are small data (kb), medium data (~10mb) , large data(~200mb)

### Step 3: evaluate the performance of the system with the large data

```
bash run.sh
```
This will run 4 different size of data, small, medium and large, these three contains the same number of files init. Test data contains much more small files in it.

The result is stored in small.txt, medium.txt, large.txt, test.txt

## Prerequisites

To keep the code clean, this software doesn't use any third-party libray


### Usage of the system

#### Parallel version
```
./decompressor -p int <txtfile
```

-p flag indictate the maxinum cores that the system will use when decompressing 
each line of the txtfile is the root of the compressed file's location

#### Sequential version
```
./decompressor <txtfile
```

### Running the tests

Inside src there are 3 packages:

decomprossor is the main package

To run the tests for the decomprossor

```
go test -bench=.
```

pipeline is the package that implements pipeline pattern and fanIn-fanOut pattern
util package has some basic function to support the other two

To run the tests for these two packages

```
go test -v
```