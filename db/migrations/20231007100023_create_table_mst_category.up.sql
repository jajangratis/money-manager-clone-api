CREATE TABLE "mst_category" (
                                "id" uuid DEFAULT uuid_generate_v4(),
                                "category_id" varchar(255),
                                "category_name" varchar(255),
                                "created_date" date NOT NULL DEFAULT CURRENT_DATE,
                                "updated_date" date NOT NULL DEFAULT CURRENT_DATE,
                                PRIMARY KEY("id")
);
