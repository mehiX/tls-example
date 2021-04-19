# TLS example

TLS practice following [Liz Rice - Secure connections](https://github.com/lizrice/secure-connections)

External tool: [MiniCA](https://github.com/jsha/minica)

## Run the project

Prepare the environment:

```bash
# Install minica if not done so already
go get github.com/jsha/minica

# generate certificates in a separate directory (./ca)
mkdir ca && cd ca && minica -d local.host,local.client

# add local.host to /etc/hosts
sudo echo "127.0.0.1 local.host" >> /etc/hosts
```

Run the server:

```bash 
go get -u ./...
go run ./cmd/server local.host:7878
```

Run the client:
```bash
go run ./cmd/client https://local.host:7878
```

The client and the server will pause at various stages asking for user input to proceed.