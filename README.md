# Development

## Make
The make file provides utility functions for working with this repository.
```
make help
```

## Dependencies
### Reflex
We make use of [reflex](https://github.com/cespare/reflex) which is a utility for hot-reloading, it's old and unmaintained but it's complete and versatile. You'll have to install it locally if you wish to monitor `btc-node-proxy` source files for changes when starting it in the host machine.

### ZMQ
If you wish to raise `btc-node-proxy` directly in the host machine you'll need to install `ZMQ` which is needed to receive block and transaction events from a `bitcoin-core` node, provided the node has been started with necessarry zmq flags.

## Development
### Host

To start the service from the host simply do `make watch`, this will start the proxy and any changes made to any files ending with `.go` will trigger a rebuild of the service.

See Dependencies section for details on how to start on the host.

### Docker


