package entity

type SiteSource string

const (
	BookDepository SiteSource = "BookDepository"
	RightStufAnime SiteSource = "RightStufAnime"
)

var (
	ListAllSource = []SiteSource{
		BookDepository,
		RightStufAnime,
	}
)

func (source SiteSource) Validate() bool {
	for _, site := range ListAllSource {
		if source == site {
			return true
		}
	}
	return false
}

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

type ProductBookDepository struct {
	Title   string
	Price   float64
	InStock bool
}
