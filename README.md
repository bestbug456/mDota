# mDota
mDota is a Machine learning alghoritm which can identify your role using your stat and in the future will be able to create a balance team


# What is mDota
mDota is a project based on golang (https://golang.org) and the SVM library libsvm-go (https://github.com/ewalker544/libsvm-go).

# How to use mDota
At the moment the only mode to use the mDota alghoritm is via cli, run the compiled code with --help to see the list of command:
* Analyze: this command take in input a trainset and at least 1 user and return the role for each user
* Big data: using golang feature to permit analysis of a big number of user in a low time, 

# Benchmark
this are the benchmark result done on our machine for the bigData command:

BenchmarkBigData100Request-4    	     500	   3268394 ns/op
BenchmarkBigData1000Request-4   	     500	   4327609 ns/op
BenchmarkBigData10000Request-4  	     300	   4128752 ns/op
BenchmarkBigData100000Request-4 	     200	   5993765 ns/op
BenchmarkBigData1000000Request-4	     500	 121815031 ns/op

which mean the single request is done in 100 nanosecond circa, not bad in our opinion! :D


# Our goal
The scope of this project is to give to the comunity of dota a tool to help to identify the best role to play a better game, and maybe give to valve a idea how to improve their research alghoritm

# How to contribute
If you want to contribute to this project you can fork the from develop branch and add your code, when you have done create a pull request to the develop branch again.

# Authors
The change to the libsvm and the whole code (excluding the _vendor) was created by Danilo 'Bestbug' and Franca 'Forcolotta'
