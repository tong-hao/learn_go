package main

import (
	"bufio"
	"fmt"
	hdfs "github.com/colinmarc/hdfs/v2"
	"strings"
)

func readHdfsColum(hdfsPath string, client  *hdfs.Client, column int) []string {
	ret := make([]string, 0)
	fInfo, _ := client.ReadDir(hdfsPath)
	for _, file := range fInfo {
		f := fmt.Sprintf("%s/%s", hdfsPath, file.Name())
		if file.IsDir() {
			continue
		}
		fmt.Println(f)
		reader,_ := client.Open(f)
		s := bufio.NewScanner(reader)
		for s.Scan() {
			line := s.Text()
			strs := strings.Split(line, ",")
			if len(strs) >= column {
				ret = append(ret, strs[column-1])
			}
		}
		reader.Close()
	}

	return ret
}

func main() {
	nameNode := "192.168.8.240:9000"
	userName := "zhuang.miao"


	hasUser := true
	hdfsPath := "/caton.hao/"
	var client  *hdfs.Client
	// write with user
	if hasUser {
		var opts hdfs.ClientOptions
		opts.User = userName
		opts.Addresses = []string{nameNode}
		client,_ = hdfs.NewClient(opts)
	} else {
		client, _ = hdfs.New(nameNode)
	}


	// write
	//{
	//	writer, err := client.Create(file)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	for i := 0; i < 100; i++ {
	//		str := fmt.Sprintf("%d,%d,%d\n", i, i+1, i+2)
	//		writer.Write([]byte(str))
	//	}
	//	writer.Flush()
	//	writer.Close()
	//}


	ret := readHdfsColum(hdfsPath, client, 1)
	for _,line := range ret {
		fmt.Println(line)
	}

}
