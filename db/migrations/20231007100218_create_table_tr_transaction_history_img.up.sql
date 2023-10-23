CREATE TABLE "tr_transaction_history_img" (
                                              "id" uuid DEFAULT uuid_generate_v4 (),
                                              "tr_transaction_history_id" uuid,
                                              "img_link" text,
                                              "created_date" date NOT NULL DEFAULT CURRENT_DATE,
                                              "updated_date" date NOT NULL DEFAULT CURRENT_DATE,
                                              PRIMARY KEY("id")
);
