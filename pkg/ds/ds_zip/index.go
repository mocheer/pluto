package ds_zip

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"strings"

	"github.com/mocheer/pluto/pkg/fn"
)

// ReadFileItem
// 即使文件在多层文件夹内部也能获取到，文件名为：文件夹/文件夹/**/名称
func ReadFileItem(fileName string, itemFileName string) ([]byte, error) {
	// 读取
	r, err := zip.OpenReader(fileName)

	if err == nil {
		defer r.Close()
		// 遍历所有文件，包括文件夹本身
		for _, f := range r.File {
			if f.Name == itemFileName {
				reader, err := f.Open()

				defer reader.Close()
				if err == nil {

					data, err := io.ReadAll(reader)
					return data, err
				}
				return nil, err
			}
		}
		err = errors.New("not found")
	}

	return nil, err
}

// Each 遍历zip文件
func Each(fileName string, callback func(*zip.File)) error {
	// 读取
	r, err := zip.OpenReader(fileName)
	// zip: not a valid zip file [recovered]
	// fn.Panic(err, fileName+"不是有效的zip文件")
	if err != nil {
		return err
	}
	defer r.Close()
	// 遍历所有文件，包括文件夹本身
	for _, f := range r.File {
		callback(f)
	}
	return nil
}

// EachByReader
// EachByReader(file,file.Size)
// EachByReader(bytes.NewReader(bs),len(bs))
func EachByReader(reader io.ReaderAt, size int64, callback func(*zip.File)) error {
	r, err := zip.NewReader(reader, size)
	if err != nil {
		return err
	}
	// 遍历所有文件，包括文件夹本身
	for _, f := range r.File {
		callback(f)
	}
	return nil
}

// EachFiles 遍历zip文件
func EachFiles(fileName string, callback func(*zip.File)) error {
	return Each(fileName, func(f *zip.File) {
		info := f.FileInfo()
		if !info.IsDir() {
			callback(f)
		}
	})
}

// EachFiles 遍历zip文件
func EachFilesReader(fileName string, callback func(io.ReadCloser)) error {
	return EachFiles(fileName, func(f *zip.File) {
		reader, err := f.Open()
		if err != nil {
			fn.Panic(err, "无法解压")
		}
		defer reader.Close()
		callback(reader)
	})
}

// EachFilesBytes 遍历zip文件
func EachFilesBytes(fileName string, callback func([]byte)) error {
	return EachFilesReader(fileName, func(r io.ReadCloser) {
		bs, err := io.ReadAll(r)
		if err != nil {
			fn.Panic(err, "无法读取")
		}
		callback(bs)
	})
}

// Extract 解压缩zip文件
func Extract(fileName string) {
	// 读取
	Each(fileName, func(f *zip.File) {
		info := f.FileInfo()
		if info.IsDir() {
			err := os.MkdirAll(f.Name, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			return
		}
		srcFile, err := f.Open()
		if err != nil {
			panic(err.Error())
		}
		defer srcFile.Close()

		newFile, err := os.Create(f.Name)
		if err != nil {
			panic(err.Error())
		}
		defer newFile.Close()

		io.Copy(newFile, srcFile)
	})

}

// Write 创建压缩包并写入文件
func Write(data []byte, fileName string) {
	// 缓存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建zip
	writer := zip.NewWriter(buf)
	defer writer.Close()

	// 接收
	f, _ := writer.Create(fileName)
	f.Write(data)

	fileName = strings.TrimSuffix(fileName, path.Ext(fileName)) + ".zip"
	os.WriteFile(fileName, buf.Bytes(), 0644)
}
