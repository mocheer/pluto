# tif

## 压缩算法

|          | Decoding | Encoding |
| -------- | -------- | -------- |
| None     | ✓        | ✓        |
| LZW      | ✓        | ✓        |
| Deflate  | ✓        | ✓        |
| PackBits | ✓        | ✓        |
| JPEG     | ✓        | ✓        |
| LERC     | ✓        | ✓        |

## 开发库

### jblindsay/go-spatial
- 在tiff格式之外还支持其他栅格数据文件
  - rst
  - dep
- 不支持bigtiff
- 会将ifds合并，排在后面的ifd优先级更高，这个机制应该是错的

### google/tiff
- 没有直接支持获取高程数据（灰度值），包括相关的lzw压缩算法等（在另外的库中，且lzw算法目前有问题，无法解析部分lzw压缩的tiff文件）

### google/tiff/geotiff
- BigTiff的文件头固定为8个字节，分别为49 49 2B 00 08 00 00 00。读取程序检测得到这8个字节即可判定文件为BigTiff格式。
- BigTiff的文件尾固定为8个全零字节

### geotiffjs/geotiff
这是一个js库

- 可解析`golang.org/x/image/tiff/lzw`算法不支持的tiff文件
- 压缩算法还支持LERC、JPEG

### image-rs/image-tiff
这是一个rust库




```go
// Tags (see p. 28-41 of the spec).
var tagMap = map[int]string{
	254: "NewSubFileType",
	256: "ImageWidth",
	257: "ImageLength",
	258: "BitsPerSample",
	259: "Compression",
	262: "PhotometricInterpretation",
	266: "FillOrder",
	269: "DocumentName",
	284: "PlanarConfiguration",
	270: "ImageDescription",
	271: "Make",
	272: "Model",
	273: "StripOffsets",
	274: "Orientation",
	277: "SamplesPerPixel",
	278: "RowsPerStrip",
	279: "StripByteCounts",
	280: "MinSampleValue",
	281: "MaxSampleValue",
	282: "XResolution",
	283: "YResolution",
	296: "ResolutionUnit",
	305: "Software",
	306: "DateTime",
	322: "TileWidth",
	323: "TileLength",
	324: "TileOffsets",
	325: "TileByteCounts",
	317: "Predictor",
	320: "ColorMap",
	338: "ExtraSamples",
	339: "SampleFormat",
	34735: "GeoKeyDirectoryTag",
	34736: "GeoDoubleParamsTag",
	34737: "GeoAsciiParamsTag",
	33550: "ModelPixelScaleTag",
	33922: "ModelTiepointTag",
	34264: "ModelTransformationTag",
	42112: "GDAL_METADATA",
	42113: "GDAL_NODATA",
	1024:  "GTModelTypeGeoKey",
	1025:  "GTRasterTypeGeoKey",
	1026:  "GTCitationGeoKey",
	2048:  "GeographicTypeGeoKey",
	2049:  "GeogCitationGeoKey",
	2050:  "GeogGeodeticDatumGeoKey",
	2051:  "GeogPrimeMeridianGeoKey",
	2061:  "GeogPrimeMeridianLongGeoKey",
	2052:  "GeogLinearUnitsGeoKey",
	2053:  "GeogLinearUnitSizeGeoKey",
	2054:  "GeogAngularUnitsGeoKey",
	2055:  "GeogAngularUnitSizeGeoKey",
	2056:  "GeogEllipsoidGeoKey",
	2057:  "GeogSemiMajorAxisGeoKey",
	2058:  "GeogSemiMinorAxisGeoKey",
	2059:  "GeogInvFlatteningGeoKey",
	2060:  "GeogAzimuthUnitsGeoKey",
	3072:  "ProjectedCSTypeGeoKey",
	3073:  "PCSCitationGeoKey",
	3074:  "ProjectionGeoKey",
	3075:  "ProjCoordTransGeoKey",
	3076:  "ProjLinearUnitsGeoKey",
	3077:  "ProjLinearUnitSizeGeoKey",
	3078:  "ProjStdParallel1GeoKey",
	3079:  "ProjStdParallel2GeoKey",
	3080:  "ProjNatOriginLongGeoKey",
	3081:  "ProjNatOriginLatGeoKey",
	3082:  "ProjFalseEastingGeoKey",
	3083:  "ProjFalseNorthingGeoKey",
	3084:  "ProjFalseOriginLongGeoKey",
	3085:  "ProjFalseOriginLatGeoKey",
	3086:  "ProjFalseOriginEastingGeoKey",
	3087:  "ProjFalseOriginNorthingGeoKey",
	3088:  "ProjCenterLongGeoKey",
	3089:  "ProjCenterLatGeoKey",
	3090:  "ProjCenterEastingGeoKey",
	3091:  "ProjFalseOriginNorthingGeoKey",
	3092:  "ProjScaleAtNatOriginGeoKey",
	3093:  "ProjScaleAtCenterGeoKey",
	3094:  "ProjAzimuthAngleGeoKey",
	3095:  "ProjStraightVertPoleLongGeoKey",
	4096:  "VerticalCSTypeGeoKey",
	4097:  "VerticalCitationGeoKey",
	4098:  "VerticalDatumGeoKey",
	4099:  "VerticalUnitsGeoKey",
	50844: "RPCCoefficientTag",
	34377: "Photoshop",
}
```

## 参考
- https://github.com/jblindsay/go-spatial
- https://github.com/google/tiff
- https://github.com/geotiffjs/geotiff.js
- https://search.asf.alaska.edu
- http://www.tuxingis.com
- http://www.gscloud.cn/   
- https://github.com/image-rs/image-tiff

