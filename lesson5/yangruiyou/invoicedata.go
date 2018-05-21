package main

import (
	"time"
	"io"
	"fmt"
	"os"
	"log"
)

func main() {


	var file *os.File
	var err error
	if file,err=os.Open(filename);err!=nil{
		log.Println("fail to open the file",err)
	}
	defer file.Close()


}

type Item struct {
	Id       string
	Price    float64
	Quantity int
	Note     string
}
type Invoice struct {
	Id         int
	CustomerId int
	Raised     time.Time
	Due        time.Time
	Paid       bool
	Note       []*Item
}

const (
	fileType    = "INVOICES"
	magicNumber = 0x125D
	fileVersion = 100
	dataFormat  = "2016-01-02"
)

type InvoiceMarshaler interface {
	MarshalInvoices(writer io.Writer, invoices []*Invoice) error
}

type InvoiceUnmarshaler interface {
	UnmarshalInvoices(reader io.Reader) ([]*Invoice, error)
}

func readInvoices(reader io.Reader, suffix string) ([]*Invoice, error) {
	var unmarshaler InvoiceUnmarshaler
	switch suffix {
	case ".gob":
		unmarshaler = GobMarshaler{}
	case ".inv":
		unmarshaler = InvMarshaler{}
	case ".jsn", ".json":
		unmarshaler = JSONMarshaler{}
	case ".txt":
		unmarshaler = TxtMarshaler{}
	case ".xml":
		unmarshaler = XMLMarshaler{}

	}
	if unmarshaler != nil {
		return unmarshaler.UnmarshalInvoices(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix:%s", suffix)
}


type JSONMarshaler struct{}


