CREATE TABLE "tr_transaction_history" (
                                          "id" uuid DEFAULT uuid_generate_v4 (),
                                          "method_id" varchar(255),
                                          "account_id" varchar(255),
                                          "category_id" varchar(255),
                                          "from" varchar(255),
                                          "to" varchar(255),
                                          "amount" numeric(255),
                                          "fee" numeric(255),
                                          "note" text,
                                          "created_date" date NOT NULL DEFAULT CURRENT_DATE,
                                          "updated_date" date NOT NULL DEFAULT CURRENT_DATE,
                                          PRIMARY KEY("id")
);