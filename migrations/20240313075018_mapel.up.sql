CREATE TABLE IF NOT EXISTS mapel (
  id varchar PRIMARY KEY NOT NULL,
  judul varchar NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
