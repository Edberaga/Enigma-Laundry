package entity

import (
	"database/sql"
	"fmt"
	"time"
)

type Transaction struct {
	No int
	TanggalMasuk time.Time
	TanggalSelesai time.Time
	Pelayanan string
	Jumlah int
	Satuan string
	Harga float64
	Total float64
	Penerima string
}

func (t *Transaction) Validate(tx *sql.Tx) error {
	if t.TanggalMasuk.After(t.TanggalSelesai) {
		fmt.Println("Tanggal Masuk tidak bisa setelah Tanggal Selesai")
		tx.Rollback()
	}
	if t.Pelayanan == "" {
		fmt.Println("Pelayanan cannot be empty")
		tx.Rollback()
	}
	if t.Jumlah <= 0 {
		fmt.Println("Jumlah must be positive")
		tx.Rollback()
	}
	if t.Harga <= 0 {
		fmt.Println("Harga must be positive")
		tx.Rollback()
	}
	if t.Total != float64(t.Jumlah)*t.Harga {
		fmt.Println("Total must be equal to Jumlah * Harga")
		tx.Rollback()
	}
	return nil
}