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

type AggregateStatResponse struct {
	Data struct {
		AggregateStat AggregateStat `json:"aggregateStat"`
	} `json:"data"`
}

type TokensResponse struct {
	Data struct {
		Tokens struct {
			Nodes    []TokenWrapper `json:"nodes,omitempty"`
			PageInfo struct {
				EndCursor   string `json:"endCursor,omitempty"`
				HasNextPage bool   `json:"hasNextPage,omitempty"`
				Limit       int    `json:"limit,omitempty"`
			} `json:"pageInfo,omitempty"`
		} `json:"tokens,omitempty"`
	} `json:"data,omitempty"`
}

type Token struct {
	CollectionName   string      `json:"collectionName,omitempty"`
	TokenId          string      `json:"tokenId,omitempty"`
	Name             string      `json:"name,omitempty"`
	Owner            string      `json:"owner,omitempty"`
	Content          Content     `json:"content"`
	Image            Content     `json:"image"`
	Description      string      `json:"description,omitempty"`
	TokenUrl         string      `json:"tokenUrl,omitempty"`
	TokenUrlMimeType string      `json:"tokenUrlMimeType,omitempty"`
	Attributes       []Attribute `json:"attributes,omitempty"`
	MintInfo         struct {
		MintContext struct {
			BlockNumber     int    `json:"blockNumber,omitempty"`
			BlockTimestamp  string `json:"blockTimestamp,omitempty"`
			TransactionHash string `json:"transactionHash,omitempty"`
		} `json:"mintContext"`
		Price struct {
			ChainTokenPrice struct {
				Currency struct {
					Address  string `json:"address,omitempty"`
					Decimals int    `json:"decimals,omitempty"`
					Name     string `json:"name,omitempty"`
				} `json:"currency"`
				Decimal float64 `json:"decimal,omitempty"`
			} `json:"chainTokenPrice"`
		} `json:"price"`
		ToAddress         string `json:"toAddress,omitempty"`
		OriginatorAddress string `json:"originatorAddress,omitempty"`
	} `json:"mintInfo"`
}

type TokenWrapper struct {
	Token Token `json:"token,omitempty"`
}

type Node struct {
	Count int    `json:"count,omitempty"`
	Owner string `json:"owner,omitempty"`
}

type TopOwnersByCount struct {
	Nodes []Node `json:"nodes,omitempty"`
}
type Content struct {
	Url           string        `json:"url,omitempty"`
	Size          string        `json:"size,omitempty"`
	MimeType      string        `json:"mimeType,omitempty"`
	MediaEncoding MediaEncoding `json:"mediaEncoding"`
}

type MediaEncoding struct {
	Large     string `json:"large,omitempty"`
	Poster    string `json:"poster,omitempty"`
	Preview   string `json:"preview,omitempty"`
	Original  string `json:"original,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type Attribute struct {
	Value     string `firestore:"value,omitempty" json:"value,omitempty"`
	TraitType string `firestore:"traitType,omitempty" json:"traitType,omitempty"`
}
