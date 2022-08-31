package reservoir

type DetailedTokensResponse struct {
	Tokens []TokenWrapper `json:"tokens,omitempty"`

	// Cursor.
	Continuation string `json:"continuation,omitempty"`
}

type TokenWrapper struct {
	Token Token
}

type Token struct {
	Contract    string      `json:"contract,omitempty"`
	TokenId     string      `json:"tokenId,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Image       string      `json:"image,omitempty"`
	Kind        string      `json:"kind,omitempty"`
	Owner       string      `json:"owner,omitempty"`
	Collection  Collection  `json:"collection"`
	Attributes  []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Key        string `json:"key,omitempty"`
	Value      string `json:"value,omitempty"`
	TokenCount int    `json:"tokenCount,omitempty"`
}

type Collection struct {
	Id              string                `json:"id,omitempty"`
	Name            string                `json:"name,omitempty"`
	Slug            string                `json:"slug,omitempty"`
	Image           string                `json:"image,omitempty"`
	Metadata        CollectionMetadata    `json:"metadata"`
	SampleImages    []string              `json:"sampleImages,omitempty"`
	TokenCount      int                   `json:"tokenCount,omitempty,string"`
	OwnerCount      int                   `json:"ownerCount,omitempty,string"`
	OnSaleCount     int                   `json:"onSaleCount,omitempty,string"`
	FloorAsk        CollectionFloorAsk    `json:"floorAsk"`
	TopBid          CollectionTopBid      `json:"topBid"`
	Rank            CollectionPeriodStat  `json:"rank"`
	Volume          CollectionPeriodStat  `json:"volume"`
	ColumeChange    CollectionPeriodStat  `json:"columeChange"`
	FloorSale       CollectionPeriodStat  `json:"floorSale"`
	FloorSaleChange CollectionPeriodStat  `json:"floorSaleChange"`
	Attributes      []CollectionAttribute `json:"attributes,omitempty"`
}

type CollectionMetadata struct {
	ImageUrl        string `json:"imageUrl,omitempty"`
	DiscordUrl      string `json:"discordUrl,omitempty"`
	Description     string `json:"description,omitempty"`
	ExternalUrl     string `json:"externalUrl,omitempty"`
	BannerImageUrl  string `json:"bannerImageUrl,omitempty"`
	TwitterUsername string `json:"twitterUsername,omitempty"`
}

type CollectionFloorAsk struct {
	Id         string `json:"id,omitempty"`
	Price      int    `json:"price,omitempty"`
	Maker      string `json:"maker,omitempty"`
	ValidFrom  int    `json:"validFrom,omitempty"`
	ValidUntil int    `json:"validUntil,omitempty"`
	Token      struct {
		Contract string `json:"contract,omitempty"`
		TokenId  string `json:"tokenId,omitempty"`
		Name     string `json:"name,omitempty"`
		Image    string `json:"image,omitempty"`
	} `json:"token"`
}

type CollectionTopBid struct {
	Id         string `json:"id,omitempty"`
	Value      int    `json:"value,omitempty"`
	Maker      string `json:"maker,omitempty"`
	ValidFrom  int    `json:"validFrom,omitempty"`
	ValidUntil int    `json:"validUntil,omitempty"`
}

type CollectionPeriodStat struct {
	OneDay    int `json:"1day,omitempty"`
	SevenDay  int `json:"7day,omitempty"`
	ThirtyDay int `json:"30day,omitempty"`
	AllTime   int `json:"allTime,omitempty"`
}

type CollectionAttribute struct {
	Key   string `json:"key,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Count int    `json:"count,omitempty"`
}
