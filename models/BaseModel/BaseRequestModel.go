package model

type BaseSearchOption struct {
	PageIndex int
	PageSize  int
	IsPage    bool
	SortName  string
	IsDesc    bool
}
