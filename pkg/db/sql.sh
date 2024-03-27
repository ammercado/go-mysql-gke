# create database
create database crypto_db;

# use database
use crypto_db;

# insert into table songs
INSERT INTO coins (created_at, updated_at, name, abbreviation) VALUES (now(), now(), "Bitcoin", "BTC");
INSERT INTO coins (created_at, updated_at, name, abbreviation) VALUES (now(), now(), "Ethereum", "ETH");
INSERT INTO coins (created_at, updated_at, name, abbreviation) VALUES (now(), now(), "Tether", "USDT");
INSERT INTO coins (created_at, updated_at, name, abbreviation) VALUES (now(), now(), "BNB", "BNB");
INSERT INTO coins (created_at, updated_at, name, abbreviation) VALUES (now(), now(), "Solana", "SOL");
INSERT INTO coins (created_at, updated_at, name, abbreviation) VALUES (now(), now(), "XRP", "XRP);