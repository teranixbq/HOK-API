package response

type ResponseHero struct {
	Id        string   `json:"id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	Role      string `json:"role"`
	Type_Hero string `json:"type_hero"`
	Skill     string `json:"skill"`
}
