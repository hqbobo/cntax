package cntax

import (
	"encoding/xml"
	"io/ioutil"
)

type Header struct {
	XMLName     xml.Name `xml:"Header"`
	EIid        string   `xml:"EIid"`
	EInvoiceTag string   `xml:"EInvoiceTag"`
	Version     string   `xml:"Version"`
}

type SellerInformation struct {
	XMLName          xml.Name `xml:"SellerInformation"`
	SellerIdNum      string   `xml:"SellerIdNum"`
	SellerName       string   `xml:"SellerName"`
	SellerAddr       string   `xml:"SellerAddr"`
	SellerTelNum     string   `xml:"SellerTelNum"`
	SellerBankName   string   `xml:"SellerBankName"`
	SellerBankAccNum string   `xml:"SellerBankAccNum"`
}

type BuyerInformation struct {
	XMLName           xml.Name `xml:"BuyerInformation"`
	BuyerIdNum        string   `xml:"BuyerIdNum"`
	BuyerName         string   `xml:"BuyerName"`
	BuyerHandlingName string   `xml:"BuyerHandlingName"`
}

type BasicInformation struct {
	XMLName                         xml.Name `xml:"BasicInformation"`
	TotalAmWithoutTax               float64  `xml:"TotalAmWithoutTax"`
	TotalTaxincludedAmount          float64  `xml:"TotalTax-includedAmount"`
	TotalTaxincludedAmountInChinese string   `xml:"TotalTax-includedAmountInChinese"`
	Drawer                          string   `xml:"Drawer"`
	RequestTime                     string   `xml:"RequestTime"`
}
type TaxSupervisionInfo struct {
	XMLName       xml.Name `xml:"TaxSupervisionInfo"`
	InvoiceNumber string   `xml:"InvoiceNumber"`
	IssueTime     string   `xml:"IssueTime"`
	TaxBureauCode string   `xml:"TaxBureauCode"`
	TaxBureauName string   `xml:"TaxBureauName"`
}
type TaxBureauSignature struct {
	XMLName            xml.Name `xml:"TaxBureauSignature"`
	Reference          string   `xml:"Reference"`
	SignatureAlgorithm string   `xml:"SignatureAlgorithm"`
	SignatureFormat    string   `xml:"SignatureFormat"`
	SignatureTime      string   `xml:"SignatureTime"`
	SignatureValue     string   `xml:"SignatureValue"`
}
type EInvoiceData struct {
	XMLName           xml.Name          `xml:"EInvoiceData"`
	SellerInformation SellerInformation `xml:"SellerInformation"`
	BuyerInformation  BuyerInformation  `xml:"BuyerInformation"`
	BasicInformation  BasicInformation  `xml:"BasicInformation"`
}

type EInvoice struct {
	XMLName            xml.Name           `xml:"EInvoice"`
	Header             Header             `xml:"Header"`
	EInvoiceData       EInvoiceData       `xml:"EInvoiceData"`
	TaxSupervisionInfo TaxSupervisionInfo `xml:"TaxSupervisionInfo"`
	TaxBureauSignature TaxBureauSignature `xml:"TaxBureauSignature"`
}

// NewTaxEInvoice creaet a new TaxFile
func NewTaxEInvoice(path string) (*EInvoice, error) {
	f := new(EInvoice)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(buf, f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// NewTaxEInvoiceByte creaet a new TaxFile with byte
func NewTaxEInvoiceByte(data []byte) (*EInvoice, error) {
	f := new(EInvoice)
	err := xml.Unmarshal(data, f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
