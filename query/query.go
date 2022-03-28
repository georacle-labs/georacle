package query

import (
	"encoding/json"

	g "github.com/georacle-labs/go-osm/geometry"
	n "github.com/georacle-labs/go-osm/nominatim"
	o "github.com/georacle-labs/go-osm/overpass"
	"github.com/golang/geo/r2"
	"golang.org/x/crypto/sha3"
)

// Type is a flag that denotes a query type
type Type uint64

const (
	// AreaQuery denotes a named area query type
	AreaQuery Type = 1

	// BBoxQuery denotes a bounding box query type
	BBoxQuery = (AreaQuery << 1)

	// ProximityQuery denotes a proximity query type
	ProximityQuery = (BBoxQuery << 1)

	// GeocodeQuery denotes a geocode query type
	GeocodeQuery = (ProximityQuery << 1)

	// ReverseGeocodeQuery denotes a reverse geocode query type
	ReverseGeocodeQuery = (GeocodeQuery << 1)
)

// Query represents a generic map query
type Query struct {
	Type   Type
	Query  interface{}
	Digest []byte
}

// MakeQuery builds and hashes a generic query
func MakeQuery(queryType Type, query interface{}) (*Query, error) {
	payload, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	h := sha3.NewLegacyKeccak256()
	sum := h.Sum(payload)
	return &Query{Type: queryType, Query: query, Digest: sum}, nil
}

// NewAreaQuery constructs a named area query
func NewAreaQuery(key, val, name string, limit uint) (*Query, error) {
	tags := o.Tags{Key: key, Value: val, Limit: limit}
	area := o.Area{Tags: tags, Name: name}

	return MakeQuery(AreaQuery, area)
}

// NewBBoxQuery constructs a bounding box query
func NewBBoxQuery(key, val string, box r2.Rect, limit uint) (*Query, error) {
	tags := o.Tags{Key: key, Value: val, Limit: limit}

	sw := box.Lo()
	ne := box.Hi()
	bbox := o.BBox{Tags: tags, South: sw.Y, West: sw.X, North: ne.Y, East: ne.X}

	return MakeQuery(BBoxQuery, bbox)
}

// NewProximityQuery constructs a proximity query
func NewProximityQuery(key, val string, center r2.Point, radius float64, limit uint) (*Query, error) {
	tags := o.Tags{Key: key, Value: val, Limit: limit}
	refPoint := g.Point{Lat: center.Y, Lon: center.X}
	proximity := o.Proximity{Tags: tags, Point: refPoint, Radius: radius}

	return MakeQuery(ProximityQuery, proximity)
}

// NewGeocodeQuery constructs a geocode query
func NewGeocodeQuery(address string) (*Query, error) {
	geocode := n.GeocodeQuery{Address: address}

	return MakeQuery(GeocodeQuery, geocode)
}

// NewReverseGeocodeQuery constructs a reverse geocode query
func NewReverseGeocodeQuery(point r2.Point) (*Query, error) {
	pt := g.Point{Lat: point.Y, Lon: point.X}
	reverse := n.ReverseGeocodeQuery{Point: pt}

	return MakeQuery(ReverseGeocodeQuery, reverse)
}
