package shared

// Page is normalized pagination input. Defaults and clamping happen at the
// HTTP edge, so use cases receive values they can trust.
type Page struct {
	Number int
	Size   int
}

func (p Page) Limit() int  { return p.Size }
func (p Page) Offset() int { return (p.Number - 1) * p.Size }

type SortDirection string

const (
	Asc  SortDirection = "ASC"
	Desc SortDirection = "DESC"
)

func (d SortDirection) IsValid() bool {
	return d == Asc || d == Desc
}
