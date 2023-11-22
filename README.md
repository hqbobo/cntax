# cntax

税务电子发票解析 (ofd)

this is used for parse CN tax file 
it require ofd type file as input and find out its key info

the key info defined

```

// TaxProject 发票项目信息
type TaxProject struct {
	Name    string  //项目名称
	Count   float64 //项目金额
	Tax     float64 //项目金额
	TaxRate float64 //税率
}

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
	Projects 
}

```

sample code
```
package main

import (
	"fmt"
	"github.com/hqbobo/cntax"
)

func main() {
	f, _ := cntax.NewTaxFile("f.ofd")
	fmt.Printf("%+v", *f)
	//parse the xml came with ofd
	fx, e := cntax.NewTaxEInvoice("f.xml")
	fmt.Printf("%+v", *fx)
}


```