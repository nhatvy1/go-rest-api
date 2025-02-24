package user_models

type Paging struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

func (p *Paging) Process() {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}
}
