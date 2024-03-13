CREATE TABLE IF NOT EXISTS kelas (
  id_kelas varchar PRIMARY KEY NOT NULL,
  walas int,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
