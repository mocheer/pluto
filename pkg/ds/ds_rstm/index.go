package ds_rstm

type RSTM struct {
	CommonBlocks  CommonBlocks
	ProductHeader ProductHeader
	ProductData   ProductData
}

type CommonBlocks struct {
	GenericHeader GenericHeader
}

type ProductHeader struct {
}

type ProductData struct {
}

type GenericHeader struct {
	MagicWord    int32
	MajorVersion int16
	MinorVersion int16
	GenericType  int32
	ProductType  int32
	Reserved     [16]byte
}
