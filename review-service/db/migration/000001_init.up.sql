CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Required for gen_random_uuid()

CREATE TABLE "review_messages" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "review" varchar NOT NULL,
  "score" float NOT NULL,
  "label" varchar NOT NULL,
  "review_id" uuid NOT NULL,
  "updated_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "reviews" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "product_id" varchar UNIQUE NOT NULL,
  "num_of_reviews" float NOT NULL
);