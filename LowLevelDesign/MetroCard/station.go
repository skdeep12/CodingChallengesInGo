package MetroCard

type Station struct {
	Name string
}

func NewStation(name string) Station {
	return Station{
		Name: name,
	}
}

func (s Station) Equals(s2 Station) bool {
	if s.Name == s2.Name {
		return true
	}
	return false
}
