package wine

// Metadata properties to wine
type Metadata struct {
	Barrel           string `json:"barrel"`
	BrandImage       string `json:"brandImage"`
	Capacity         string `json:"capacity"`
	Cellar           string `json:"cellar"`
	CellarURL        string `json:"cellarURL"`
	Color            string `json:"color"`
	Cork             string `json:"cork"`
	Do               string `json:"do"`
	DoImage          string `json:"doImage"`
	Graduation       string `json:"graduation"`
	Grape            string `json:"grape"`
	PlaceholderImage string `json:"placeholderImage"`
	Path             string `json:"path"`
	Where            string `json:"where"`
}

// Structure of wine type
type Structure struct {
	ID       string   `json:"id"`
	Images   []string `json:"images"`
	Name     string   `json:"name"`
	URL      string   `json:"url"`
	Metadata Metadata `json:"metadata"`
}
