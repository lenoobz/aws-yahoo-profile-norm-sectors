package consts

// Collection names
const (
	YAHOO_ASSET_PROFILES_COLLECTION = "yahoo_asset_profiles" // Should match with Colnames's key of AppConf
	ASSET_SECTORS_COLLECTION        = "asset_sectors"
)

const (
	BOND     = "BOND"
	EQUITY   = "EQUITY"
	BALANCED = "BALANCED"
)

const (
	DATA_SOURCE = "YAHOO_FINANCE"
)

var Sectors = []struct {
	Name string
	Code string
}{
	{
		Name: "Industrials",
		Code: "IND",
	},
	{
		Name: "Producer Durables",
		Code: "IND",
	},
	{
		Name: "Financials",
		Code: "FIN",
	},
	{
		Name: "Financial Services",
		Code: "FIN",
	},
	{
		Name: "Health Care",
		Code: "HEC",
	},
	{
		Name: "Healthcare",
		Code: "HEC",
	},
	{
		Name: "Consumer Discretionary",
		Code: "CND",
	},
	{
		Name: "Consumer Staples",
		Code: "CNS",
	},
	{
		Name: "Consumer Goods",
		Code: "CNS",
	},
	{
		Name: "Consumer Cyclical",
		Code: "CNS",
	},
	{
		Name: "Consumer Defensive",
		Code: "CNS",
	},
	{
		Name: "Information Technology",
		Code: "TEC",
	},
	{
		Name: "Technology",
		Code: "TEC",
	},
	{
		Name: "Materials",
		Code: "MTL",
	},
	{
		Name: "Basic Materials",
		Code: "MTL",
	},
	{
		Name: "Materials & Processing",
		Code: "MTL",
	},
	{
		Name: "Energy",
		Code: "ENG",
	},
	{
		Name: "Oil & Gas",
		Code: "ENG",
	},
	{
		Name: "Utilities",
		Code: "UTL",
	},
	{
		Name: "Communication Services",
		Code: "TEL",
	},
	{
		Name: "Telecommunications",
		Code: "TEL",
	},
	{
		Name: "Real Estate",
		Code: "RLE",
	},
	{
		Name: "Other",
		Code: "OTH",
	},
	{
		Name: "Consumer Services",
		Code: "OTH",
	},
}
