# NFT indexer
Nft indexing done right.

## Installation
### Ethereum
#### Converting ABIs to go code
Install the latest version of `abigen` on your system:

```
$ go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

To transform an Ethereum ABI to a go package, see the following example:

`$ abigen --abi ./pkg/indexer/ethereum/abi/ERC721.json --pkg tokens --type Erc721 --out ./pkg/indexer/ethereum/tokens/erc721.go`.

## Components
This project consists of multiple components and microservices in order to achieve maximum scalability.

### Indexer
The NFT indexer processes and stores information about NFT collections.

#### Service 1

### Discovery
The NFT discoverer's job is to find new NFT collections from different sources like OpenSea and blockchain mint events.
Once a new collection has been found, it will submit it to the indexer for processing.

### Public API
HTTP endpoint that enables clients to submit custom collections to be indexed. 

### Message broker
The message broker queues the NFT collections to be processed by the indexer via a FIFO strategy. 
The public API can add items to this queue and the indexer takes them out.
