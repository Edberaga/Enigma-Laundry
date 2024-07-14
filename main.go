package main

import (
	"bufio"
	"database/sql"
	en "enigma-laundry/entity"
	fun "enigma-laundry/function"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Edberaga7"
	dbname   = "enigma_laundry_challenge"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	db := connectDb()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. View Customers")
		fmt.Println("2. Insert Customer")
		fmt.Println("3. Update Customer")
		fmt.Println("4. Delete Customer")
		fmt.Println("5. View Services")
		fmt.Println("6. Insert Service")
		fmt.Println("7. Update Service")
		fmt.Println("8. Delete Service")
		fmt.Println("9. View Transactions")
		fmt.Println("10. Insert Transaction")
		fmt.Println("0. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			
			fun.ViewCustomers(db)
		case 2:
			var name string
			var phoneNo string
			fmt.Println("Enter name:")
			fmt.Scan(&name)
			fmt.Println("Enter phone number:")
			fmt.Scan(&phoneNo)
			fun.InsertCustomer(db, en.Customer{Name: name, PhoneNo: phoneNo})
		case 3:
			var name string
			var phoneNo string
			fmt.Println("Enter name:")
			fmt.Scan(&name)
			fmt.Println("Enter phone number:")
			fmt.Scan(&phoneNo)
			fun.UpdateCustomer(db, en.Customer{Name: name, PhoneNo: phoneNo})
		case 4:
			var name string
			fmt.Println("Enter name:")
			fmt.Scan(&name)
			fun.DeleteCustomer(db, name)
		case 5:
			fun.ViewServices(db)
		case 6:
			var name string
			var satuan string
			var price float64

			// Loop until we get a valid service name
			for {
				fmt.Println("Enter service name:")
				scanner.Scan()
				name = scanner.Text()
				if name != "" {
					break
				}
			}

			// Loop until we get a valid satuan type
			for {
				fmt.Println("Enter Satuan type: ")
				scanner.Scan()
				satuan = scanner.Text()
				if satuan != "" {
					break
				}
			}

			fmt.Println("Enter price:")
			readPrice, _ := reader.ReadString('\n')
			price, _ = strconv.ParseFloat(strings.TrimSpace(readPrice), 64)

			fun.InsertService(db, en.Service{ServiceName: name, Satuan: satuan, Price: price})
		case 7:
			var serviceToUpdate string
			var name string
			var satuan string
			var price float64

			for {
				fmt.Println("Enter the Service name you wish to Update:")
				scanner.Scan()
				serviceToUpdate= scanner.Text()
				if serviceToUpdate != "" {
					break
				}
			}

			for {
				fmt.Println("Enter service name:")
				scanner.Scan()
				name = scanner.Text()
				if name != "" {
					break
				}
			}

			for {
				fmt.Println("Enter Satuan type: ")
				scanner.Scan()
				satuan = scanner.Text()
				if satuan != "" {
					break
				}
			}

			fmt.Println("Enter price:")
			readPrice, _ := reader.ReadString('\n')
			price, _ = strconv.ParseFloat(strings.TrimSpace(readPrice), 64)

			fun.UpdateService(db, serviceToUpdate, en.Service{ServiceName: name, Satuan: satuan, Price: price})
		case 8:
			var name string
			for {
				fmt.Println("Enter service name:")
				scanner.Scan()
				name = scanner.Text()
				if name != "" {
					break
				}
			}
			fun.DeleteService(db, name)
		case 9:
			fun.ViewTransactions(db)
		case 10:
			var no, jumlah int
			var pelayanan,penerima string
			var tanggalMasuk, tanggalSelesai string
			fmt.Println("Enter no:")
			fmt.Scan(&no)
			fmt.Println("Enter tanggal masuk (YYYY-MM-DD):")
			fmt.Scan(&tanggalMasuk)
			fmt.Println("Enter tanggal selesai (YYYY-MM-DD):")
			fmt.Scan(&tanggalSelesai)

			for {
				fmt.Println("Enter service name:")
				scanner.Scan()
				pelayanan = scanner.Text()
				if pelayanan != "" {
					break
				}
			}
			fmt.Println("Enter jumlah:")
			fmt.Scan(&jumlah)
			// total = float64(jumlah) * harga
			fmt.Println("Enter penerima:")
			fmt.Scan(&penerima)

			tMasuk, _ := time.Parse("2006-01-02", tanggalMasuk)
			tSelesai, _ := time.Parse("2006-01-02", tanggalSelesai)

			fun.PurchaseService(db, en.Transaction{No: no, TanggalMasuk: tMasuk, TanggalSelesai: tSelesai, Pelayanan: pelayanan, Jumlah: jumlah, Penerima: penerima})
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
		pause()
	}
}

func pause() {
	fmt.Println("Press any Key to continue... (not enter)")
	key := ""
	for {
		fmt.Scan(&key)
		if key != "" {
			break
		}
	}
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database")
	return db
}