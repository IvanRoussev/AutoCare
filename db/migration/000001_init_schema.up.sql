CREATE TABLE "cars"
(
  "vin" VARCHAR(20) PRIMARY KEY,
  "username" VARCHAR(30) NOT NULL,
  "make" VARCHAR(20) NOT NULL,
  "model" VARCHAR(30) NOT NULL,
  "year" int NOT NULL
);
CREATE TABLE "maintenances"
(
  "maintenance_id" serial PRIMARY KEY,
  "car_vin" VARCHAR(20) NOT NULL,
  "maintenance_type" VARCHAR(30) NOT NULL,
  "mileage" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(NOW())
)
;
CREATE INDEX ON "cars" ("vin");
CREATE INDEX ON "maintenances" ("car_vin");
ALTER TABLE "maintenances" ADD FOREIGN KEY ("car_vin") REFERENCES "cars" ("vin")