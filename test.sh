#!/usr/bin/env bash

rm test.out
go build -o test.out
./test.out -d 1 -t 100000
./test.out -d 1 -t 100000 -p
./test.out -d 2 -t 100000
./test.out -d 2 -t 100000 -p
./test.out -d 3 -t 100000
./test.out -d 3 -t 100000 -p
./test.out -d 4 -t 100000
./test.out -d 4 -t 100000 -p
./test.out -d 5 -t 100000
./test.out -d 5 -t 100000 -p
./test.out -d 6 -t 100000
./test.out -d 6 -t 100000 -p
./test.out -d 7 -t 100000
./test.out -d 7 -t 100000 -p
./test.out -d 8 -t 100000
./test.out -d 8 -t 100000 -p
./test.out -d 9 -t 100000
./test.out -d 9 -t 100000 -p
./test.out -d 10 -t 10000
./test.out -d 10 -t 10000 -p
./test.out -d 11 -t 10000
./test.out -d 11 -t 10000 -p
./test.out -d 12 -t 10000
./test.out -d 12 -t 10000 -p
./test.out -d 13 -t 10000
./test.out -d 13 -t 10000 -p
./test.out -d 14 -t 10000
./test.out -d 14 -t 10000 -p
./test.out -d 15 -t 10000
./test.out -d 15 -t 10000 -p
./test.out -d 16 -t 10000
./test.out -d 16 -t 10000 -p
rm test.out