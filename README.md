# go-url-shortener

This is a simple URL shortener written in Go. It allows generating short aliases for long URLs.

## Features

- Generate short URLs given a long URL
- Redirect short URLs to the original long URL

## Usage

To start the server:

```
go run main.go
```

The server will start on port 8080 by default and open a simple ui on localhost:8080/ui. You can use the ui to generate short URLs and delete them.

Generate a short URL using cli:

```
curl -X POST -d '{"url":"http://long.url/here"}' http://localhost:8080/api/shorten
```

This will return a JSON response with the short URL:

```json
{
  "id": "FGm3TgWAjDL1gdv3Lv3i4a4XlcVxsqCUBPLDHEH4xyz1nfqDrp7Tg2jiFc7KMs67kN1mG4EESl5SwAnBE4VthLc3l3CZvsAHsLp9",
  "url": "http://long.url/here",
  "shortUrl": "7NDuXo",
  "created": "2024-01-12T20:38:09.6147719+01:00"
}
```

Generate a short URL using cli:

```
curl -X DELETE http://localhost:8080/api/delete/{id}
```

To redirect a short URL:

Just send a GET request to the short URL path, and you will be redirected to the original long URL.
