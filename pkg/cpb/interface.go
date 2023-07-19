package cpb

type Marshaler interface {
	MarshalCPB() (any, uint64)
}
