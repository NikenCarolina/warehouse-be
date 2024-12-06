INSERT INTO "MasterSupplier" ("SupplierName")
VALUES 
    ('Supplier A'),
    ('Supplier B'),
    ('Supplier C');

INSERT INTO "MasterCustomer" ("CustomerName")
VALUES 
    ('Customer X'),
    ('Customer Y'),
    ('Customer Z');

INSERT INTO "MasterProduct" ("ProductName")
VALUES 
    ('Product 1'),
    ('Product 2'),
    ('Product 3');

INSERT INTO "MasterWarehouse" ("WhsName")
VALUES 
    ('Warehouse North'),
    ('Warehouse South'),
    ('Warehouse East');

INSERT INTO "TransaksiPenerimaanBarangHeader" ("TrxInNo", "WhsIdf", "TrxInDate", "TrxInSuppIdf", "TrxInNotes")
VALUES 
    ('TRX-IN-001', 1, '2024-12-01', 1, 'First batch of items'),
    ('TRX-IN-002', 2, '2024-12-02', 2, 'Second batch of items'),
    ('TRX-IN-003', 3, '2024-12-03', 3, 'Third batch of items');

INSERT INTO "TransaksiPenerimaanBarangDetail" ("TrxInIDF", "TrxInDProductIdf", "TrxInDQtyDus", "TrxInDQtyPcs")
VALUES 
    (1, 1, 10, 100),
    (1, 2, 5, 50),
    (2, 2, 7, 70),
    (2, 3, 3, 30),
    (3, 1, 15, 150),
    (3, 3, 10, 100);

INSERT INTO "TransaksiPengeluaranBarangHeader" ("TrxOutNo", "WhsIdf", "TrxOutDate", "TrxOutSuppIdf", "TrxOutNotes")
VALUES 
    ('TRX-OUT-001', 1, '2024-12-05', 1, 'Dispatch to Customer A'),
    ('TRX-OUT-002', 2, '2024-12-06', 2, 'Dispatch to Customer B'),
    ('TRX-OUT-003', 3, '2024-12-07', 3, 'Dispatch to Customer C');

INSERT INTO "TransaksiPengeluaranBarangDetail" ("TrxOutIDF", "TrxOutDProductIdf", "TrxOutDQtyDus", "TrxOutDQtyPcs")
VALUES 
    (1, 1, 3, 30),
    (1, 2, 2, 20),
    (2, 2, 4, 40),
    (2, 3, 1, 10),
    (3, 1, 5, 50),
    (3, 3, 6, 60);
