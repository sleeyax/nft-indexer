package zora

var queryAggregateStatTemplate = `
query MyQuery {
  aggregateStat {
	ownerCount(where: { collectionAddresses: "{{.CollectionAddress}}" })
	ownersByCount(
	  where: { collectionAddresses: "{{.CollectionAddress}}" }
	  pagination: { limit: {{.TopOwnersLimit}} }
	) {
	  nodes {
		count
		owner
	  }
	}
	salesVolume(where: { collectionAddresses: "{{.CollectionAddress}}" }) {
	  chainTokenPrice
	  totalCount
	  usdcPrice
	}
	nftCount(where: { collectionAddresses: "{{.CollectionAddress}}" })
  }
}`

var queryTokensTemplate = `
        query MyQuery {
          tokens(where: { collectionAddresses: "{{.CollectionAddress}}"}, networks: {network: ETHEREUM, chain: MAINNET}, pagination: {after: "{{.After}}", limit: {{.Limit}}}, sort: {sortKey: TOKEN_ID, sortDirection: ASC}) {
            nodes {
              token {
                tokenId
                tokenUrl
                attributes {
                  displayType
                  traitType
                  value
                }
                image {
                  url
                }
                mintInfo {
                  toAddress
                  originatorAddress
                  price {
                    chainTokenPrice {
                      decimal
                      currency {
                        address
                        decimals
                        name
                      }
                    }
                  }
                  mintContext {
                    blockNumber
                    transactionHash
                    blockTimestamp
                  }
                }
              }
            }
            pageInfo {
              endCursor
              hasNextPage
              limit
            }
          }
        }
      `
