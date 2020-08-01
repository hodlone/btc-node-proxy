# Development

## Make
The make file provides utility functions for working with this repository.
```
make help
```

## Dependencies
### Reflex
We make use of [reflex](https://github.com/cespare/reflex) which is a utility for hot-reloading, it's old and unmaintained but it's complete and versatile. You'll have to install it locally if you wish to monitor `btc-node-proxy` source files for changes when starting it in the host machine.

Install reflex:
```
go get -v github.com/cespare/reflex && \
go install github.com/cespare/reflex
```

If you get command not found when entering `reflex` into the terminal try creating a symlink to your bin path:
```
sudo ln -s "$GOPATH/bin/reflex" /usr/local/bin/
```

### ZMQ
If you wish to raise `btc-node-proxy` directly in the host machine you'll need to install `ZMQ` which is needed to receive block and transaction events from a `bitcoin-core` node, provided the node has been started with necessarry zmq flags.
