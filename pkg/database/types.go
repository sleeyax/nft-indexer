package database

import "nft-indexer/pkg/indexer/thirdparty/zora"

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
	Unindexed          CreationFlow = "unindexed"
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
	Timestamp int64  `firestore:"timestamp,omitempty"`
	Twitter   string `firestore:"twitter,omitempty"`
	Discord   string `firestore:"discord,omitempty"`
	External  string `firestore:"external,omitempty"`
	Medium    string `firestore:"medium,omitempty"`
	Slug      string `firestore:"slug,omitempty"`
	Telegram  string `firestore:"telegram,omitempty"`
	Instagram string `firestore:"instagram,omitempty"`
	Wiki      string `firestore:"wiki,omitempty"`
	Facebook  string `firestore:"facebook,omitempty"`
}

type Partnership struct {
	Name string `firestore:"name,omitempty"`
	Link string `firestore:"link,omitempty"`
}

type Metadata struct {
	Name         string        `firestore:"name,omitempty"`
	Description  string        `firestore:"description,omitempty"`
	Symbol       string        `firestore:"symbol,omitempty"`
	ProfileImage string        `firestore:"profileImage,omitempty"`
	BannerImage  string        `firestore:"bannerImage,omitempty"`
	Links        Links         `firestore:"links,omitempty"`
	Benefits     []string      `firestore:"benefits,omitempty"`
	Partnerships []Partnership `firestore:"partnerships,omitempty"`
	DisplayType  string        `firestore:"displayType,omitempty"`
}

type AttributeMetadata struct {
	// Number of tokens with this attribute/trait.
	Count int `firestore:"count,omitempty"`

	// Percentage of tokens with this attribute/trait.
	Percent int `firestore:"percent,omitempty"`

	// Equals '1 / (percent / 100)'.
	RarityScore int `firestore:"rarityScore,omitempty"`
}

type Attribute struct {
	// Defines how the attribute is shown on OpenSea.
	DisplayType DisplayType `firestore:"displayType,omitempty"`

	// Number of NFTs with this attribute/trait.
	Count int `firestore:"count,omitempty"`

	// Percentage of NFTs with this attribute/trait.
	Percent int `firestore:"percent,omitempty"`

	// Inner AttributeMetadata values.
	Values map[string]AttributeMetadata `firestore:"values,omitempty"`
}

type Create struct {
	Step      CreationFlow           `firestore:"step,omitempty"`
	UpdatedAt int64                  `firestore:"updatedAt,omitempty"`
	Error     map[string]interface{} `firestore:"error,omitempty"`
	Progress  int                    `firestore:"progress,omitempty"`
}

type Export struct {
	Done bool `firestore:"done,omitempty"`
}

type State struct {
	Version int    `firestore:"version,omitempty"`
	Create  Create `firestore:"create,omitempty"`
	Export  Export `firestore:"export,omitempty"`
}

type Stats struct {
	ContractAddress string `firestore:"contractAddress,omitempty"`
	AvgPrice        int    `firestore:"avgPrice,omitempty"`
	SalesVolume     int    `firestore:"salesVolume,omitempty"`
	OwnerCount      int    `firestore:"ownerCount,omitempty"`
	TokenCount      int    `firestore:"tokenCount,omitempty"`
}

type CollectionStats struct {
	ChainId                   string      `firestore:"chainId,omitempty"`
	CollectionAddress         string      `firestore:"collectionAddress,omitempty"`
	Volume                    float64     `firestore:"volume,omitempty"`
	NumSales                  int         `firestore:"numSales,omitempty"`
	VolumeUSDC                float64     `firestore:"volumeUSDC,omitempty"`
	NumOwners                 int         `firestore:"numOwners,omitempty"`
	NumNfts                   int         `firestore:"numNfts,omitempty"`
	TopOwnersByOwnedNftsCount []zora.Node `firestore:"topOwnersByOwnedNftsCount,omitempty"`
	UpdatedAt                 int64       `firestore:"updatedAt,omitempty"`
}

type StatsOverTime struct {
	Daily   Stats `firestore:"daily,omitempty"`
	Weekly  Stats `firestore:"weekly,omitempty"`
	Monthly Stats `firestore:"monthly,omitempty"`
}

type NFTCollection struct {
	// Blockchain identifier.
	ChainId string `firestore:"chainId,omitempty"`

	// Contract address.
	Address string `firestore:"address,omitempty"`

	// ERC token standard.
	TokenStandard TokenStandard `firestore:"tokenStandard,omitempty"`

	// Whether the collection is verified.
	HasBlueCheck bool `firestore:"hasBlueCheck,omitempty"`

	// The address that created the contract.
	Deployer string `firestore:"deployer,omitempty"`

	// Current owner of the contract.
	Owner string `firestore:"owner,omitempty"`

	// Number of unique owners.
	NumOwners int `firestore:"numOwners,omitempty"`

	// Unix timestamp that indicates when NumOwners has been last updated.
	NumOwnersUpdatedAt int `firestore:"numOwnersUpdatedAt,omitempty"`

	// Unix timestamp that the contract was deployed at (in ms).
	DeployedAt int `firestore:"deployedAt,omitempty"`

	// Block nr the collection was deployed at.
	DeployedAtBlock int `firestore:"deployedAtBlock,omitempty"`

	// Metadata conforming OpenSea's metadata standards.
	// See: https://docs.opensea.io/docs/metadata-standards
	Metadata Metadata `firestore:"metadata,omitempty"`

	// Short slug that is used to refer to this collection in a URL.
	Slug string `firestore:"slug,omitempty"`

	// Number of available tokens in the collection, excluding burned/destroyed tokens.
	NumNfts int `firestore:"numNfts,omitempty"`

	// Total number of trait types in the collection.
	NumTraitTypes int `firestore:"numTraitTypes,omitempty"`

	// The address of the person who initiated the NFT database (index or reindex).
	IndexInitiator string `firestore:"indexInitiator,omitempty"`

	// Indexer state.
	State State `firestore:"state,omitempty"`

	// Statistics.
	Stats StatsOverTime `firestore:"stats,omitempty"`

	// Attributes/traits.
	Attributes map[string]Attribute `firestore:"attributes,omitempty"`
}
