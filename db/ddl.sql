CREATE DATABASE IF NOT EXISTS echoapp;
USE echoapp;
DROP TABLE IF EXISTS member;
CREATE TABLE IF NOT EXISTS member (
  number    INT PRIMARY KEY,
  name      VARCHAR(255) NOT NULL,
  position  CHAR(2) NOT NULL,
  created_at BIGINT       NOT NULL
);

