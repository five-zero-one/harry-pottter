package harrypotter

type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Blood    string `json:"blood"`
	Species  string `json:"species"`
	Patronus string `json:"patronus"`
	Born     string `json:"born"`
	Quote    string `json:"quote"`
	ImgURL   string `json:"imgUrl"`
}
