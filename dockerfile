FROM golang:1.13.12-alpine3.12 as base

RUN apk add --no-cache make gcc g++ zeromq zeromq-dev
WORKDIR /app
COPY . .
EXPOSE 4000
RUN go build -o main .

########## DEV STAGE ##########
FROM base as dev
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk add --no-cache git && \
    go get -v github.com/cespare/reflex && \
    go install github.com/cespare/reflex

ENTRYPOINT ["make", "watch"]

######### PROD STAGE #########
FROM alpine as prod
WORKDIR /app
RUN apk add --no-cache make gcc g++ zeromq zeromq-dev
COPY --from=base /app/main /app/.
EXPOSE 4000
CMD ["/app/main"]