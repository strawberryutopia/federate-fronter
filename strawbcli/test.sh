#!/bin/bash

echo ====================
echo Testing...
echo ====================


echo "Lucy"
go run main.go federate --name Lucy --avatar "https://pbs.twimg.com/media/Ec6Zv21WsAAIwOS?format=png&name=medium"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name

echo
echo
echo "Ivie"
go run main.go federate --name Ivie --avatar "https://ca.slack-edge.com/T04FFM9U7-U4UBY7Y5V-02da57a94a10-512"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name

echo
echo
echo "Jesper"
go run main.go federate --name Jesper --avatar "https://i.imgur.com/4vlRY23.png"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name

echo
echo
echo "No Config Set"
go run main.go federate --name "No Config Set" --avatar "https://www.pngitem.com/pimgs/m/575-5759580_anonymous-avatar-image-png-transparent-png.png"
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name

# TODO: Test nil

echo
echo
echo "From PK"
go run main.go federate
go run main.go federate check strawb | jq .real_name
go run main.go federate check pcw | jq .real_name
