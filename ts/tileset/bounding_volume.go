package tileset

import "errors"

type BoundingVolume struct {
	Region *[]float64 `json:"region,omitempty"`
	Box    *[]float64 `json:"box,omitempty"`
	Sphere *[]float64 `json:"sphere,omitempty"`
}

func (b *BoundingVolume) SetBox(box []float64) error {
	if len(box) != 12 {
		return errors.New("box must 12 element")
	}
	if b.Region != nil || b.Sphere != nil {
		b.Region = nil
		b.Sphere = nil
	}
	b.Box = &box
	return nil
}

func (b *BoundingVolume) SetRegion(region []float64) error {
	if len(region) != 6 {
		return errors.New("region must 6 element")
	}
	if b.Box != nil || b.Sphere != nil {
		b.Box = nil
		b.Sphere = nil
	}
	b.Region = &region
	return nil
}

func (b *BoundingVolume) SetSphere(sphere []float64) error {
	if len(sphere) != 4 {
		return errors.New("sphere must 4 element")
	}
	if b.Box != nil || b.Region != nil {
		b.Box = nil
		b.Region = nil
	}
	b.Sphere = &sphere
	return nil
}

func (b *BoundingVolume) GetRegion() []float64 {
	return *b.Region
}

func (b *BoundingVolume) GetBox() []float64 {
	return *b.Box
}

func (b *BoundingVolume) GetSphere() []float64 {
	return *b.Sphere
}

func (b *BoundingVolume) GetData() []float64 {
	if b.Region != nil {
		return *b.Region
	}
	if b.Box != nil {
		return *b.Box
	}
	if b.Sphere != nil {
		return *b.Sphere
	}
	return nil
}
