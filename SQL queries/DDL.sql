CREATE DATABASE enigma_laundry_challenge;

CREATE TABLE customer (
	nama_customer VARCHAR(100) NOT NULL,
	no_HP VARCHAR(15) NOT NULL,
	PRIMARY KEY(nama_customer)
);

CREATE TABLE service (
	nama VARCHAR(100) NOT NULL,
	satuan VARCHAR(10) NOT NULL,
	harga DECIMAL NOT NULL,
	PRIMARY KEY(nama)
);

CREATE TABLE tx_laundry (
	nomor INT NOT NULL,
	tanggal_masuk DATE NOT NULL,
	tanggal_selesai DATE NOT NULL,
	pelayanan VARCHAR(100) NOT NULL,
	jumlah INT NOT NULL,
	satuan VARCHAR(10) NOT NULL,
	harga DECIMAL NOT NULL,
	total DECIMAL NOT NULL,
	penerima VARCHAR(100) NOT NULL,
	PRIMARY KEY(nomor),
	FOREIGN KEY(pelayanan) REFERENCES service(nama),
	FOREIGN KEY(penerima) REFERENCES customer(nama_customer)
);