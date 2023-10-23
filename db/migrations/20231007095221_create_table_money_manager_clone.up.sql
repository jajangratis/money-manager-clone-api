CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "mst_transaction_method" (
                                          "id" uuid DEFAULT uuid_generate_v4(),
                                          "method_id" varchar(255),
                                          "method_name" varchar(255),
                                          "created_date" date NOT NULL DEFAULT CURRENT_DATE,
                                          "updated_date" date NOT NULL DEFAULT CURRENT_DATE,
                                      Primary Key("id")
);