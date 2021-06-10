# go_90test
This is the resolution for given technical challenge.

## RUN

To execute, just clone the repository and run `docker-compose up`

It's also possible to build and run it manually from command line.

## client-api

client-api is a module located on /client-api subfolder. Main package is
under /client-api/cmd

## portdomsvc (Port Domain Service)

portdomsvc module is also organized as a module on /portdomsvc

## ports

default ports are 8088 for client-api rest service and 9090 for grpc server

## rest services

### Endpoints
URL | Description
------------ | -------------
**/v1/port/** | POST or PUT executing a post with a Port object to this endpoint to send and store it
**/v1/port/{KEY}** | GET Retrieves and returns the port with the given code (404 NOT FOUND if port doesn't exist).
**/v1/port/{KEY}** | DELETE Sending a DELETE request to this endpoint will delete port with given key.


### JSON format

Example of the json format needed for POST a net Port object

```json
{
    "key": "BHAHD",
    "country": "Bahrain",
    "province": "Bahrain",
    "city": "Al HIdd",
    "code": "52500",
    "name": "Al Hidd",
    "alias": [],
    "regions": [],
    "unlocs": [
      "BHAHD"
    ]
}

```
