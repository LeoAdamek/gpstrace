CREATE TABLE "locations" (
  "id" SERIAL PRIMARY KEY,
  "entity" CHARACTER VARYING(64),
  "latitude" DOUBLE PRECISION,
  "longitude" DOUBLE PRECISION,
  "timestamp" TIMESTAMP WITH TIME ZONE,
  "ingest_timestamp" TIMESTAMP NOT NULL
);