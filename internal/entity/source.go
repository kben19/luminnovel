package entity

type SiteSource string

const (
	BookDepository SiteSource = "BookDepository"
	RightStufAnime SiteSource = "RightStufAnime"
)

type Source struct {
	Title  string
	Volume string
	Site   string
	Domain string
	Path   string
	Params map[string]string
	Header map[string]string
	Body   string
}

type ProductRightStufAnime struct {
	Title    string
	Price    float64
	PriceFmt string
	InStock  bool
	PreOrder bool
}
