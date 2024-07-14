package function

import (
	"database/sql"
	en "enigma-laundry/entity"
	"fmt"
)

//Customer function
func InsertCustomer(db *sql.DB, customer en.Customer) {
	if err := customer.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	query := "INSERT INTO customer (nama_customer, no_HP, total_fee) VALUES ($1, $2, 0)"
	_, err := db.Exec(query, customer.Name, customer.PhoneNo)
	if err != nil {
		fmt.Println("Insert error:", err)
	} else {
		fmt.Println("Customer successfully inserted")
	}
}

func ViewCustomers(db *sql.DB) {
	rows, err := db.Query("SELECT nama_customer, no_HP, total_fee FROM customer")
	if err != nil {
		fmt.Println("View error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var customer en.Customer
		if err := rows.Scan(&customer.Name, &customer.PhoneNo, &customer.TotalFee); err != nil {
			fmt.Println("Scan error:", err)
		} else {
			fmt.Printf("Name: %s, Phone: %s\n, Total: %2.f\n", customer.Name, customer.PhoneNo, customer.TotalFee)
		}
	}
}

func UpdateCustomer(db *sql.DB, customer en.Customer) {
	if err := customer.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	query := "UPDATE customer SET no_HP = $1 WHERE nama_customer = $2"
	_, err := db.Exec(query, customer.PhoneNo, customer.Name)
	if err != nil {
		fmt.Println("Update error:", err)
	} else {
		fmt.Println("Customer successfully updated")
	}
}

func DeleteCustomer(db *sql.DB, name string) {
	query := "DELETE FROM customer WHERE nama_customer = $1"
	_, err := db.Exec(query, name)
	if err != nil {
		fmt.Println("Delete error:", err)
	} else {
		fmt.Println("Customer successfully deleted")
	}
}