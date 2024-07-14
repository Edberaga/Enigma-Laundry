package function

import (
	"database/sql"
	en "enigma-laundry/entity"
	"fmt"
)

func PurchaseService(db *sql.DB, t en.Transaction) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	getServicePrice(&t, tx)
	InsertTransaction(t, tx)
	takenFee := getCustomerFee(t.Penerima, tx)
	updateCustomerFee(takenFee, t.Penerima, tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Succesfully insert for customer: ", t.Penerima)
	}
}

//SERVICE TRANSACTION: Retrieve satuan and Harga value according to service name, calculate the total and then update to tx table
func getServicePrice(t *en.Transaction, tx *sql.Tx) {
	querySatuan := "SELECT satuan FROM service WHERE nama = $1;"
	queryHarga := "SELECT harga FROM service WHERE nama = $1;"
	var satuan string = ""
	var harga float64 = 0

	err := tx.QueryRow(querySatuan, t.Pelayanan).Scan(&satuan)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error during retrievibng Satuan service : ", err, "||| Transaction has rollback!")
	} else {
		t.Satuan = satuan
		fmt.Println("Transaction successfully get service satuan", satuan, "Transaction satuan is: ", t.Satuan)
	}

	err = tx.QueryRow(queryHarga, t.Pelayanan).Scan(&harga)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error during retrievibng Satuan service : ", err, "||| Transaction has rollback!")
	} else {
		t.Harga = harga
		fmt.Println("Transaction successfully get service satuan and harga", harga, "Transaction harga is: ", t.Harga)
	}
}

//Insert Transction
func InsertTransaction(t en.Transaction, tx *sql.Tx) {
	fmt.Println("Current Harga: ", t.Harga)
	t.Total = t.Harga * float64(t.Jumlah)
	fmt.Println("Total: ", t.Total)
	if err := t.Validate(tx); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	query := "INSERT INTO tx_laundry (nomor, tanggal_masuk, tanggal_selesai, pelayanan, jumlah, satuan, harga, total, penerima) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := tx.Exec(query, t.No, t.TanggalMasuk, t.TanggalSelesai, t.Pelayanan, t.Jumlah, t.Satuan, t.Harga, t.Total, t.Penerima)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error during update customer fee: ", err, "||| Transaction has rollback!")
	} else {
		fmt.Println("Transaction successfully inserted")
	}
}

//CUSTOMER TRANSACTION #GetTotalFee, and update Total Fee
func getCustomerFee(name string, tx *sql.Tx) float64 {
	query := "SELECT total FROM tx_laundry WHERE penerima = $1;"
	takenFee := 0
	err := tx.QueryRow(query, name).Scan(&takenFee)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error during update customer fee: ", err, "||| Transaction has rollback!")
	}else {
		fmt.Println("Transaction successfully get customer fee", float64(takenFee))
	}
	return float64(takenFee)
}

func updateCustomerFee(newFee float64, name string, tx *sql.Tx) {
	var currentFee float64

	// Retrieve the current total_fee for the customer
	query := "SELECT total_fee FROM customer WHERE nama_customer = $1;"
	err := tx.QueryRow(query, name).Scan(&currentFee)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error retrieving current total fee:", err, "||| Transaction has rollback!")
		return
	}

	updatedFee := currentFee + newFee
	fmt.Println("Current Fee: ", currentFee)
	fmt.Println("New Fee: ", newFee)
	fmt.Println("Updated Fee: ", updatedFee)

	updateQuery := "UPDATE customer SET total_fee = $1 WHERE nama_customer = $2;"
	_, err = tx.Exec(updateQuery, updatedFee, name)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error during update customer fee:", err, "||| Transaction has rollback!")
	} else {
		fmt.Println("Transaction successfully updated customer fee")
	}
}

func ViewTransactions(db *sql.DB) {
	rows, err := db.Query("SELECT nomor, tanggal_masuk, tanggal_selesai, pelayanan, jumlah, satuan, harga, total, penerima FROM tx_laundry")
	if err != nil {
		fmt.Println("View error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var transaction en.Transaction
		if err := rows.Scan(&transaction.No, &transaction.TanggalMasuk, &transaction.TanggalSelesai, &transaction.Pelayanan, &transaction.Jumlah, &transaction.Satuan, &transaction.Harga, &transaction.Total, &transaction.Penerima); err != nil {
			fmt.Println("Scan error:", err)
		} else {
			fmt.Printf("No: %d, TanggalMasuk: %s, TanggalSelesai: %s, Pelayanan: %s, Jumlah: %d, Satuan: %s, Harga: %.2f, Total: %.2f, Penerima: %s\n",
				transaction.No, transaction.TanggalMasuk, transaction.TanggalSelesai, transaction.Pelayanan, transaction.Jumlah, transaction.Satuan, transaction.Harga, transaction.Total, transaction.Penerima)
		}
	}
}