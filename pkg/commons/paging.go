package commons

type PagingUser struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (p *PagingUser) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}
