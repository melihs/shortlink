package v1

func (p *Payments) Len() int {
	return len(p.GetList())
}

func (p *Payments) Less(i, j int) bool {
	return p.GetList()[i].GetName() < p.GetList()[j].GetName()
}

func (p *Payments) Swap(i, j int) {
	p.List[i], p.List[j] = p.GetList()[j], p.GetList()[i]
}
