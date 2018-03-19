CREATE DATABASE ssn owner postgres encoding 'utf8';
\c ssn;

/* Drop Tables */

DROP TABLE IF EXISTS "user";

/* Create Tables */

CREATE TABLE "user"
(
	id serial NOT NULL UNIQUE,
	user_name VARCHAR(32),
	email VARCHAR(256),
	password VARCHAR(256),
	full_name VARCHAR(256),
	birthday DATE,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE user_id_SEQ INCREMENT 1 RESTART 1;
