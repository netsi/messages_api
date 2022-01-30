##
# build stage
##
FROM golang:1.17 as build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY build/data/messages.csv /messages.csv
COPY api/ ./api
COPY internal/ ./internal
COPY main.go ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /app

##
# run stage
##
FROM alpine:3.15

COPY --from=build /app /
COPY --from=build /messages.csv /build/data/messages.csv

CMD ["/app"]