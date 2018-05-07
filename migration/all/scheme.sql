CREATE DATABASE ssn owner postgres encoding 'utf8';
\c ssn;

/* Drop Tables */
DROP TABLE IF EXISTS "star_counts";
DROP TABLE IF EXISTS "post_hashtags";
DROP TABLE IF EXISTS "posts_stars";
DROP TABLE IF EXISTS "team_players";
DROP TABLE IF EXISTS "tournament_teams";
DROP TABLE IF EXISTS "hashtags";
DROP TABLE IF EXISTS "images";
DROP TABLE IF EXISTS "videos";
DROP TABLE IF EXISTS "posts";
DROP TABLE IF EXISTS "players";
DROP TABLE IF EXISTS "masters";
DROP TABLE IF EXISTS "matchs";
DROP TABLE IF EXISTS "teams";
DROP TABLE IF EXISTS "tournaments";
DROP TABLE IF EXISTS "user_follows";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "locations";

/* Create Tables */

CREATE TABLE "users"
(
	id serial NOT NULL UNIQUE,
	user_name VARCHAR(32) UNIQUE,
	email VARCHAR(256) UNIQUE,
	password VARCHAR(256),
	first_name VARCHAR(256),
	last_name VARCHAR(256),
	city VARCHAR(256),
	country VARCHAR(256),
	about VARCHAR(500),
	quote VARCHAR(256),
	birthday DATE DEFAULT NULL,
	role VARCHAR(50) DEFAULT 'user',
	score int DEFAULT 0,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE users_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "user_follows"
(
	id serial NOT NULL UNIQUE,
	user_id int NOT NULL,
	follow_id int NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id),
	CONSTRAINT user_follows_user_id_follow_id_UNIQUE UNIQUE (user_id, follow_id)
)WITHOUT OIDS;

ALTER SEQUENCE user_follows_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "posts"
(
	id serial NOT NULL UNIQUE,
	user_id int NOT NULL,
	caption text,
	location_id int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE posts_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "locations"
(
	id serial NOT NULL UNIQUE,
	place_id varchar UNIQUE,
	post_count int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE locations_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "hashtags"
(
	id serial NOT NULL UNIQUE,
	"key_word" varchar UNIQUE,
	post_count int DEFAULT 0,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE hashtags_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "post_hashtags"
(
	id serial NOT NULL UNIQUE,
	post_id int,
	hashtag_id int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id),
	CONSTRAINT post_hashtags_post_id_hashtag_id_UNIQUE UNIQUE (post_id, hashtag_id)
) WITHOUT OIDS;

ALTER SEQUENCE post_hashtags_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "images"
(
	id serial NOT NULL UNIQUE,
	post_id int NOT NULL,
	name varchar(272),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE images_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "videos"
(
	id serial NOT NULL UNIQUE,
	post_id int NOT NULL,
	name varchar(272),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE videos_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "post_stars"
(
	id serial NOT NULL UNIQUE,
	user_id int NOT NULL,
	post_id int NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id),
	CONSTRAINT post_stars_post_id_user_id_UNIQUE UNIQUE (post_id, user_id)
) WITHOUT OIDS;

ALTER SEQUENCE post_stars_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "star_counts"
(
	id serial NOT NULL UNIQUE,
	owner_id int,
	owner_type text,
	quantity int DEFAULT 0,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id),
	CONSTRAINT star_counts_owner_id_owner_type_UNIQUE UNIQUE (owner_id, owner_type)
) WITHOUT OIDS;

ALTER SEQUENCE star_counts_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "teams"
(
	id serial NOT NULL UNIQUE,
	name VARCHAR(256) NOT NULL,
	description text,
	max_members int DEFAULT 16,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE teams_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "team_players"
(
	id serial NOT NULL UNIQUE,
	user_id int NOT NULL,
	team_id int NOT NULL,
	position VARCHAR(10) NOT NULL, 
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id),
	CONSTRAINT team_players_user_id_team_id_UNIQUE UNIQUE (user_id, team_id)
) WITHOUT OIDS;

ALTER SEQUENCE team_players_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "matches"
(
	id serial NOT NULL UNIQUE,
	tournament_id int,
	description text,
	start_date timestamp NOT NULL,
	team1_id int NOT NULL,
	team2_id int NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE matches_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "tournaments"
(
	id serial NOT NULL UNIQUE,
	name varchar(256) NOT NULL,
	description text,
	start_date timestamp NOT NULL,
	end_date timestamp NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id)
) WITHOUT OIDS;

ALTER SEQUENCE tournaments_id_SEQ INCREMENT 1 RESTART 1;


CREATE TABLE "tournament_teams"
(
	id serial NOT NULL UNIQUE,
	tournament_id int NOT NULL,
	team_id int NOT NULL,
	score int DEFAULT 0,
	"group" VARCHAR(256),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id),
	CONSTRAINT tournament_teams_tournament_id_team_id_UNIQUE UNIQUE (tournament_id, team_id)
) WITHOUT OIDS;

ALTER SEQUENCE tournament_teams_id_SEQ INCREMENT 1 RESTART 1;

CREATE TABLE "masters"
(
	id serial NOT NULL UNIQUE,
	owner_id int,
	owner_type text,
	user_id int NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id),
	CONSTRAINT masters_owner_id_owner_type_user_id_UNIQUE UNIQUE (owner_id, owner_type, user_id)
) WITHOUT OIDS;

ALTER SEQUENCE masters_id_SEQ INCREMENT 1 RESTART 1;

ALTER TABLE user_follows
	ADD FOREIGN KEY (user_id)
	REFERENCES "users" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE user_follows
	ADD FOREIGN KEY (follow_id)
	REFERENCES "users" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE posts
	ADD FOREIGN KEY (user_id)
	REFERENCES "users" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE posts
	ADD FOREIGN KEY (location_id)
	REFERENCES "locations" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE post_hashtags
	ADD FOREIGN KEY (post_id)
	REFERENCES "posts" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE post_hashtags
	ADD FOREIGN KEY (hashtag_id)
	REFERENCES "hashtags" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE post_stars
	ADD FOREIGN KEY (post_id)
	REFERENCES "posts" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE post_stars
	ADD FOREIGN KEY (user_id)
	REFERENCES "users" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE team_players
	ADD FOREIGN KEY (user_id)
	REFERENCES "users" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE team_players
	ADD FOREIGN KEY (team_id)
	REFERENCES "teams" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE matches
	ADD FOREIGN KEY (tournament_id)
	REFERENCES "tournaments" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE matches
	ADD FOREIGN KEY (team1_id)
	REFERENCES "teams" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE matches
	ADD FOREIGN KEY (team2_id)
	REFERENCES "teams" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE masters
	ADD FOREIGN KEY (user_id)
	REFERENCES "users" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE tournament_teams
	ADD FOREIGN KEY (tournament_id)
	REFERENCES "tournaments" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;

ALTER TABLE tournament_teams
	ADD FOREIGN KEY (team_id)
	REFERENCES "teams" (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;