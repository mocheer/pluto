package ds_tif

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const (
	tNewSubfileType            = 254
	tImageWidth                = 256
	tImageLength               = 257
	tBitsPerSample             = 258
	tCompression               = 259
	tPhotometricInterpretation = 262
	tFillOrder                 = 266
	tDocumentName              = 269
	tPlanarConfiguration       = 284

	tStripOffsets    = 273
	tOrientation     = 274
	tSamplesPerPixel = 277
	tRowsPerStrip    = 278
	tStripByteCounts = 279

	tTileWidth      = 322
	tTileLength     = 323
	tTileOffsets    = 324
	tTileByteCounts = 325

	tXResolution    = 282
	tYResolution    = 283
	tResolutionUnit = 296

	tSoftware     = 305
	tPredictor    = 317
	tColorMap     = 320
	tExtraSamples = 338
	tSampleFormat = 339

	tGDAL_METADATA = 42112
	tGDAL_NODATA   = 42113

	tModelPixelScaleTag     = 33550
	tModelTransformationTag = 34264
	tModelTiepointTag       = 33922
	tGeoKeyDirectoryTag     = 34735
	tGeoDoubleParamsTag     = 34736
	tGeoAsciiParamsTag      = 34737
	tIntergraphMatrixTag    = 33920

	tGTModelTypeGeoKey              = 1024
	tGTRasterTypeGeoKey             = 1025
	tGTCitationGeoKey               = 1026
	tGeographicTypeGeoKey           = 2048
	tGeogCitationGeoKey             = 2049
	tGeogGeodeticDatumGeoKey        = 2050
	tGeogPrimeMeridianGeoKey        = 2051
	tGeogLinearUnitsGeoKey          = 2052
	tGeogLinearUnitSizeGeoKey       = 2053
	tGeogAngularUnitsGeoKey         = 2054
	tGeogAngularUnitSizeGeoKey      = 2055
	tGeogEllipsoidGeoKey            = 2056
	tGeogSemiMajorAxisGeoKey        = 2057
	tGeogSemiMinorAxisGeoKey        = 2058
	tGeogInvFlatteningGeoKey        = 2059
	tGeogAzimuthUnitsGeoKey         = 2060
	tGeogPrimeMeridianLongGeoKey    = 2061
	tProjectedCSTypeGeoKey          = 3072
	tPCSCitationGeoKey              = 3073
	tProjectionGeoKey               = 3074
	tProjCoordTransGeoKey           = 3075
	tProjLinearUnitsGeoKey          = 3076
	tProjLinearUnitSizeGeoKey       = 3077
	tProjStdParallel1GeoKey         = 3078
	tProjStdParallel2GeoKey         = 3079
	tProjNatOriginLongGeoKey        = 3080
	tProjNatOriginLatGeoKey         = 3081
	tProjFalseEastingGeoKey         = 3082
	tProjFalseNorthingGeoKey        = 3083
	tProjFalseOriginLongGeoKey      = 3084
	tProjFalseOriginLatGeoKey       = 3085
	tProjFalseOriginEastingGeoKey   = 3086
	tProjFalseOriginNorthingGeoKey  = 3087
	tProjCenterLongGeoKey           = 3088
	tProjCenterLatGeoKey            = 3089
	tProjCenterEastingGeoKey        = 3090
	tProjCenterNorthingGeoKey       = 3091
	tProjScaleAtNatOriginGeoKey     = 3092
	tProjScaleAtCenterGeoKey        = 3093
	tProjAzimuthAngleGeoKey         = 3094
	tProjStraightVertPoleLongGeoKey = 3095
	tVerticalCSTypeGeoKey           = 4096
	tVerticalCitationGeoKey         = 4097
	tVerticalDatumGeoKey            = 4098
	tVerticalUnitsGeoKey            = 4099

	tPhotoshop = 34377
)

// Compression types (defined in various places in the spec and supplements).
const (
	cNone       = 1
	cCCITT      = 2
	cG3         = 3 // Group 3 Fax.
	cG4         = 4 // Group 4 Fax.
	cLZW        = 5
	cJPEGOld    = 6 // Superseded by cJPEG.
	cJPEG       = 7
	cDeflate    = 8 // zlib compression.
	cPackBits   = 32773
	cDeflateOld = 32946 // Superseded by cDeflate.
)

// Values for the tPredictor tag (page 64-65 of the spec).
const (
	prNone       = 1
	prHorizontal = 2
)

// Values for the tResolutionUnit tag (page 18).
const (
	resNone    = 1
	resPerInch = 2 // Dots per inch.
	resPerCM   = 3 // Dots per centimeter.
)

type imageMode int

const (
	mBilevel imageMode = iota
	mPaletted
	mGray
	mGrayInvert
	mRGB
	mRGBA
	mNRGBA
)

const (
	PI_WhiteIsZero = 0
	PI_BlackIsZero = 1
	PI_RGB         = 2
	PI_Paletted    = 3
	PI_TransMask   = 4 // transparency mask
	PI_CMYK        = 5
	PI_YCbCr       = 6
	PI_CIELab      = 8
)
