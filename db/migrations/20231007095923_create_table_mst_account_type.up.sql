CREATE TABLE "mst_account_type" (
                                    "id" uuid DEFAULT uuid_generate_v4(),
                                    "account_id" varchar(255),
                                    "account_name" varchar(255),
                                    "created_date" date NOT NULL DEFAULT CURRENT_DATE,
                                    "updated_date" date NOT NULL DEFAULT CURRENT_DATE,
                                PRIMARY KEY("id")
);