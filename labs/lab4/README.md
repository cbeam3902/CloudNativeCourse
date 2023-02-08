# Lab 4
This lab introduces how to make a basic webserver using GO. It implements a simple map based database to keep track of prices of items and can be changed by making a simple url request.

To demo the code, run `go run webserver.go` and in a different terminal do `curl "localhost:8000/list"`to bring up the list of items in the database or  or `curl "localhost:8000/create?item=Piano&price=500"` to create a new item in the database.

Different types of calls:
1. list
2. price - Need: item
3. create - Need: item, price
4. update - Need: item, price
5. delete - Need: item