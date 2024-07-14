/*CUSTOMER QUERY*/

/*CUSTOMER DUMMY DATA*/
INSERT INTO customer (nama_customer, no_HP, total_fee) VALUES ('Edbert', '082169672989', 0);
INSERT INTO customer (nama_customer, no_HP, total_fee) VALUES ('Stanly', '082169672989', 0);
INSERT INTO customer (nama_customer, no_HP, total_fee) VALUES ('Andres', '082169672989', 0);

/*CUSTOMER VIEW QUERY*/
SELECT * FROM customer;
SELECT * FROM customer WHERE nama_customer = 'Edbert';

/*CUSTOMER INSERT QUERY*/
INSERT INTO customer (nama_customer, no_HP, total_fee) VALUES ('Ray', '082169672989', 0);

/*CUSTOMER UPDATE QUERY*/
UPDATE customer SET no_HP = '081354843800' WHERE nama_customer = 'Andre';

/*CUSTOMER DELETE QUERY*/
DELETE FROM customer WHERE nama_customer = 'Stanly';

/*SERVICE QUERY*/

/*SERVICE DUMMY DATA */
INSERT INTO service (nama, satuan, harga) VALUES ('Cuci Baju', 'KG', 7000);
INSERT INTO service (nama, satuan, harga) VALUES ('Laundry Bedcover', 'Buah', 5000);
INSERT INTO service (nama, satuan, harga) VALUES ('Laundry Repair', 'Buah', 2000);

/*CUSTOMER VIEW QUERY*/
SELECT * FROM service;
SELECT * FROM service WHERE nama = 'Cuci Baju';

/*CUSTOMER INSERT QUERY*/
INSERT INTO service (nama, satuan, harga) VALUES('Laundry Uniform', 'Buah', 3000);

/*CUSTOMER UPDATE QUERY*/
UPDATE service SET nama = 'Cuci Baju + Celana', satuan = 'KG', harga = '7500' WHERE nama = 'Cuci Baju';

/*CUSTOMER DELETE QUERY*/
DELETE FROM service WHERE nama = 'Laundry Repair';