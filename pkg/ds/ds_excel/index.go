package ds_excel

import (
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

// @see https://xuri.me/excelize/zh-hans/

// ReadFile
func ReadFile(fileName string) ([][]string, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	// 单个row包含每个col的string值
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Read
func Read(r io.Reader) ([][]string, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	// 单个row包含每个col的string值
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
