# NFT indexer
Concurrent NFT indexer.


> [!WARNING]  
>  **No longer maintained**
> 
> I used to work on this project as an experimental side-project while working on the [Infinity NFT marketplace](https://github.com/infinitydotxyz) back in 2022.
> Given the project has long since been discontinued and its source code disclosed, I decided to make this project public as well.
> Batteries are not included as some things may be missinng or broken. If you wish to continue this project or reuse parts of it then you're on you own :) 

## Installation
Requires go version 1.18+ to be installed on your system.

```
$ git clone https://github.com/sleeyax/nft-indexer.git
$ go mod download
```

To get started: `$ go run ./cmd/indexer/main.go`.

Run tests: `$ go test ./...`.

**JetBrain's GoLand IDE is recommended for development.**

## Services
This project consists of multiple components in order to achieve decent scalability and separation of concerns.

### Indexer
The NFT indexer processes and stores information about NFT collections.

### Discoverer
The NFT discoverer's job is to find new NFT collections from different sources like OpenSea and blockchain mint events.
Once a new collection has been found, it could submit it to the indexer for processing.

### Public API
HTTP endpoint that enables clients to submit custom collections to be indexed. 
This endpoint is also accessible via a WebSocket connection, so developers can consume indexing events live as the collection is being indexed. 

### Message broker
The message broker queues the NFT collections to be processed by the indexer via a FIFO strategy. 
The public API can add items to this queue and the indexer takes them out. 
This system shouldn't be implemented entirely from scratch; production-ready solution like Kafka, RabbitMQ or GCP queue already exist.

## Guides
### Ethereum
#### Converting ABIs to go code
Install the latest version of `abigen` on your system:

```
$ go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

To transform an Ethereum ABI to a go package, see the following example:

`$ abigen --abi ./pkg/indexer/ethereum/abi/ERC721.json --pkg tokens --type Erc721 --out ./pkg/indexer/ethereum/tokens/erc721.go`.
