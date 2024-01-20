package ds_tif_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_tif"
	"github.com/mocheer/xena/pkg/proj4"
	"github.com/samber/lo"
)

func TestTiff(t *testing.T) {
	fileName := "./testdata/SRTM3_V4_90m.tif" // GCS_WGS_1984
	tfs, _ := ds_tif.Read(fileName)

	t.Log(tfs.Tif.IFDs()[0])
	t.Log(string(tfs.Tif.IFDs()[0].Fields()[18].Value().Bytes()))
	// 107.85853701103325 34.12743138973333 107.34437034436667 33.849098056400045
	t.Log(tfs.BBox())
	// 100,100
	t.Log(tfs.GetColRow(107.42770367769998, 34.04409805640001))
	// 107.42770367769998, 34.04409805640001
	t.Log(tfs.GetLonLat(100, 100))
	// 1757
	t.Log(tfs.GetAlt(100, 100))

}

func TestBids_tiff(t *testing.T) {
	// UTM 投影：经度采用6度分带，纬度采用8度分带，从80S到84N共20个纬度带（X带多4度），分别用C到X的字母来表示
	fileName := "./testdata/utm/54R_20200101-20210101.tif" //这是一个bids_tiff格式的文件
	tfs, _ := ds_tif.Read(fileName)

	tfs.IFD_Index = 0
	t.Log(tfs.Tif.IFDs())
	t.Log(tfs.Origin())
	t.Log(tfs.BBox())
	t.Log(tfs.Scale())
	//
	t.Log(tfs.GetLonLat(100, 100))
	xy := tfs.GetLonLat(100, 100)
	t.Log(proj4.UTM_WGS84_ZONE(54).Inverse([]float64{xy[0], xy[1]}))
}

func TestTiff2(t *testing.T) {
	fileName := "./testdata/ASTGTMV003_N53E125_dem.tif" // lzw 压缩
	tfs, _ := ds_tif.Read(fileName)

	t.Log(len(tfs.Tif.IFDs()))
	t.Log(tfs.Width())
	t.Log(tfs.Height())
	t.Log(tfs.Origin())
	t.Log(tfs.BBox())
	t.Log(tfs.Scale())
	//
	t.Log(tfs.GetLonLat(100, 3600))
	t.Log(tfs.GetAlt(2, 0))
	t.Log(tfs.GetAltByLonLat(125, 53))

}

func TestTiff3(t *testing.T) {
	fileName := "./testdata/ASTGTMV003_N34E107_dem.tif"
	tfs, _ := ds_tif.Read(fileName)

	t.Log(len(tfs.Tif.IFDs()))
	t.Log(tfs.Width())
	t.Log(tfs.Height())
	t.Log(tfs.Origin())
	t.Log(tfs.BBox())
	t.Log(tfs.Scale())
	//
	t.Log(tfs.GetLonLat(100, 3600))
	t.Log("alt", tfs.GetAlt(2, 0))
	t.Log(tfs.GetAltByLonLat(107.42770367769998, 34.04409805640001))

}

func TestTiff4(t *testing.T) {
	fileName := "./testdata/ASTGTMV003_N43E093_dem.tif"
	tfs, _ := ds_tif.Read(fileName)
	//
	t.Log(len(tfs.Tif.IFDs()))
	t.Log(tfs.Width())                              // 宽度
	t.Log(tfs.Height())                             // 高度
	t.Log(tfs.Origin())                             // 起点
	t.Log(tfs.BBox())                               // 范围
	t.Log(tfs.Scale())                              // 系数
	t.Log(tfs.Scale()[0] * 111110)                  // 30.86 说明这个水平分辨率只有30米，这个算法是一个模糊值
	t.Log(tfs.GetColRow(93.878768, 43.074783))      // 获取经纬度对应的网格位置
	t.Log(tfs.GetAltByLonLat(93.878768, 43.074783)) // 获取当前经纬度对应的高度 1458 => 1455.0753793690708  Cesium.sampleTerrainMostDetailed(terrainProvider,[Cartographic.fromDegrees()])
	t.Log(tfs.GetAltByLonLat(93.278768, 43.074783))
}

func TestTiff5(t *testing.T) {
	fileName := "./testdata/ASTGTMV003_N03E112_dem.tif"
	tfs, err := ds_tif.Read(fileName)
	if err != nil {
		t.Fatal(err)
	}
	//
	t.Log(len(tfs.Tif.IFDs()))
	t.Log(tfs.Width())             // 宽度
	t.Log(tfs.Height())            // 高度
	t.Log(tfs.Origin())            // 起点
	t.Log(tfs.BBox())              // 范围
	t.Log(tfs.Scale())             // 系数
	t.Log(tfs.Scale()[0] * 111110) // 30.86 说明这个水平分辨率只有30米，这个算法是一个模糊值
	t.Log(len(lo.Filter(tfs.Data, func(v float64, _ int) bool { return v != 0 })))
	t.Log(string(tfs.Tif.IFDs()[0].Fields()[18].Value().Bytes()))
}

func TestTiff_llsk(t *testing.T) {
	// 岭里水库
	fileName := "./testdata/llsk/Extract_crea61.tif" //高斯投影3度带，117
	tfs, _ := ds_tif.Read(fileName)
	ifd := tfs.Tif.IFDs()[0]
	t.Log(ifd)
	t.Log((ifd.Fields()[16]))
	t.Log(string(ifd.Fields()[16].Value().Bytes()))
	t.Log(string(ifd.Fields()[17].Value().Bytes()))
	// 107.85853701103325 34.12743138973333 107.34437034436667 33.849098056400045
	t.Log(tfs.BBox())

}

func TestTiff_nc(t *testing.T) {
	fileName := "./testdata/nc.tif"
	tfs, _ := ds_tif.Read(fileName)
	ifd := tfs.Tif.IFDs()[0]
	t.Log(ifd)
	t.Log(tfs.Width())
	t.Log(tfs.Height())
	t.Log(tfs.BBox())
	t.Log(tfs.Scale())
	t.Log(tfs.Origin())
	// ds_json.Save("./testdata/nc.json", tfs.Data)
}

func TestTiff2_nc(t *testing.T) {
	fileName := "./testdata/20230812_164500.tif"
	tfs, _ := ds_tif.Read(fileName)
	ifd := tfs.Tif.IFDs()
	t.Log(ifd[0], len(ifd))
}

func TestTiff3_mq(t *testing.T) {
	fileName := "./testdata/闽清倾斜摄影对应DEM/10M.tif"
	tfs, _ := ds_tif.Read(fileName)
	ifd := tfs.Tif.IFDs()
	t.Log(ifd[0], len(ifd))
	t.Log(tfs.Origin())
	t.Log(tfs.BBox())
	t.Log(tfs.Scale())
	//
	t.Log(tfs.GetAlt(1000, 1000))
	t.Log(tfs.GetLonLat(1000, 1000))
	t.Log(tfs.GetAltByLonLat(674663.2124192547, 2874738.825797036))
	t.Log(tfs.GetAltByLonLat(674669.8599475692, 2875052.5104912943))

	t.Log(tfs.Scale()[0])
	// ds_json.Save("./testdata/nc.json", tfs.Data)

}

func TestTiff34(t *testing.T) {
	fileName := "./testdata/alos/AP_27052_FBS_F0740_RT1/AP_27052_FBS_F0740_RT1.dem.tif"
	tfs, _ := ds_tif.Read(fileName)
	ifd := tfs.Tif.IFDs()
	t.Log(ifd[0], len(ifd))
	t.Log(tfs.Origin())
	t.Log(tfs.BBox())
	t.Log(tfs.Scale())
	//
	t.Log(tfs.GetAlt(1000, 1000))
	t.Log(tfs.GetLonLat(1000, 1000))
	t.Log(tfs.GetAltByLonLat(674663.2124192547, 2874738.825797036))
	t.Log(tfs.GetAltByLonLat(674669.8599475692, 2875052.5104912943))

	t.Log(tfs.Scale()[0])

}
