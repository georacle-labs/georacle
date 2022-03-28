package query

import (
	"testing"

	"github.com/golang/geo/r2"
)

func TestAreaCreate(t *testing.T) {
	_, err := NewAreaQuery("amenity", "cafe", "London", 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBBoxCreate(t *testing.T) {
	sw := r2.Point{Y: 40.771900, X: -73.974600}
	ne := r2.Point{Y: 40.797500, X: -73.946900}
	box := r2.RectFromPoints(sw, ne)

	_, err := NewBBoxQuery("public_transport", "station", box, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProximityCreate(t *testing.T) {
	center := r2.Point{Y: 48.858093, X: 2.294694}
	rad := 500.0

	_, err := NewProximityQuery("amenity", "restaurant", center, rad, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGeocodeCreate(t *testing.T) {
	_, err := NewGeocodeQuery("221B Baker St, London NW1 6XE, UK")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReverseGeocodeQuery(t *testing.T) {
	p := r2.Point{Y: 51.5233879, X: -0.1582367}

	_, err := NewReverseGeocodeQuery(p)
	if err != nil {
		t.Fatal(err)
	}
}
