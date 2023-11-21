package cntax

import (
	"archive/zip"
	"encoding/xml"
	"strconv"
)

// 解析发票基础数据
const TaxTypeBase = "OFD.xml"

type OFD struct {
	XMLName xml.Name `xml:"OFD"`
	DocBody DocBody  `xml:"DocBody"`
}

type DocBody struct {
	XMLName xml.Name `xml:"DocBody"`
	DocInfo DocInfo  `xml:"DocInfo"`
}

type DocInfo struct {
	XMLName      xml.Name    `xml:"DocInfo"`
	DocID        string      `xml:"DocID"`
	Author       string      `xml:"Author"`
	CreationDate string      `xml:"CreationDate"`
	ModDate      string      `xml:"ModDate"`
	Creator      string      `xml:"Creator"`
	CustomDatas  CustomDatas `xml:"CustomDatas"`
}

type CustomDatas struct {
	XMLName    xml.Name     `xml:"CustomDatas"`
	CustomData []CustomData `xml:"CustomData"`
}

type CustomData struct {
	XMLName xml.Name `xml:"CustomData"`
	Name    string   `xml:"Name,attr"`
	Val     string   `xml:",chardata"`
}

// 解析基础文件
func (tf *TaxFile) parseBaseFile(f *zip.File) error {
	var ofd OFD
	// 打开文件
	file, err := f.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	buf := make([]byte, f.FileInfo().Size())
	file.Read(buf)
	err = xml.Unmarshal(buf, &ofd)
	if err != nil {
		return err
	}
	//遍历内容
	for _, v := range ofd.DocBody.DocInfo.CustomDatas.CustomData {
		var e error
		if v.Name == "发票号码" {
			tf.DocNum = v.Val
		}
		if v.Name == "销售方纳税人识别号" {
			tf.FromNo = v.Val
		}
		if v.Name == "购买方纳税人识别号" {
			tf.ToNo = v.Val
		}
		if v.Name == "开票日期" {
			tf.Data = v.Val
		}
		// 主页金额不含税
		// if v.Name == "合计金额" {
		// 	tf.Total, e = strconv.ParseFloat(v.Val, 64)
		// 	if e != nil {
		// 		return
		// 	}
		// }
		if v.Name == "合计税额" {
			tf.TotalTax, e = strconv.ParseFloat(v.Val, 64)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

// 解析发票内容
const TaxContent = "Doc_0/Pages/Page_0/Content.xml"

type Page struct {
	XMLName xml.Name    `xml:"Page"`
	Content PageContent `xml:"Content"`
}
type PageContent struct {
	XMLName xml.Name  `xml:"Content"`
	Layer   PageLayer `xml:"Layer"`
}
type PageLayer struct {
	XMLName    xml.Name     `xml:"Layer"`
	TextObject []TextObject `xml:"TextObject"`
}
type TextObject struct {
	XMLName xml.Name `xml:"TextObject"`
	ID      string   `xml:"ID,attr"`
	Val     string   `xml:"TextCode"`
}

func (tf *TaxFile) parseContent(f *zip.File) error {
	var p Page

	file, err := f.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	buf := make([]byte, f.FileInfo().Size())
	file.Read(buf)
	if err = xml.Unmarshal(buf, &p); err != nil {
		return err
	}

	for _, v := range p.Content.Layer.TextObject {
		var e error
		if v.ID == "6924" {
			tf.ToName = v.Val
		}
		if v.ID == "6927" {
			tf.FromName = v.Val
		}
		if v.ID == "6934" {
			tf.ToalCn = v.Val
		}
		if v.ID == "6937" {
			tf.Operator = v.Val
		}
		if v.ID == "6936" {
			tf.Total, e = strconv.ParseFloat(v.Val, 64)
			if e != nil {
				return e
			}
		}

	}
	return nil
}

// 解析发票类型
const TaxTypeFile = "Doc_0/Tpls/Tpl_0/Content.xml"

type TplPage struct {
	XMLName xml.Name       `xml:"Page"`
	Content TplPageContent `xml:"Content"`
}
type TplPageContent struct {
	XMLName xml.Name       `xml:"Content"`
	Layer   []TplPageLayer `xml:"Layer"`
}
type TplPageLayer struct {
	XMLName    xml.Name        `xml:"Layer"`
	TextObject []TplTextObject `xml:"TextObject"`
}
type TplTextObject struct {
	XMLName xml.Name `xml:"TextObject"`
	ID      string   `xml:"ID,attr"`
	Val     string   `xml:"TextCode"`
}

func (tf *TaxFile) parseTaxTypeFile(f *zip.File) error {
	var p TplPage
	// 打开文件
	file, err := f.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	buf := make([]byte, f.FileInfo().Size())
	file.Read(buf)
	if err = xml.Unmarshal(buf, &p); err != nil {
		return err
	}
	for _, l := range p.Content.Layer {
		for _, o := range l.TextObject {
			//发票类型Id
			if o.ID == "3" {
				if o.Val == "电子发票（普通发票）" {
					tf.Type = TypeNormal
				} else {
					tf.Type = TypeSpecial
				}
			}
		}
	}
	return nil
}
