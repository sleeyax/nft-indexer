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
	Timestamp int
	Twitter   string
	Discord   string
	External  string
	Medium    string
	Slug      string
	Telegram  string
	Instagram string
	Wiki      string
	Facebook  string
}

type Partnership struct {
	Name string
	Link string
}

type Metadata struct {
	Name         string
	Description  string
	Symbol       string
	ProfileImage string
	BannerImage  string
	Links        Links
	Benefits     []string
	Partnerships []Partnership
	DisplayType  string
}

type AttributeMetadata struct {
	// Number of tokens with this attribute/trait.
	Count int

	// Percentage of tokens with this attribute/trait.
	Percent int

	// Equals '1 / (percent / 100)'.
	RarityScore int
}

type Attribute struct {
	// Defines how the attribute is shown on OpenSea.
	DisplayType DisplayType

	// Number of NFTs with this attribute/trait.
	Count int

	// Percentage of NFTs with this attribute/trait.
	Percent int

	// Inner AttributeMetadata values.
	Values map[string]AttributeMetadata
}

type Create struct {
	Step      CreationFlow
	UpdatedAt int64
	Error     map[string]interface{}
	Progress  int
}

type Export struct {
	Done bool
}

type State struct {
	Version int
	Create  Create
	Export  Export
}

type Stats struct {
	ContractAddress string
	AvgPrice        int
	SalesVolume     int
	OwnerCount      int
	TokenCount      int
}

type StatsOverTime struct {
	Daily   Stats
	Weekly  Stats
	Monthly Stats
}

type NFTCollection struct {
	// Blockchain identifier.
	ChainId string

	// Contract address.
	Address string

	// ERC token standard.
	TokenStandard TokenStandard

	// Whether the collection is verified.
	HasBlueCheck bool

	// The address that created the contract.
	Deployer string

	// Current owner of the contract.
	Owner string

	// Number of unique owners.
	NumOwners int

	// Unix timestamp that indicates when NumOwners has been last updated.
	NumOwnersUpdatedAt int

	// Unix timestamp that the contract was deployed at (in ms).
	DeployedAt int

	// Block nr the collection was deployed at.
	DeployedAtBlock int

	// Metadata conforming OpenSea's metadata standards.
	// See: https://docs.opensea.io/docs/metadata-standards
	Metadata Metadata

	// Short slug that is used to refer to this collection in a URL.
	Slug string

	// Number of available tokens in the collection, excluding burned/destroyed tokens.
	NumNfts int

	// Total number of trait types in the collection.
	NumTraitTypes int

	// The address of the person who initiated the NFT database (index or reindex).
	IndexInitiator string

	// Indexer state.
	State State

	// Statistics.
	Stats StatsOverTime

	// Attributes/traits.
	Attributes map[string]Attribute
}
