RUN apk add --no-cache zeromq-dev musl-dev pkgconfig alpine-sdk libsodium-dev
RUN CGO_LDFLAGS="$CGO_LDFLAGS -lstdc++ -lm -lsodium" \
  CGO_ENABLED=1 \
  GOOS=linux \
  go build -v -a --ldflags '-extldflags "-static" -v'