package report

import (
	"context"
	"errors"
	"log"
	"strconv"
)

const (
	InvoiceIndex      = 1
	NamaProdukIndex   = 8
	StatusIndex       = 3
	JumlahProdukIndex = 13
	HargaAwalIndex    = 14
	HargaJualIndex    = 16
)

const (
	CommissionNameIndex        = 0
	InvoiceCommissionIndex     = 3
	ProductNameCommissionIndex = 2
	ServiceFeeCommissionIndex  = 10
)

var (
	// KEEP IN MIND function extractValues that use the same index for mapping values
	headerUsed = map[int]string{
		InvoiceIndex:      "Nomor Invoice",        //B
		StatusIndex:       "Status Terakhir",      //D
		NamaProdukIndex:   "Nama Produk",          // I
		JumlahProdukIndex: "Jumlah Produk Dibeli", //N
		HargaAwalIndex:    "Harga Awal (IDR)",     //O
		HargaJualIndex:    "Harga Jual (IDR)",     //Q
	}

	// KEEP IN MIND function extractCommissionValues uses the same index for mapping values
	headerCommission = map[int]string{
		CommissionNameIndex:        "Commission Name",
		InvoiceCommissionIndex:     "Invoice No",
		ProductNameCommissionIndex: "Product Name",
		ServiceFeeCommissionIndex:  "Service Fee Gross",
	}
)

type SummaryReport struct {
	Header []string      `json:"header"`
	Value  []interface{} `json:"value"`
}

type TransactionReport struct {
	Invoice    string
	Status     string
	NamaProduk string
	Jumlah     int
	HargaAwal  int64
	HargaJual  int64
}

type InvoiceDetail struct {
	ProductName []string
	Status      string
}

type CommissionReport struct {
	Invoice         string
	NamaProduk      string
	CommissionName  string
	ServiceFeeGross int64
}

// Returns Header, and value of string array respectively
func (usecase *Usecase) CalculateMonthlySummaryReporting(ctx context.Context, path string, pathCommission string) (SummaryReport, error) {
	fileRaw, err := readFile(path)
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	firstSheet := fileRaw[0]
	// Header starts from row 4 based on generated reports
	err = validateHeader(firstSheet[4], headerUsed)
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	// extract values
	trxList, err := extractValues(firstSheet[5:])
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	summary, mapInvoices := calculateTransactions(trxList)

	if pathCommission == "" {
		return summary, nil
	}

	commissionFile, err := readFile(pathCommission)
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	// Header starts from row 0 based on generated reports
	err = validateHeader(commissionFile[0][0], headerCommission)
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	// Fetch the first sheet TOTAL
	commissionList, err := extractCommissionValues(commissionFile[0][1:])
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	commissionSummary := calculateCommissions(mapInvoices, commissionList)

	return SummaryReport{
		Header: append(summary.Header, commissionSummary.Header...),
		Value:  append(summary.Value, commissionSummary.Value...),
	}, nil
}

func calculateCommissions(mapInvoices map[string]InvoiceDetail, commissionList []CommissionReport) SummaryReport {
	var (
		serviceFee   int64
		deliveryFee  int64
		invoiceCount = map[string]InvoiceDetail{}
	)
	for _, commission := range commissionList {
		if _, ok := mapInvoices[commission.Invoice]; !ok {
			continue
		}
		switch commission.CommissionName {
		case "Biaya Layanan Power Merchant Pro":
			serviceFee += commission.ServiceFeeGross
		case "Biaya Layanan Bebas Ongkir Power Merchant Pro":
			deliveryFee += commission.ServiceFeeGross
		}
		// Check for possibility an invoice is not listed in commission list
		invoiceCount[commission.Invoice] = mapInvoices[commission.Invoice]
	}
	if len(invoiceCount) != len(mapInvoices) {
		log.Printf("Invoice count is not expected, Actual: %d Expected: %d \n", len(invoiceCount), len(mapInvoices))
		for invoice, detail := range mapInvoices {
			if _, ok := invoiceCount[invoice]; !ok {
				log.Printf("Invoice Missing %s Status %s\n", invoice, detail.Status)
			}
		}
	}
	return SummaryReport{
		Header: []string{"Service Fee", "Delivery Fee"},
		Value:  []interface{}{serviceFee, deliveryFee},
	}
}

func calculateTransactions(trxList []TransactionReport) (SummaryReport, map[string]InvoiceDetail) {
	mapInvoices := map[string]InvoiceDetail{}
	totalJumlah := 0
	totalHargaAwal := int64(0)
	TPV := int64(0)
	NMV := int64(0)
	for _, trx := range trxList {
		invoiceDetail := mapInvoices[trx.Invoice]
		invoiceDetail.ProductName = append(invoiceDetail.ProductName, trx.NamaProduk)
		invoiceDetail.Status = trx.Status
		mapInvoices[trx.Invoice] = invoiceDetail
		totalJumlah += trx.Jumlah

		// canceled order
		if trx.Status == "Dibatalkan Penjual" {
			continue
		}
		totalHargaAwal += trx.HargaAwal
		TPV += trx.HargaJual

		// finished order
		if trx.Status == "Pesanan Selesai" {
			NMV += trx.HargaJual
		}
	}
	return SummaryReport{
		Header: []string{"Total Order", "Total Item", "Total Purchase", "Gross Profit", "Potential Profit", "Discounted"},
		Value: []interface{}{
			len(mapInvoices),
			totalJumlah,
			TPV,
			NMV,
			TPV - NMV,
			totalHargaAwal - TPV,
		},
	}, mapInvoices
}

func extractCommissionValues(values [][]string) ([]CommissionReport, error) {
	commList := make([]CommissionReport, len(values))
	for index, row := range values {
		mapValues := map[int]string{}
		for indexMap := range headerCommission {
			mapValues[indexMap] = row[indexMap]
		}
		if mapValues[ServiceFeeCommissionIndex] == "" {
			mapValues[ServiceFeeCommissionIndex] = "0"
		}
		serviceFee, err := strconv.ParseInt(mapValues[ServiceFeeCommissionIndex], 10, 64)
		if err != nil {
			log.Printf("Row: %d, err: %v \n", index, err)
			return nil, errors.New("invalid service fee")
		}
		commList[index] = CommissionReport{
			Invoice:         mapValues[InvoiceCommissionIndex],
			NamaProduk:      mapValues[ProductNameCommissionIndex],
			CommissionName:  mapValues[CommissionNameIndex],
			ServiceFeeGross: serviceFee,
		}
	}
	return commList, nil
}

func extractValues(values [][]string) ([]TransactionReport, error) {
	trxList := make([]TransactionReport, len(values))
	for index, row := range values {
		mapValues := map[int]string{}
		for indexMap := range headerUsed {
			mapValues[indexMap] = row[indexMap]
		}
		jumlah, err := strconv.Atoi(mapValues[JumlahProdukIndex])
		if err != nil {
			log.Printf("Row: %d, err: %v \n", index, err)
			return nil, errors.New("invalid jumlah value")
		}
		hargaAwal, err := strconv.ParseInt(mapValues[HargaAwalIndex], 10, 64)
		if err != nil {
			log.Printf("Row: %d, err: %v \n", index, err)
			return nil, errors.New("invalid harga awal value")
		}
		hargaJual, err := strconv.ParseInt(mapValues[HargaJualIndex], 10, 64)
		if err != nil {
			log.Printf("Row: %d, err: %v \n", index, err)
			return nil, errors.New("invalid harga jual value")
		}
		trxList[index] = TransactionReport{
			Invoice:    mapValues[InvoiceIndex],
			Status:     mapValues[StatusIndex],
			NamaProduk: mapValues[NamaProdukIndex],
			Jumlah:     jumlah,
			HargaAwal:  hargaAwal,
			HargaJual:  hargaJual,
		}
	}
	return trxList, nil
}

func validateHeader(headers []string, headerMap map[int]string) error {
	for index, header := range headerMap {
		if index >= len(headers) {
			return errors.New("invalid index header")
		}
		if actualHeader := headers[index]; actualHeader != header {
			log.Printf("index: %d, header: %s \n", index, actualHeader)
			return errors.New("header is not valid")
		}
	}
	return nil
}
