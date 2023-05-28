package rctx

type Permission struct {
	NeedToken  bool
	NeedCasbin bool
}

func NewPermission() *Permission {
	return &Permission{NeedToken: true, NeedCasbin: true}
}
func (p *Permission) WithToken(needToken bool) *Permission {
	p.NeedToken = needToken
	return p
}

func (p *Permission) WithCasbin(needCasbin bool) *Permission {
	p.NeedCasbin = needCasbin
	return p
}
