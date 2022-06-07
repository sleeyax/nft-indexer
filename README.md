# NFT indexer
Nft indexing done right.

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
