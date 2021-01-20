#! /bin/bash
let num=0
while test ${num} -le 10; do
    go run ./mutex1.go
    ((num++))
done
