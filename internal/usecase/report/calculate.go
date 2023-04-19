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

// Returns Header, and value of string array respectively
func (usecase *Usecase) CalculateMonthlySummaryReporting(ctx context.Context, path string) (SummaryReport, error) {
	fileRaw, err := readFile(path)
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	firstSheet := fileRaw[0]
	// Header starts from row 4 based on generated reports
	err = validateHeader(firstSheet[4])
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

	summary, err := calculateTransactions(trxList)
	if err != nil {
		log.Println(err)
		return SummaryReport{}, err
	}

	return summary, nil
}

func calculateTransactions(trxList []TransactionReport) (SummaryReport, error) {
	mapInvoices := map[string][]string{}
	totalJumlah := 0
	totalHargaAwal := int64(0)
	TPV := int64(0)
	NMV := int64(0)
	for _, trx := range trxList {
		mapInvoices[trx.Invoice] = append(mapInvoices[trx.Invoice], trx.NamaProduk)
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
	}, nil
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

func validateHeader(headers []string) error {
	for index, header := range headerUsed {
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
