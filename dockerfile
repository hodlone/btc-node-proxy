FROM gcr.io/nodehodl/go-base:ssh as base

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
RUN apk add --no-cache gcc g++ zeromq zeromq-dev
RUN addgroup -g 2468 simps && \
    adduser -s /bin/sh -G simps -u 2468 -D liluser
WORKDIR /home/liluser/app
USER liluser
COPY --from=base /app/main /home/liluser/app/.
EXPOSE 4000
CMD ["/home/liluser/app/main"]