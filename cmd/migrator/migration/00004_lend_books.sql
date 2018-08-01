-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."lend_books" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "book_id" uuid,
  "user_id" uuid,
  "from" timestamptz,
  "to" timestamptz, 
  CONSTRAINT "lend_books_pkey" PRIMARY KEY ("id"),
  FOREIGN KEY ("book_id") REFERENCES "public"."books"("id"),
  FOREIGN KEY ("user_id") REFERENCES "public"."users"("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."lend_books"
