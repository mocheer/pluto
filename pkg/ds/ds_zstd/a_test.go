package ds_zstd_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ds/ds_zstd"
)

func TestEncode(t *testing.T) {
	ds.EachFiles("./testdata/train", func(filename string, fi os.FileInfo) {
		data, err := ds.ReadFile(filename)
		if err != nil {
			log.Println(err)
		}
		r, err := ds_zstd.Encode(data)
		if err != nil {
			t.Log(err)
			return
		}
		origin, err := ds_zstd.Decode(r)
		if err != nil {
			t.Log(err)
			return
		}

		if !bytes.Equal(origin, data) {
			log.Println("还原错误", len(origin))
		}
		ds.Save("./testdata/zst/"+fi.Name()+".text", r)
		log.Println(len(data), len(r), len(data)-len(r))
	})

}

func TestEncode2(t *testing.T) {

	ds.EachFiles("./testdata/train", func(filename string, fi os.FileInfo) {
		b := &bytes.Buffer{}
		f, _ := os.Open(filename)
		err := ds_zstd.Compress(f, b)
		if err != nil {
			t.Log(err)
			return
		}
		ds.Save("./testdata/zst/"+fi.Name(), b.Bytes())
	})

}

// 失败原因是样本太大??
// @see https://github.com/klauspost/compress/issues/1042
// 太大会出现 panic: runtime error: slice bounds out of range [-129026:] [recovered, repanicked]
// 太小压缩效率低
func TestTrain(t *testing.T) {

	var samples [][]byte

	ds.EachFiles("./testdata/train", func(filename string, fi os.FileInfo) {
		data, err := ds.ReadFile(filename)
		if err != nil {
			t.Log(filename, err)
		}

		if len(data) <= 128 {
			return
		}
		// 大于256KB报错？
		if len(data) >= 1024*256 {
			t.Log(filename, "数据样本大小不适合", len(data))
			return
		}

		samples = append(samples, data)
	})
	//
	d := ds_zstd.TrainDict(samples)
	ds.Save("./testdata/dict/train.dict", d)
}

func TestDict(t *testing.T) {
	// testdata/train/hrbxt
	// testdata/train/6VbjevBv
	data, err := ds.ReadFile("./testdata/train/6VbjevBv")
	if err != nil {
		t.Log(err)
		return
	}
	// testdata\dict\train.dict
	dict, _ := ds.ReadFile("./testdata/dict/train.dict")
	r, err := ds_zstd.EncodeWithDict(data, dict)
	if err != nil {
		t.Log(err)
		return
	}
	r2, err := ds_zstd.Encode(data)
	if err != nil {
		t.Log(err)
		return
	}
	origin, err := ds_zstd.DecodeWithDict(r, dict)
	if err != nil {
		t.Log(err)
		return
	}
	if !bytes.Equal(origin, data) {
		t.Log("还原错误", len(origin), string(origin))
	}
	totalLength := float64(len(data))
	t.Log(totalLength, float64(len(r))/totalLength)
	t.Log(totalLength, float64(len(r2))/totalLength)
	t.Log(len(r2) - len(r))
	// t.Log(string(origin))

}
