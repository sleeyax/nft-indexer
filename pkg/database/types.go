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

type CreationStep string

const (
	UnindexedStep          CreationStep = "unindexed"
	CollectionCreatorStep  CreationStep = "collection-creator"
	CollectionMetadataStep CreationStep = "collection-metadata"
	CollectionMintsStep    CreationStep = "collection-mints"
	TokenMetadataStep      CreationStep = "token-metadata"
	TokenMetadataUriStep   CreationStep = "token-metadata-uri"
	AggregateMetadataStep  CreationStep = "aggregate-metadata"
	CacheImageStep         CreationStep = "cache-image"
	ValidateImageStep      CreationStep = "validate-image"
	CompleteStep           CreationStep = "complete"
	IncompleteStep         CreationStep = "incomplete"
	UnknownStep            CreationStep = "unknown"
	InvalidStep            CreationStep = "invalid"
)

type Links struct {
	Timestamp int64  `firestore:"timestamp,omitempty" json:"timestamp,omitempty"`
	Twitter   string `firestore:"twitter,omitempty" json:"twitter,omitempty"`
	Discord   string `firestore:"discord,omitempty" json:"discord,omitempty"`
	External  string `firestore:"external,omitempty" json:"external,omitempty"`
	Medium    string `firestore:"medium,omitempty" json:"medium,omitempty"`
	Slug      string `firestore:"slug,omitempty" json:"slug,omitempty"`
	Telegram  string `firestore:"telegram,omitempty" json:"telegram,omitempty"`
	Instagram string `firestore:"instagram,omitempty" json:"instagram,omitempty"`
	Wiki      string `firestore:"wiki,omitempty" json:"wiki,omitempty"`
	Facebook  string `firestore:"facebook,omitempty" json:"facebook,omitempty"`
}

type Partnership struct {
	Name string `firestore:"name,omitempty" json:"name,omitempty"`
	Link string `firestore:"link,omitempty" json:"link,omitempty"`
}

type Metadata struct {
	Name         string        `firestore:"name,omitempty" json:"name,omitempty"`
	Description  string        `firestore:"description,omitempty" json:"description,omitempty"`
	Symbol       string        `firestore:"symbol,omitempty" json:"symbol,omitempty"`
	ProfileImage string        `firestore:"profileImage,omitempty" json:"profileImage,omitempty"`
	BannerImage  string        `firestore:"bannerImage,omitempty" json:"bannerImage,omitempty"`
	Links        Links         `firestore:"links,omitempty" json:"links"`
	Benefits     []string      `firestore:"benefits,omitempty" json:"benefits,omitempty"`
	Partnerships []Partnership `firestore:"partnerships,omitempty" json:"partnerships,omitempty"`
	DisplayType  string        `firestore:"displayType,omitempty" json:"displayType,omitempty"`
}

type AttributeMetadata struct {
	// Number of tokens with this attribute/trait.
	Count int `firestore:"count,omitempty" json:"count,omitempty"`

	// Percentage of tokens with this attribute/trait.
	Percent int `firestore:"percent,omitempty" json:"percent,omitempty"`

	// Equals '1 / (percent / 100)'.
	RarityScore int `firestore:"rarityScore,omitempty" json:"rarityScore,omitempty"`
}

type Attribute struct {
	// Defines how the attribute is shown on OpenSea.
	DisplayType DisplayType `firestore:"displayType,omitempty" json:"displayType,omitempty"`

	// Number of NFTs with this attribute/trait.
	Count int `firestore:"count,omitempty" json:"count,omitempty"`

	// Percentage of NFTs with this attribute/trait.
	Percent int `firestore:"percent,omitempty" json:"percent,omitempty"`

	// Inner AttributeMetadata values.
	Values map[string]AttributeMetadata `firestore:"values,omitempty" json:"values,omitempty"`
}

type Create struct {
	Step            CreationStep           `firestore:"step,omitempty" json:"step,omitempty"`
	UpdatedAt       int64                  `firestore:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Error           map[string]interface{} `firestore:"error,omitempty" json:"error,omitempty"`
	Progress        int                    `firestore:"progress,omitempty" json:"progress,omitempty"`
	ReservoirCursor string                 `firestore:"reservoirCursor,omitempty" json:"reservoirCursor"`
	ZoraCursor      string                 `firestore:"zoraCursor,omitempty" json:"zoraCursor"`
}

type Export struct {
	Done bool `firestore:"done,omitempty"`
}

type State struct {
	Version int    `firestore:"version,omitempty" json:"version,omitempty"`
	Create  Create `firestore:"create,omitempty" json:"create"`
	Export  Export `firestore:"export,omitempty" json:"export"`
}

type Stats struct {
	ContractAddress string `firestore:"contractAddress,omitempty" json:"contractAddress,omitempty"`
	AvgPrice        int    `firestore:"avgPrice,omitempty" json:"avgPrice,omitempty"`
	SalesVolume     int    `firestore:"salesVolume,omitempty" json:"salesVolume,omitempty"`
	OwnerCount      int    `firestore:"ownerCount,omitempty" json:"ownerCount,omitempty"`
	TokenCount      int    `firestore:"tokenCount,omitempty" json:"tokenCount,omitempty"`
}

type ZoraNFTStats struct {
	ChainId                   string      `firestore:"chainId,omitempty" json:"chainId,omitempty"`
	CollectionAddress         string      `firestore:"collectionAddress,omitempty" json:"collectionAddress,omitempty"`
	Volume                    float64     `firestore:"volume,omitempty" json:"volume,omitempty"`
	NumSales                  int         `firestore:"numSales,omitempty" json:"numSales,omitempty"`
	VolumeUSDC                float64     `firestore:"volumeUSDC,omitempty" json:"volumeUSDC,omitempty"`
	NumOwners                 int         `firestore:"numOwners,omitempty" json:"numOwners,omitempty"`
	NumNfts                   int         `firestore:"numNfts,omitempty" json:"numNfts,omitempty"`
	TopOwnersByOwnedNftsCount []zora.Node `firestore:"topOwnersByOwnedNftsCount,omitempty" json:"topOwnersByOwnedNftsCount,omitempty"`
	UpdatedAt                 int64       `firestore:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type StatsOverTime struct {
	Daily   Stats `firestore:"daily,omitempty" json:"daily"`
	Weekly  Stats `firestore:"weekly,omitempty" json:"weekly"`
	Monthly Stats `firestore:"monthly,omitempty" json:"monthly"`
}

type NFTCollection struct {
	// Blockchain identifier.
	ChainId string `firestore:"chainId,omitempty" json:"chainId,omitempty"`

	// Contract address.
	Address string `firestore:"address,omitempty" json:"address,omitempty"`

	// ERC token standard.
	TokenStandard TokenStandard `firestore:"tokenStandard,omitempty" json:"tokenStandard,omitempty"`

	// Whether the collection is verified.
	HasBlueCheck bool `firestore:"hasBlueCheck,omitempty" json:"hasBlueCheck,omitempty"`

	// The address that created the contract.
	Deployer string `firestore:"deployer,omitempty" json:"deployer,omitempty"`

	// Current owner of the contract.
	Owner string `firestore:"owner,omitempty" json:"owner,omitempty"`

	// Number of unique owners.
	NumOwners int `firestore:"numOwners,omitempty" json:"numOwners,omitempty"`

	// Unix timestamp that indicates when NumOwners has been last updated.
	NumOwnersUpdatedAt int `firestore:"numOwnersUpdatedAt,omitempty" json:"numOwnersUpdatedAt,omitempty"`

	// Unix timestamp that the contract was deployed at (in ms).
	DeployedAt int `firestore:"deployedAt,omitempty" json:"deployedAt,omitempty"`

	// Block nr the collection was deployed at.
	DeployedAtBlock int `firestore:"deployedAtBlock,omitempty" json:"deployedAtBlock,omitempty"`

	// Metadata conforming OpenSea's metadata standards.
	// See: https://docs.opensea.io/docs/metadata-standards
	Metadata Metadata `firestore:"metadata,omitempty" json:"metadata"`

	// Short slug that is used to refer to this collection in a URL.
	Slug string `firestore:"slug,omitempty" json:"slug,omitempty"`

	// Number of available tokens in the collection, excluding burned/destroyed tokens.
	NumNfts int `firestore:"numNfts,omitempty" json:"numNfts,omitempty"`

	// Total number of trait types in the collection.
	NumTraitTypes int `firestore:"numTraitTypes,omitempty" json:"numTraitTypes,omitempty"`

	// The address of the person who initiated the NFT database (index or reindex).
	IndexInitiator string `firestore:"indexInitiator,omitempty" json:"indexInitiator,omitempty"`

	// Indexer state.
	State State `firestore:"state,omitempty" json:"state"`

	// Statistics.
	Stats StatsOverTime `firestore:"stats,omitempty" json:"stats"`

	// Attributes/traits.
	Attributes map[string]Attribute `firestore:"attributes,omitempty" json:"attributes,omitempty"`

	// NFT collection statistics received from Zora.
	// Note that this information isn't stored in the database.
	ZoraStats *ZoraNFTStats `firestore:"-" json:"zoraStats,omitempty"`

	// List of NFTs in this collection.
	Tokens []Erc721Token `firestore:"-" json:"tokens"`
}

type Erc721Token struct {
	Slug                 string              `firestore:"slug,omitempty" json:"slug,omitempty"`
	TokenId              string              `firestore:"tokenId,omitempty" json:"tokenId,omitempty"`
	TokenIdNumeric       int                 `firestore:"tokenIdNumeric,omitempty" json:"tokenIdNumeric,omitempty"`
	ChainId              string              `firestore:"chainId,omitempty" json:"chainId,omitempty"`
	CollectionAddress    string              `firestore:"collectionAddress,omitempty" json:"collectionAddress,omitempty"`
	NumTraitTypes        int                 `firestore:"numTraitTypes,omitempty" json:"numTraitTypes,omitempty"`
	Metadata             Erc721TokenMetadata `firestore:"metadata,omitempty" json:"metadata"`
	UpdatedAt            int64               `firestore:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	Owner                string              `firestore:"owner,omitempty" json:"owner,omitempty"`
	TokenStandard        TokenStandard       `firestore:"tokenStandard,omitempty" json:"tokenStandard,omitempty"`
	Image                Erc721TokenImage    `firestore:"image,omitempty" json:"image"`
	MintedAt             int64               `firestore:"mintedAt,omitempty" json:"mintedAt,omitempty"`
	Minter               string              `firestore:"minter,omitempty" json:"minter,omitempty"`
	MintTxHash           string              `firestore:"mintTxHash,omitempty" json:"mintTxHash,omitempty"`
	MintPrice            float64             `firestore:"mintPrice,omitempty" json:"mintPrice,omitempty"`
	MintCurrencyAddress  string              `firestore:"mintCurrencyAddress,omitempty" json:"mintCurrencyAddress,omitempty"`
	MintCurrencyDecimals int                 `firestore:"mintCurrencyDecimals,omitempty" json:"mintCurrencyDecimals,omitempty"`
	MintCurrencyName     string              `firestore:"mintCurrencyName,omitempty" json:"mintCurrencyName,omitempty"`
}

type Erc721TokenImage struct {
	Url       string `firestore:"url,omitempty" json:"url,omitempty"`
	UpdatedAt int64  `firestore:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type Erc721TokenAttribute struct {
	Value     string `firestore:"value,omitempty" json:"value,omitempty"`
	TraitType string `firestore:"traitType,omitempty" json:"trait_type,omitempty"`
}

type Erc721TokenMetadata struct {
	Name            string                 `firestore:"name,omitempty" json:"name,omitempty"`
	Title           string                 `firestore:"title,omitempty" json:"title,omitempty"`
	Image           string                 `firestore:"image,omitempty" json:"image,omitempty"`
	ImageData       string                 `firestore:"imageData,omitempty" json:"image_data,omitempty"`
	ExternalUrl     string                 `firestore:"externalUrl,omitempty" json:"external_url,omitempty"`
	Description     string                 `firestore:"description,omitempty" json:"description,omitempty"`
	BackgroundColor string                 `firestore:"backgroundColor,omitempty" json:"background_color,omitempty"`
	AnimationUrl    string                 `firestore:"animationUrl,omitempty" json:"animation_url,omitempty"`
	YoutubeUrl      string                 `firestore:"youtubeUrl,omitempty" json:"youtube_url,omitempty"`
	Attributes      []Erc721TokenAttribute `firestore:"attributes,omitempty" json:"attributes,omitempty"`
}
