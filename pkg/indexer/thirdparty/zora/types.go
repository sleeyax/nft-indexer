package zora

type SalesVolume struct {
	ChainTokenPrice float64 `json:"chainTokenPrice"`
	TotalCount      int     `json:"totalCount"`
	UsdcPrice       float64 `json:"usdcPrice"`
}

type AggregateStat struct {
	OwnerCount    int              `json:"ownerCount,omitempty"`
	OwnersByCount TopOwnersByCount `json:"ownersByCount"`
	SalesVolume   SalesVolume      `json:"salesVolume"`
	NftCount      int              `json:"nftCount,omitempty"`
}

type GraphQLResponse struct {
	Data struct {
		AggregateStat AggregateStat `json:"aggregateStat"`
	} `json:"data"`
}

type Node struct {
	Count int    `json:"count,omitempty"`
	Owner string `json:"owner,omitempty"`
}

type TopOwnersByCount struct {
	Nodes []Node `json:"nodes,omitempty"`
}
