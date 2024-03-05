package ds_ssh_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/mocheer/pluto/pkg/ds_ssh"
)

func TestXxx(t *testing.T) {
	// 调用函数，传入相应参数
	client, err := ds_ssh.New("root", "JSPTB_201Created_@)!_103", "192.168.118.103:22")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	content, err := client.ReadFile("/vdb/istrong/charon/public/configs/app.toml")
	fmt.Println(string(content))
}

func TestDir(t *testing.T) {
	// 调用函数，传入相应参数
	client, err := ds_ssh.New("root", "JSPTB_201Created_@)!_138", "192.168.118.138:22")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	content, err := client.GetFilenames("/strong/backup/postgis_data")
	fmt.Println(content)
}
