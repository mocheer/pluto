package ds_dat_test

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"
)

func Test1(t *testing.T) {
	// 打开.dat文件
	file, err := os.Open("testdata/SL_BO_20230812_075513_SA370000001M.dat")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	// 声明变量用于解码二进制数据
	//
	var volumeLabel [4]byte
	var versionNo [4]byte
	var fileLength uint32
	var rayOrder uint32 // unint16
	err = binary.Read(file, binary.BigEndian, &volumeLabel)
	err = binary.Read(file, binary.BigEndian, &versionNo)
	err = binary.Read(file, binary.BigEndian, &fileLength)
	err = binary.Read(file, binary.BigEndian, &rayOrder)
	fmt.Println(string(volumeLabel[0:]))
	fmt.Println(string(versionNo[0:]))
	fmt.Println(fileLength)
	fmt.Println(rayOrder)
	// 读取到终止符需要自动停止
	var country [6]byte
	err = binary.Read(file, binary.LittleEndian, &country)

	fmt.Println(string(country[:]))

	var province [32]byte
	err = binary.Read(file, binary.BigEndian, &province)
	fmt.Println("province:", string(province[0:]), province)

}
