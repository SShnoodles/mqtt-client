# mqtt client

if you deploy mosquitto server, then deploy this.

## Download 2 files

* mqtt-client
* config.yml

## Run

./mqtt-client

## Use

pub "test/request" => sub "test/response"

## Build

GOOS=linux GOARCH=amd64 go build -o mqtt-client