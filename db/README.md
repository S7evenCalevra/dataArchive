# README

## Requirements

docker run -e AUTH_USERNAME='testguy' -e AUTH_PASSWORD='123Security' -it -p 8080:4000 a99acd8b2888

Genreal Flow :
SNOW Service catalog --> sends http request with captured values as post to API --> API takes information, marshals to a struct and sends to db -> returns status of request

May require a get handler to fetch and check if data exists?