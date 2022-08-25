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
