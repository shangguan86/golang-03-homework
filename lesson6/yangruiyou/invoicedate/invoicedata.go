package main

import (
	"time"
	"io"
	"fmt"
)

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
	Note       string
	Items      []*Item
}

const (
	fileType    = "INVOICES"
	magicNumber = 0x125D
	fileVersion = 100
	dataFormat  = "2018-6-5"
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
