package entities

// FundBreakdown struct
type FundBreakdown struct {
	Ticker     string             `json:"ticker,omitempty"`
	AssetClass string             `json:"assetClass,omitempty"`
	Sectors    []*SectorBreakdown `json:"sectors,omitempty"`
}

// SectorBreakdown struct
type SectorBreakdown struct {
	SectorCode  string  `json:"sectorCode,omitempty"`
	SectorName  string  `json:"sectorName,omitempty"`
	FundPercent float64 `json:"fundPercent,omitempty"`
}
