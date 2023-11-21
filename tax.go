package cntax

import (
	"archive/zip"
	"os"
)

const (
	TypeNormal  = 1 //普通发票
	TypeSpecial = 2 //专票
)

// TaxFile 基础结构
type TaxFile struct {
	path     string      //文件路径
	file     *zip.Reader //文件
	Type     int         //发票类型
	DocNum   string      //发票编号
	FromName string      //发票发起公司
	FromNo   string      //发票发起公司税务编号
	ToName   string      //发票对象公司
	ToNo     string      //发票对象公司税务编号
	Total    float64     //总金额
	ToalCn   string      //总金额大写
	TotalTax float64     //总税额
	Operator string      //开票人
	Data     string      //开票日期
}

// NewTaxFile creaet a new TaxFile
func NewTaxFile(path string) (*TaxFile, error) {
	f := new(TaxFile)
	f.path = path
	// 打开压缩文件
	zipFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()
	stat, _ := zipFile.Stat()
	// 创建zip文件读取器
	f.file, err = zip.NewReader(zipFile, stat.Size())
	if err != nil {
		return nil, err
	}

	// 遍历压缩包中的文件
	for _, file := range f.file.File {
		if file.Name == TaxTypeBase {
			if e := f.parseBaseFile(file); e != nil {
				return nil, e
			}
		}
		if file.Name == TaxContent {
			if e := f.parseContent(file); e != nil {
				return nil, e
			}
		}
		if file.Name == TaxTypeFile {
			if e := f.parseTaxTypeFile(file); e != nil {
				return nil, e
			}
		}
	}
	return f, nil
}
