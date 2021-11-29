package types

type DateCategory string

const (
	UnknownDateCategory DateCategory = "unknown"
	Day1                DateCategory = "day1"
	Day7                DateCategory = "day7"
	Day30               DateCategory = "day30"
)

func (s DateCategory) String() string {
	return string(s)
}

func ParseDateCategory(s string) DateCategory {
	switch s {
	case string(Day1):
		return Day1
	case string(Day7):
		return Day7
	case string(Day30):
		return Day30
	default:
		return UnknownDateCategory
	}
}
