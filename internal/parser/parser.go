package parser

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func (p *StrParser) Settings(ctx context.Context, datas Datas) error {
	if err := p.checkAndSetCookie(ctx, datas); err != nil {
		return err
	}
	return nil
}

var datas []Datas

func (p *StrParser) GetData() ([]Datas, error) {
	datas = []Datas{}
	dataFile, err := os.ReadDir(Free_account)
	if err != nil {
		return nil, fmt.Errorf("read dir error: %w", err)
	}
	if len(dataFile) == 0 {
		log.Print("empty free_account dir")
		return nil, ErrEmptyData
	}
	for _, file := range dataFile {
		log.Printf("open file: %s", file.Name())
		dataFile, err := os.Open(Free_account + file.Name())
		if err != nil {
			return nil, fmt.Errorf("open data file error: %w", err)
		}
		defer func() {
			if err := dataFile.Close(); err != nil {
				log.Fatalf("data file close error: %v", err)
			}
		}()
		data, err := io.ReadAll(dataFile)
		if err != nil {
			return nil, fmt.Errorf("read data file error: %w", err)
		}
		strDatas := Datas{}
		strDatas.FileName = file.Name()

		strData := Data{}
		data1 := strings.ReplaceAll(string(data), "\r", "")
		splitData := strings.Split(data1, "\n")

		loginpass := splitData[0]
		datalogin := strings.Split(loginpass, ":")
		login := datalogin[0]
		pass := datalogin[1]

		strData.LoginFB = login
		strData.PassFB = pass

		proxy := splitData[1]
		dataproxy := strings.Split(proxy, "@")
		ipport := dataproxy[0]
		loginpassproxy := dataproxy[1]
		loginpassPX := strings.Split(loginpassproxy, ":")
		loginPX := loginpassPX[0]
		passPX := loginpassPX[1]

		strData.IpPortPX = ipport
		strData.LoginPX = loginPX
		strData.PassPX = passPX

		cookie := splitData[2]

		strData.Cookies = cookie
		strDatas.Datas = strData

		datas = append(datas, strDatas)
	}
	log.Printf("datas: %v", datas)
	return datas, nil
}
