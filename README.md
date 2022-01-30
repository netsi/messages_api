# Messages API

Message API is a simple go application that exposes 4 routes.

1 public:

* POST /api/v1/message

3 private:

* GET /api/v1/message/{message_id}
* PUT /api/v1/message/{message_id}
* GET /api/v1/messages

the private routes are protected with basic authentication. The project uses gin-gonic which is a popular
well-documented web framework. Additionally it uses a library for generating UUIDs.


As you can see I'm storing everything in memory. For production I would have used a NoSQL database like mongoDB to
store the `users` and the `messages`. To make the transition in the future easier I've created interfaces for the repositories
so that the only change we would need to do to use an actual Database would be to implement that interface.

Normally I would not put a binary file (see /build/data/messages.csv into the repository) I noticed though that I could
not download the file from Google Drive directly using a client like google.golang.org/api/drive/v3 so to avoid this
I've added it into the repo.

## Prerequisites

You'll have to install docker to be able to run the application locally see
the [official site](https://docs.docker.com/get-docker/) for the instructions on how to install it.

## Running Locally

You can simply start the local environment by calling `make start` on the project root folder. Then you can
use `make stop` to stop the application.

## Endpoints

### Send Message Request

`curl -v -X POST http://localhost:8080/api/v1/message -H 'Content-Type: application/json' -d '{"name":"some-name", "email":"some-email@gmail.com", "text":"some-text"}`

### Send Message Response

```
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Sun, 30 Jan 2022 15:31:17 GMT
< Content-Length: 168
```

```json
{
  "id": "131FE0E8-310C-48A2-A556-A85EAF8E94FA",
  "name": "some-name",
  "email": "some-email@gmail.com",
  "text": "some-text",
  "creation_date": "2022-01-30T16:31:17.973082776+01:00"
}
```

### Get Message Request

`curl -v -X GET http://localhost:8080/api/v1/message/131FE0E8-310C-48A2-A556-A85EAF8E94FA -u admin:back-challenge`

### Get Message Response

```
HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Sun, 30 Jan 2022 15:32:59 GMT
< Content-Length: 168
```

```json
{
  "id": "131FE0E8-310C-48A2-A556-A85EAF8E94FA",
  "name": "some-name",
  "email": "some-email@gmail.com",
  "text": "some-text",
  "creation_date": "2022-01-30T16:31:17.973082776+01:00"
}
```

### Update Message Request

`curl -v -X PUT http://localhost:8080/api/v1/message/131FE0E8-310C-48A2-A556-A85EAF8E94FA -u admin:back-challenge -H 'Content-Type: application/json' -d '{"text":"updated-text"}'`

### Update Message Response

```
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Sun, 30 Jan 2022 15:34:15 GMT
< Content-Length: 171
```

```json
{
  "id": "131FE0E8-310C-48A2-A556-A85EAF8E94FA",
  "name": "some-name",
  "email": "some-email@gmail.com",
  "text": "updated-text",
  "creation_date": "2022-01-30T16:31:17.973082776+01:00"
}
```

### Get Messages Request

`curl -v -X GET http://localhost:8080/api/v1/messages -u admin:back-challenge`

### GetMessages Response

```json
{
  "message": [
    {
      "id": "32B20FDD-C578-481A-B559-E3171FA135DD",
      "name": "some-name",
      "email": "some-email@gmail.com",
      "text": "some-text300",
      "creation_date": "2022-01-30T16:46:20.90391772+01:00"
    },
    {
      "id": "FE7F018C-EF74-40A7-9E55-54720D75354A",
      "name": "some-name",
      "email": "some-email@gmail.com",
      "text": "some-text299",
      "creation_date": "2022-01-30T16:46:19.59191301+01:00"
    }
  ],
  "next_page": "http://localhost:8080/api/v1/messages?offset=2",
  "total_count": 300
}
```