CREATE DATABASE ssn owner postgres encoding 'utf8';
\c ssn;

/* Drop Tables */

DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "post";

/* Create Tables */

CREATE TABLE "user"
(
	id serial NOT NULL UNIQUE,
	user_name VARCHAR(32) UNIQUE,
	email VARCHAR(256) UNIQUE,
	password VARCHAR(256),
	full_name VARCHAR(256),
	birthday DATE,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE user_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE post
(
	id serial NOT NULL UNIQUE,
	user_id int NOT NULL,
	caption text,
	source_image_file_name varchar(272),
	source_video_file_name varchar(256),
	location_id int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE post_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "location"
(
	id serial NOT NULL UNIQUE,
	place_id varchar UNIQUE,
	post_count int,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE location_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE hashtag
(
	id serial NOT NULL UNIQUE,
	"key" varchar UNIQUE,
	post_count int,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE hashtag_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE post_hashtags
(
	id serial NOT NULL UNIQUE,
	post_id int,
	hashtag_id int,
	PRIMARY KEY (id),
	CONSTRAINT post_hashtags_post_id_hashtag_id_UNIQUE UNIQUE (post_id, hashtag_id)
) WITHOUT OIDS;

ALTER SEQUENCE post_hashtags_id_SEQ INCREMENT 1 RESTART 1;

ALTER TABLE post
	ADD FOREIGN KEY (user_id)
	REFERENCES "user" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE post
	ADD FOREIGN KEY (location_id)
	REFERENCES "location" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE post_hashtags
	ADD FOREIGN KEY (post_id)
	REFERENCES "post" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE post_hashtags
	ADD FOREIGN KEY (hashtag_id)
	REFERENCES "hashtag" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;
