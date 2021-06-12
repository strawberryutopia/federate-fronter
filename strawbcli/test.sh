#!/bin/bash

echo ====================
echo Testing...
echo ====================

# TODO: This is okay for now, but really this aught to be some
# proper Go tests, with mocks and stuff


echo "Lucy"
go run main.go -d federate --name Lucy --avatar "https://pbs.twimg.com/media/Ec6Zv21WsAAIwOS?format=png&name=medium"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name
go run main.go federate check work | jq .real_name

echo
echo
echo "Ivie"
go run main.go -d federate --name Ivie --avatar "https://ca.slack-edge.com/T04FFM9U7-U4UBY7Y5V-02da57a94a10-512"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name
go run main.go federate check work | jq .real_name

echo
echo
echo "Jesper"
go run main.go -d federate --name Jesper --avatar "https://i.imgur.com/4vlRY23.png"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name
go run main.go federate check work | jq .real_name

echo
echo
echo "Strawb System (e.g. nil fronter)"
go run main.go -d federate --name "Strawb System" --avatar "https://www.pngitem.com/pimgs/m/575-5759580_anonymous-avatar-image-png-transparent-png.png"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name
go run main.go federate check work | jq .real_name

# TODO: Test nil

echo
echo
echo "From PK"
go run main.go federate
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name
go run main.go federate check work | jq .real_name
