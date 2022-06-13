package product

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func formatRupiah(num float64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("Rp%v", num)
}
