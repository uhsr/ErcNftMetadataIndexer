# ErcNftMetadataIndexer

## Description



## Features

- Indexes ERC-721 and ERC-1155 token metadata using event logs and off-chain data sources.
- Stores indexed NFT metadata in a PostgreSQL database optimized for efficient querying.
- Exposes a GraphQL API for retrieving NFT metadata based on various filters like token ID, owner, and attributes.
- Caches frequently accessed NFT metadata in Redis to reduce database load and improve response times.
- Automatically refreshes NFT metadata when token URIs are updated or ownership is transferred.
- Supports configurable rate limiting and API key authentication to prevent abuse.
- Implements a robust error handling and logging mechanism for debugging and monitoring.
- Deploys with Docker and Kubernetes for scalable and resilient infrastructure.
## Installation

```bash
go get github.com/uhsr/ErcNftMetadataIndexer
```

## Usage

```go
package main

import (
    "ercnftmetadataindexer/internal/ercnftmetadataindexer"
)

func main() {
    app := ercnftmetadataindexer.NewApp(false)
    app.Run()
}
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
