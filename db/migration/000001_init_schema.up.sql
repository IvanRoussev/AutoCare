CREATE TABLE "owners"
(
  "id" bigserial PRIMARY KEY,
  "first_name" VARCHAR(20) NOT NULL,
  "last_name" VARCHAR(20) NOT NULL,
  "country" VARCHAR(20) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(NOW())
);

CREATE TABLE "cars"
(
    "vin" VARCHAR(20) PRIMARY KEY,
    "owner_id" int NOT NULL ,
    "make" VARCHAR(20) NOT NULL,
    "model" VARCHAR(30) NOT NULL,
    "year" int NOT NULL
);


CREATE TABLE "maintenances"
(
  "maintenance_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "car_vin" VARCHAR(20) NOT NULL ,
  "maintenance_type" VARCHAR(30) NOT NULL ,
  "mileage" int NOT NULL ,
  "created_at" timestamptz NOT NULL DEFAULT(NOW())
);


CREATE INDEX ON "cars" ("vin");
CREATE INDEX ON "maintenances" ("car_vin");
ALTER TABLE "cars" ADD FOREIGN KEY ("owner_id") REFERENCES "owners" ("id");
ALTER TABLE "maintenances" ADD FOREIGN KEY ("car_vin") REFERENCES "cars" ("vin")