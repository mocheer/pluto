package baidupano

import (
	"testing"
	"time"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts/ctp"
)

type SData struct {
}

func TestBaiduMap(t *testing.T) {

	c := ctp.New()
	c.SetRequestTimeout(time.Second * 5)

	// sid := "0900210012200815111613880HO"
	// auth := "z6OIFeM5@a7G2aI667Q3D2R7BZ2Xfg8yuxNEHBNxTTHtF2F941W31QB2AyAT9xXwvkGcuVtcvY1SGpuEtzzljyYxvyheuzBtyEOIxXwvCQMuLxVtE5Wl1GDw8wkvyUuLoy9Ki3v@vcuVtvc3CuVtvcPPuztxgHxwzPD4vJtx7IKHwiKDv7uvhgMuzVVtvrMhuBTzEtJggagyYxegvcguxNEHBNxTTE"
	// secket := `drHbkBkj%2FHu3y7oNd1VZ4OLcsXjBFPExdgkEfbSuE6w%3D%2CAOSOnABh031VeAizN__5O0RGyIVo2ejGPk86_18n1WXJNugbrmFhIWABlhIXX2C2v6t3jgMCUEsIAuVdsZ96pI2wnRJY16Nmkenl7bRLtuaKqBc_HtHrS7_cWPK_ero3_m9x_6dDz6ySJnrk78lY4zgkvXtMoV6x9DrNg6UIQB-TeAHS_0nV7Xl1lgc7-KeI`
	// sdataURL := fmt.Sprintf(`https://mapsv0.bdimg.com/?qt=sdata&sid=%s&pc=1&auth=%s&seckey=%s&udt=20200825&fn=jsonp.p10585726`, sid, auth, secket)

	sdata, err := c.Get("https://mapsv1.bdimg.com/?qt=pdata&sid=09016200122207211646188987E&pos=0_0&z=1&udt=20200825&from=PC&auth=2wDWPMeYUX3KBwZOR8yGIwA%3DT%40FaHgLRuxNEHBTEzBEtF2F941W31QB34AACV%3Dz8yvkGcuVtvvhguVtvyheuBztGLVPFcEvCQMuxBzRtDOVjCEBw8wkvAFuzajPyYxv%40vcuVtcvY1SGpuxztGn6LFcEvcPPuztxgHxwBDDDv2quTG3FxjL1wWvvhgMuzVVtvrMhuHERztHee%40ewzvf0wd0vyOFICUCAyM&seckey=drHbkBkj%2FHu3y7oNd1VZ4OLcsXjBFPExdgkEfbSuE6w%3D%2CAOSOnABh031VeAizN__5O19MpRna2DtF83KZ04hos8-aVF5W5Dh3hNBJSf3atr0Z37PcYhH0mMb52C0Y6rZs05CxkCo1NcntyijGjMzF34BVsRoZ7UHAmTSLOpQq9igkEaRk-iG5ncJ0w_4qT-ZEhIMk9S3XLgQYQUegt-Dh9VLlzjHLNHPEEm-hPSwC7GU3")
	if err != nil {

	}

	ds.Save("./testdata/09016200122207211646188987E/0_0.png", sdata)
}
