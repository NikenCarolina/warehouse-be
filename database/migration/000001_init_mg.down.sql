-- Drop the transaction detail tables first
DROP TABLE IF EXISTS "TransaksiPenerimaanBarangDetail";
DROP TABLE IF EXISTS "TransaksiPengeluaranBarangDetail";

-- Drop the transaction header tables
DROP TABLE IF EXISTS "TransaksiPenerimaanBarangHeader";
DROP TABLE IF EXISTS "TransaksiPengeluaranBarangHeader";

-- Drop the master tables
DROP TABLE IF EXISTS "MasterWarehouse";
DROP TABLE IF EXISTS "MasterProduct";
DROP TABLE IF EXISTS "MasterCustomer";
DROP TABLE IF EXISTS "MasterSupplier";
