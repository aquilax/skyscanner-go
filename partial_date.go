package skyscanner

// PartialDate is Skyscaner API partial date
type PartialDate string

func NewPartialDate(pds string) *PartialDate {
	pd := PartialDate(pds)
	return &pd
}

func (pd *PartialDate) String() string {
	return string(*pd)
}
