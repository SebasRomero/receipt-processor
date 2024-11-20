
# Receipt processor

Hi, to run this project you must go to the root folder and generate a docker build:

## Commands to run the project

```
$ docker build -t receipt-app .
```

Then, to run the project you should use this command:

```
$ docker run -p 8081:8080 -d receipt-app
```

The default port is 8080.

## Routes

There are 4 different routes in the application.

GET /health -> It should return an ok.

GET /api/v1/receipts -> It should return all the receipts.

GET /api/v1/receipts/{id}/points -> It should return the points from that receipt.

POST /api/v1/receipts/process -> It should save the receipt and return the id.

## Test

I did just one test using the examples from the read me from the repository, so to run it you can use:

```
go test ./...
```
