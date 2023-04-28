package gcpcomputepricing

type Subtype struct {
	Description string            `json:"description"`
	Regions     map[string]Region `json:"regions"`
}

type Region struct {
	Prices       []Price `json:"price"`
	Name         string  `json:"name"`
	InternalName string  `json:"internal_name"`
}

type Price struct {
	Val      uint64 `json:"val"`
	Currency string `json:"currency"`
	Nanos    uint64 `json:"nanos"`
}
