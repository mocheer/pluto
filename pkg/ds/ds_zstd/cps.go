package ds_zstd

import (
	"bytes"
	"io"

	"github.com/klauspost/compress/dict"
	"github.com/klauspost/compress/zstd"
)

// Encode input to output.
func Encode(data []byte) ([]byte, error) {
	var b bytes.Buffer
	enc, err := zstd.NewWriter(&b)
	if err != nil {
		return nil, err
	}
	defer enc.Close()
	enc.Write(data)
	return b.Bytes(), err
}

// Decode
func Decode(data []byte) ([]byte, error) {
	dr := bytes.NewReader(data)
	d, err := zstd.NewReader(dr)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	buf, err := io.ReadAll(d)
	if err != nil {
		return buf, err
	}
	return buf, err
}

// Compress input to output.
func Compress(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out)
	if err != nil {
		return err
	}
	_, err = io.Copy(enc, in)
	if err != nil {
		enc.Close()
		return err
	}
	return enc.Close()
}

// Decompress
func Decompress(in io.Reader, out io.Writer) error {
	d, err := zstd.NewReader(in)
	if err != nil {
		return err
	}
	defer d.Close()

	// Copy content...
	_, err = io.Copy(out, d)
	return err
}

// TrainDict
// samples 样本，一般是来源于多个文件
func TrainDict(samples [][]byte) []byte {
	// 弃用 zstd.BuildDict
	// zstd.BuildDict()
	// 样本总大小最好是目标字典大小的100倍
	// 海量相似小文件，数据小于10KB左右的，一般设置字典大小是64KB-112KB
	// 完整的配置文件、中等长度的文本（如文章）、代码文件,一般设置为112KB-256KB
	// 数据越小、越相似，字典越有效，且字典本身可以更小（如64KB）。数据越大、越多样，字典收益越小，过大反而浪费内存。
	dictData, err := dict.BuildZstdDict(samples, dict.Options{
		MaxDictSize: 131072,                    // 字典大小，这里设置为128KB，字典大小不建议超过 256KB
		HashBytes:   6,                         // 最小匹配长度，通常4-8
		ZstdLevel:   zstd.SpeedBestCompression, // 字典针对的压缩级别

	})
	if err != nil {
		return nil
	}

	// zstd.WithEncoderDict(dictData)
	return dictData
}
