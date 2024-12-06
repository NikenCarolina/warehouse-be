CREATE TABLE "MasterSupplier" (
    "SupplierPK" SERIAL PRIMARY KEY,
    "SupplierName" VARCHAR(255) NOT NULL
);

CREATE TABLE "MasterCustomer" (
"CustomerPK" SERIAL PRIMARY KEY,
"CustomerName" VARCHAR(255) NOT NULL
);

CREATE TABLE "MasterProduct" (
"ProductPK" SERIAL PRIMARY KEY,
"ProductName" VARCHAR(255) NOT NULL
);

CREATE TABLE "MasterWarehouse" (
"WhsPK" SERIAL PRIMARY KEY,
"WhsName" VARCHAR(255) NOT NULL
);

CREATE TABLE "TransaksiPenerimaanBarangHeader" (
"TrxInPK" SERIAL PRIMARY KEY,
"TrxInNo" VARCHAR(255) NOT NULL,
"WhsIdf" INT NOT NULL,
"TrxInDate" DATE NOT NULL,
"TrxInSuppIdf" INT NOT NULL,
"TrxInNotes" VARCHAR(255),
    FOREIGN KEY ("WhsIdf") REFERENCES "MasterWarehouse" ("WhsPK") ON DELETE CASCADE,
    FOREIGN KEY ("TrxInSuppIdf") REFERENCES "MasterSupplier" ("SupplierPK") ON DELETE CASCADE
);

CREATE TABLE "TransaksiPenerimaanBarangDetail" (
"TrxInDPK" SERIAL PRIMARY KEY,
"TrxInIDF" INT NOT NULL,
"TrxInDProductIdf" INT NOT NULL,
"TrxInDQtyDus" INT NOT NULL,
"TrxInDQtyPcs" INT NOT NULL,
    FOREIGN KEY ("TrxInIDF") REFERENCES "TransaksiPenerimaanBarangHeader" ("TrxInPK") ON DELETE CASCADE,
    FOREIGN KEY ("TrxInDProductIdf") REFERENCES "MasterProduct" ("ProductPK") ON DELETE CASCADE
);

CREATE TABLE "TransaksiPengeluaranBarangHeader" (
"TrxOutPK" SERIAL PRIMARY KEY,
"TrxOutNo" VARCHAR(255) NOT NULL,
"WhsIdf" INT NOT NULL,
"TrxOutDate" DATE NOT NULL,
"TrxOutSuppIdf" INT NOT NULL,
"TrxOutNotes" VARCHAR(255),
    FOREIGN KEY ("WhsIdf") REFERENCES "MasterWarehouse" ("WhsPK") ON DELETE CASCADE,
    FOREIGN KEY ("TrxOutSuppIdf") REFERENCES "MasterSupplier" ("SupplierPK") ON DELETE CASCADE
);

CREATE TABLE "TransaksiPengeluaranBarangDetail" (
"TrxOutDPK" SERIAL PRIMARY KEY,
"TrxOutIDF" INT NOT NULL,
"TrxOutDProductIdf" INT NOT NULL,
"TrxOutDQtyDus" INT NOT NULL,
"TrxOutDQtyPcs" INT NOT NULL,
    FOREIGN KEY ("TrxOutIDF") REFERENCES "TransaksiPengeluaranBarangHeader" ("TrxOutPK") ON DELETE CASCADE,
    FOREIGN KEY ("TrxOutDProductIdf") REFERENCES "MasterProduct" ("ProductPK") ON DELETE CASCADE
);
