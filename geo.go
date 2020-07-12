package mobile

import (
	"github.com/gopub/types"
)

type Point types.Point

func NewPoint() *Point {
	return new(Point)
}

type Place types.Place

func NewPlace() *Place {
	return new(Place)
}

func (p *Place) SetCoordinate(c *Point) {
	p.Coordinate = (*types.Point)(c)
}

func (p *Place) GetCoordinate() *Point {
	return (*Point)(p.Coordinate)
}

type Country types.Country

type CountryList struct {
	List []*Country
}

func NewCountryList() *CountryList {
	return &CountryList{}
}

func (l *CountryList) Size() int {
	return len(l.List)
}

func (l *CountryList) Get(i int) *Country {
	return l.List[i]
}

func GetCountryList() *CountryList {
	l := GetCountries()
	ll := NewCountryList()
	ll.List = make([]*Country, len(l))
	for i, c := range l {
		ll.List[i] = (*Country)(c)
	}
	return ll
}
