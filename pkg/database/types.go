package database

type TokenStandard string

const (
	ERC721  TokenStandard = "ERC721"
	ERC1155 TokenStandard = "ERC1155"
)

type DisplayType string

const (
	Date            DisplayType = "date"
	Number          DisplayType = "number"
	BoostNumber     DisplayType = "boost_number"
	BoostPercentage DisplayType = "boost_percentage"
)

type CreationFlow string

const (
	CollectionCreator  CreationFlow = "collection-creator"
	CollectionMetadata CreationFlow = "collection-metadata"
	CollectionMints    CreationFlow = "collection-mints"
	TokenMetadata      CreationFlow = "token-metadata"
	TokenMetadataUri   CreationFlow = "token-metadata-uri"
	AggregateMetadata  CreationFlow = "aggregate-metadata"
	CacheImage         CreationFlow = "cache-image"
	ValidateImage      CreationFlow = "validate-image"
	Complete           CreationFlow = "complete"
	Incomplete         CreationFlow = "incomplete"
	Unknown            CreationFlow = "unknown"
	Invalid            CreationFlow = "invalid"
)

type Links struct {
	Timestamp int    `json:"timestamp,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Discord   string `json:"discord,omitempty"`
	External  string `json:"external,omitempty"`
	Medium    string `json:"medium,omitempty"`
	Slug      string `json:"slug,omitempty"`
	Telegram  string `json:"telegram,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Wiki      string `json:"wiki,omitempty"`
	Facebook  string `json:"facebook,omitempty"`
}

type Partnership struct {
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}

type Metadata struct {
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Symbol       string        `json:"symbol,omitempty"`
	ProfileImage string        `json:"profileImage,omitempty"`
	BannerImage  string        `json:"bannerImage,omitempty"`
	Links        Links         `json:"links,omitempty"`
	Benefits     []string      `json:"benefits,omitempty"`
	Partnerships []Partnership `json:"partnerships,omitempty"`
	DisplayType  string        `json:"displayType,omitempty"`
}

type AttributeMetadata struct {
	// Number of tokens with this attribute/trait.
	Count int `json:"count,omitempty"`

	// Percentage of tokens with this attribute/trait.
	Percent int `json:"percent,omitempty"`

	// Equals '1 / (percent / 100)'.
	RarityScore int `json:"rarityScore,omitempty"`
}

type Attribute struct {
	// Defines how the attribute is shown on OpenSea.
	DisplayType DisplayType `json:"displayType,omitempty"`

	// Number of NFTs with this attribute/trait.
	Count int `json:"count,omitempty"`

	// Percentage of NFTs with this attribute/trait.
	Percent int `json:"percent,omitempty"`

	// Inner AttributeMetadata values.
	Values map[string]AttributeMetadata `json:"values,omitempty"`
}

type Create struct {
	Step      CreationFlow           `json:"step,omitempty"`
	UpdatedAt int64                  `json:"updatedAt,omitempty"`
	Error     map[string]interface{} `json:"error,omitempty"`
	Progress  int                    `json:"progress,omitempty"`
}

type Export struct {
	Done bool `json:"done,omitempty"`
}

type State struct {
	Version int    `json:"version,omitempty"`
	Create  Create `json:"create,omitempty"`
	Export  Export `json:"export,omitempty"`
}

type Stats struct {
	ContractAddress string `json:"contractAddress,omitempty"`
	AvgPrice        int    `json:"avgPrice,omitempty"`
	SalesVolume     int    `json:"salesVolume,omitempty"`
	OwnerCount      int    `json:"ownerCount,omitempty"`
	TokenCount      int    `json:"tokenCount,omitempty"`
}

type StatsOverTime struct {
	Daily   Stats `json:"daily,omitempty"`
	Weekly  Stats `json:"weekly,omitempty"`
	Monthly Stats `json:"monthly,omitempty"`
}

type NFTCollection struct {
	// Blockchain identifier.
	ChainId string `json:"chainId,omitempty"`

	// Contract address.
	Address string `json:"address,omitempty"`

	// ERC token standard.
	TokenStandard TokenStandard `json:"tokenStandard,omitempty"`

	// Whether the collection is verified.
	HasBlueCheck bool `json:"hasBlueCheck,omitempty"`

	// The address that created the contract.
	Deployer string `json:"deployer,omitempty"`

	// Current owner of the contract.
	Owner string `json:"owner,omitempty"`

	// Number of unique owners.
	NumOwners int `json:"numOwners,omitempty"`

	// Unix timestamp that indicates when NumOwners has been last updated.
	NumOwnersUpdatedAt int `json:"numOwnersUpdatedAt,omitempty"`

	// Unix timestamp that the contract was deployed at (in ms).
	DeployedAt int `json:"deployedAt,omitempty"`

	// Block nr the collection was deployed at.
	DeployedAtBlock int `json:"deployedAtBlock,omitempty"`

	// Metadata conforming OpenSea's metadata standards.
	// See: https://docs.opensea.io/docs/metadata-standards
	Metadata Metadata `json:"metadata,omitempty"`

	// Short slug that is used to refer to this collection in a URL.
	Slug string `json:"slug,omitempty"`

	// Number of available tokens in the collection, excluding burned/destroyed tokens.
	NumNfts int `json:"numNfts,omitempty"`

	// Total number of trait types in the collection.
	NumTraitTypes int `json:"numTraitTypes,omitempty"`

	// The address of the person who initiated the NFT database (index or reindex).
	IndexInitiator string `json:"indexInitiator,omitempty"`

	// Indexer state.
	State State `json:"state,omitempty"`

	// Statistics.
	Stats StatsOverTime `json:"stats,omitempty"`

	// Attributes/traits.
	Attributes map[string]Attribute `json:"attributes,omitempty"`
}
