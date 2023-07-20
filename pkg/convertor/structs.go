package convertor

import (
	"fmt"
	"strings"

	t "github.com/mehanizm/iuliia-go"
)

type DataCSV struct {
	ClientName        string `csv:"customer"`
	ClientEmail       string `csv:"email"`
	ClientAddress     string
	ClientCountry     string `csv:"country"`
	ClientStreetLine1 string `csv:"street_line_1"`
	ClientStreetLine2 string `csv:"street_line_2"`
	ClientCity        string `csv:"city"`
	ClientRegion      string `csv:"region"`
	ClientCode        string `csv:"postal_code"`

	Number      string `csv:"number"`
	Date        string `csv:"date"`
	PaymentDate string `csv:"payment_date"`
	Currency    string `csv:"accounting_currency"`
	TotalAmount string `csv:"total_amount"`
	NetAmount   string `csv:"net_amount"`
	VAT         string `csv:"tax_1_rate"`
	VATFull     string
	VATShort    string
	VATAmount   string `csv:"tax_1_amount"`
}

func (d *DataCSV) preRender() {
	d.ClientName = t.Wikipedia.Translate(d.ClientName)

	d.Date = strings.ReplaceAll(d.Date, "/", ".")
	d.Date = strings.ReplaceAll(d.Date, "-", ".")

	d.PaymentDate = strings.ReplaceAll(d.PaymentDate, "/", ".")
	d.PaymentDate = strings.ReplaceAll(d.PaymentDate, "-", ".")

	var address []string
	if d.ClientStreetLine1 != "" {
		address = append(address, d.ClientStreetLine1)
	}
	if d.ClientStreetLine2 != "" {
		address = append(address, d.ClientStreetLine2)
	}
	if d.ClientCity != "" {
		address = append(address, d.ClientCity)
	}
	if d.ClientRegion != "" {
		address = append(address, d.ClientRegion)
	}
	if d.ClientCode != "" {
		address = append(address, d.ClientCode)
	}

	d.ClientAddress = strings.Join(address, ", ")

	d.VATShort = strings.TrimSuffix(d.VAT, ".00%")
	d.VATFull = fmt.Sprintf("VAT (%s%%)", d.VATShort)
}
