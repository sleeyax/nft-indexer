package opensea

type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

type PrimaryAsset struct {
	Address                     string      `json:"address,omitempty"`
	AssetContractType           string      `json:"asset_contract_type,omitempty"`
	CreatedDate                 string      `json:"created_date,omitempty"`
	Name                        string      `json:"name,omitempty"`
	NftVersion                  string      `json:"nft_version,omitempty"`
	OpenseaVersion              interface{} `json:"opensea_version,omitempty"`
	Owner                       int         `json:"owner,omitempty"`
	SchemaName                  string      `json:"schema_name,omitempty"`
	Symbol                      string      `json:"symbol,omitempty"`
	TotalSupply                 string      `json:"total_supply,omitempty"` // not accurate
	Description                 string      `json:"description,omitempty"`
	ExternalLink                string      `json:"external_link,omitempty"`
	ImageUrl                    string      `json:"image_url,omitempty"`
	DefaultToFiat               bool        `json:"default_to_fiat,omitempty"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points,omitempty"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points,omitempty"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers,omitempty"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points,omitempty"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points,omitempty"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points,omitempty"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points,omitempty"`
	PayoutAddress               string      `json:"payout_address,omitempty"`
}

type Collection struct {
	BannerImageUrl              string         `json:"banner_image_url,omitempty"`
	ChatUrl                     string         `json:"chat_url,omitempty"`
	CreatedDate                 string         `json:"created_date,omitempty"`
	DefaultToFiat               bool           `json:"default_to_fiat,omitempty"`
	Description                 string         `json:"description,omitempty"`
	DevBuyerFeeBasisPoints      string         `json:"dev_buyer_fee_basis_points,omitempty"`
	DevSellerFeeBasisPoints     string         `json:"dev_seller_fee_basis_points,omitempty"`
	DiscordUrl                  string         `json:"discord_url,omitempty"`
	DisplayData                 DisplayData    `json:"display_data"`
	ExternalUrl                 string         `json:"external_url,omitempty"`
	Featured                    bool           `json:"featured,omitempty"`
	FeaturedImageUrl            string         `json:"featured_image_url,omitempty"`
	Hidden                      bool           `json:"hidden,omitempty"`
	SafelistRequestStatus       string         `json:"safelist_request_status,omitempty"`
	ImageUrl                    string         `json:"image_url,omitempty"`
	IsSubjectToWhitelist        bool           `json:"is_subject_to_whitelist,omitempty"`
	LargeImageUrl               string         `json:"large_image_url,omitempty"`
	MediumUsername              string         `json:"medium_username,omitempty"`
	Name                        string         `json:"name,omitempty"`
	OnlyProxiedTransfers        bool           `json:"only_proxied_transfers,omitempty"`
	OpenseaBuyerFeeBasisPoints  string         `json:"opensea_buyer_fee_basis_points,omitempty"`
	OpenseaSellerFeeBasisPoints string         `json:"opensea_seller_fee_basis_points,omitempty"`
	PayoutAddress               string         `json:"payout_address,omitempty"`
	RequireEmail                bool           `json:"require_email,omitempty"`
	ShortDescription            string         `json:"short_description,omitempty"`
	Slug                        string         `json:"slug,omitempty"`
	TelegramUrl                 string         `json:"telegram_url,omitempty"`
	TwitterUsername             string         `json:"twitter_username,omitempty"`
	InstagramUsername           string         `json:"instagram_username,omitempty"`
	WikiUrl                     string         `json:"wiki_url,omitempty"`
	PrimaryAssetContracts       []PrimaryAsset `json:"primary_asset_contracts,omitempty"`
}

type ContractResponse struct {
	Collection                  Collection  `json:"collection"`
	Address                     string      `json:"address,omitempty"`
	ContractType                string      `json:"contract_type,omitempty"`
	CreatedDate                 string      `json:"created_date,omitempty"`
	Name                        string      `json:"name,omitempty"`
	NftVersion                  string      `json:"nft_version,omitempty"`
	OpenseaVersion              interface{} `json:"opensea_version,omitempty"`
	Owner                       int         `json:"owner,omitempty"`
	SchemaName                  string      `json:"schema_name,omitempty"`
	Symbol                      string      `json:"symbol,omitempty"`
	TotalSupply                 string      `json:"total_supply,omitempty"`
	Description                 string      `json:"description,omitempty"`
	ExternalLink                string      `json:"external_link,omitempty"`
	ImageUrl                    string      `json:"image_url,omitempty"`
	DefaultToFiat               bool        `json:"default_to_fiat,omitempty"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points,omitempty"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points,omitempty"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers,omitempty"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points,omitempty"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points,omitempty"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points,omitempty"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points,omitempty"`
	PayoutAddress               interface{} `json:"payout_address,omitempty"`
}
