package ds_zstd

import (
	"bytes"
	"io"
	"log"

	"github.com/klauspost/compress/dict"
	"github.com/klauspost/compress/zstd"
)

// Encode input to output.
func Encode(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	enc, err := zstd.NewWriter(b)
	if err != nil {
		return nil, err
	}
	enc.Write(data)
	enc.Close() // Close之后，data才能算真正写入，b.Bytes才有数据
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

func EncodeWithDict(data, dict []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	encoder, err := zstd.NewWriter(b, zstd.WithEncoderDict(dict))
	if err != nil {
		return nil, err
	}
	encoder.Write(data)
	encoder.Close()
	return b.Bytes(), nil
}

func DecodeWithDict(data, dict []byte) ([]byte, error) {
	dr := bytes.NewReader(data)
	decoder, err := zstd.NewReader(dr, zstd.WithDecoderDicts(dict))
	if err != nil {
		return nil, err
	}
	defer decoder.Close()

	result, err := io.ReadAll(decoder)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EncodeWithDict2
// dst 批量编码时可以复用dst，避免重复分配内存，make([]byte, 0, 4096)
// dict 字典
// data 数据
func EncodeWithDict2(dst, dict, data []byte) ([]byte, error) {
	encoder, err := zstd.NewWriter(nil, zstd.WithEncoderDict(dict))
	if err != nil {
		return nil, err
	}
	defer encoder.Close()

	result := encoder.EncodeAll(data, dst[:0])
	return result, nil
}

func DecodeWithDict2(dst, dict, data []byte) ([]byte, error) {
	decoder, err := zstd.NewReader(nil, zstd.WithDecoderDicts(dict))
	if err != nil {
		return nil, err
	}
	defer decoder.Close()

	result, err := decoder.DecodeAll(data, dst[:0])
	if err != nil {
		return nil, err
	}

	return result, nil
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
// 这个似乎只是实验性功能，不能很好的用于生成环境
// 训练的样本大小似乎很有要求，一旦太大（超过256KB）可能会报错? 如果有大量小样本生成字典的效率会变差。
func TrainDict(samples [][]byte) []byte {
	// 样本总大小最好是目标字典大小的100倍
	// 海量相似小文件，数据小于10KB左右的，一般设置字典大小是64KB-112KB
	// 完整的配置文件、中等长度的文本（如文章）、代码文件,一般设置为112KB-256KB
	// 数据越小、越相似，字典越有效，且字典本身可以更小（如64KB）。数据越大、越多样，字典收益越小，过大反而浪费内存。
	// 弃用zstd.BuildDict，因为dict.BuildZstdDict封装了zstd.BuildDict方法，使用起来更简便
	out := io.Discard
	dictData, err := dict.BuildZstdDict(samples, dict.Options{
		HashBytes:      6,                         // 最小匹配长度，单位B,更长的哈希匹配更精确但机会少,更短的哈希匹配更灵活但哈希冲突多
		ZstdLevel:      zstd.SpeedBestCompression, // 字典针对的压缩级别
		MaxDictSize:    4096,                      // 生成的字典大小限制,单位B,字典不建议超过 256KB,一般来说字典越大,压缩效率越高,但占用内存更多,解压效率也会变差
		ZstdDictID:     0,                         // Random
		ZstdDictCompat: false,
		Output:         out,
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	// zstd.WithEncoderDict(dictData)
	return dictData
}
