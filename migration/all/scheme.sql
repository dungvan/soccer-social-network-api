CREATE DATABASE circle owner postgres encoding 'utf8';
\c circle;



/* Drop Tables */

DROP TABLE IF EXISTS collection_item;
DROP TABLE IF EXISTS collection_outfit;
DROP TABLE IF EXISTS collection;
DROP TABLE IF EXISTS detection_item;
DROP TABLE IF EXISTS rel_outfit_hashtag;
DROP TABLE IF EXISTS hashtag;
DROP TABLE IF EXISTS item_star_count;
DROP TABLE IF EXISTS post_item_star_history;
DROP TABLE IF EXISTS user_item_report;
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS mst_selection;
DROP TABLE IF EXISTS outfit_comment_count;
DROP TABLE IF EXISTS outfit_ranking;
DROP TABLE IF EXISTS outfit_star_count;
DROP TABLE IF EXISTS post_comment_history;
DROP TABLE IF EXISTS post_star_history;
DROP TABLE IF EXISTS user_outfit_report;
DROP TABLE IF EXISTS outfit;
DROP TABLE IF EXISTS outfit_location;
DROP TABLE IF EXISTS user_follow;
DROP TABLE IF EXISTS user_option_explore;
DROP TABLE IF EXISTS user_ranking;
DROP TABLE IF EXISTS visual_search_history;
DROP TABLE IF EXISTS user_app;
DROP TABLE IF EXISTS hashtag;




/* Create Tables */

CREATE TABLE collection
(
	id_collection serial NOT NULL UNIQUE,
	id_user_app int NOT NULL UNIQUE,
	collection_name varchar,
	-- False : Collection is public,
	-- True: Collection is private.
	private_flag boolean,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_collection)
) WITHOUT OIDS;


CREATE TABLE collection_item
(
	id_collection_item serial NOT NULL UNIQUE,
	id_collection int NOT NULL UNIQUE,
	id_item int NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_collection_item)
) WITHOUT OIDS;


CREATE TABLE collection_outfit
(
	id_collection_item serial NOT NULL UNIQUE,
	id_collection int NOT NULL UNIQUE,
	id_outfit int NOT NULL UNIQUE,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_collection_item)
) WITHOUT OIDS;


CREATE TABLE detection_item
(
	id_detection_item serial NOT NULL UNIQUE,
	-- Key related to table "Outfit"
	id_outfit int NOT NULL,
	id_item int NOT NULL,
	coordinate_position box NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_detection_item)
) WITHOUT OIDS;


ALTER SEQUENCE detection_item_id_detection_item_SEQ INCREMENT 1 RESTART 1;


CREATE TABLE hashtag
(
	id_hashtag serial NOT NULL UNIQUE,
	hashtag_key_word varchar(99) NOT NULL UNIQUE,
	post_count int DEFAULT 0 NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_hashtag)
) WITHOUT OIDS;


CREATE TABLE item
(
	id_item serial NOT NULL UNIQUE,
	im_name varchar(128) NOT NULL UNIQUE,
	brand_name varchar(5000) NOT NULL,
	product_name varchar(5000) NOT NULL,
	source_image_url varchar(256) NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_item)
) WITHOUT OIDS;

ALTER SEQUENCE item_id_item_SEQ INCREMENT 1 RESTART 1;


CREATE TABLE item_star_count
(
	id_item_star_count serial NOT NULL UNIQUE,
	id_item int NOT NULL,
	star_count int DEFAULT 0 NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_item_star_count)
) WITHOUT OIDS;


CREATE TABLE mst_selection
(
	id_mst_selection serial NOT NULL,
	-- コード区分
	code_type varchar(50) NOT NULL,
	code_order int NOT NULL,
	code_value varchar(2000),
	PRIMARY KEY (id_mst_selection)
) WITHOUT OIDS;


-- ユーザがアップロードした着こなし情報を格納する
CREATE TABLE outfit
(
	id_outfit serial NOT NULL UNIQUE,
	-- Key related to table "User"
	id_user_app int NOT NULL,
	-- 文字数制限はフロント要件を確認中のため仮の値です。
	caption varchar(5000),
	source_image_file_name varchar(256),
	id_location int,
	outfit_post_datetime timestamp,
	-- False : Outfit is public,
	-- True: Outfit is private.
	private_flag boolean DEFAULT 'false' NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_outfit)
) WITHOUT OIDS;


ALTER SEQUENCE outfit_id_outfit_SEQ INCREMENT 1 RESTART 1;


CREATE TABLE outfit_comment_count
(
	id_outfit_comment_count int NOT NULL UNIQUE,
	id_outfit int NOT NULL UNIQUE,
	comment_count int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_outfit_comment_count)
) WITHOUT OIDS;


CREATE TABLE outfit_location
(
	id_location serial NOT NULL UNIQUE,
	place_id varchar,
	post_count int,
	PRIMARY KEY (id_location)
) WITHOUT OIDS;


CREATE TABLE outfit_ranking
(
	id_outfit_ranking serial NOT NULL UNIQUE,
	id_outfit int NOT NULL UNIQUE,
	gender varchar(5),
	star_count_at_day int,
	star_count_at_week int,
	star_count_at_season int,
	star_count_at_all int,
	-- 桁数確認中
	location_continent varchar,
	location_country varchar,
	location_city varchar,
	location_prefecture varchar,
	location_neighbourhood varchar,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_outfit_ranking)
) WITHOUT OIDS;


-- 着こなしごとのいいねされた数を格納する
CREATE TABLE outfit_star_count
(
	id_outfit_star_count serial NOT NULL UNIQUE,
	-- Key related to table "Outfit"
	id_outfit int NOT NULL,
	star_count int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_outfit_star_count)
) WITHOUT OIDS;


ALTER SEQUENCE outfit_star_count_id_outfit_star_count_SEQ INCREMENT 1 RESTART 1;


CREATE TABLE post_comment_history
(
	id_post_comment_history serial NOT NULL UNIQUE,
	id_outfit int NOT NULL UNIQUE,
	id_user_app int NOT NULL UNIQUE,
	comment_detail varchar,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_post_comment_history)
) WITHOUT OIDS;


-- ユーザが着こなしにいいねをした履歴を格納
CREATE TABLE post_item_star_history
(
	id_post_item_star_history serial NOT NULL,
	tap_star_datetime timestamp,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	id_item int NOT NULL,
	id_user_app int NOT NULL UNIQUE,
	PRIMARY KEY (id_post_item_star_history)
) WITHOUT OIDS;


-- ユーザが着こなしにいいねをした履歴を格納
CREATE TABLE post_star_history
(
	id_post_star_history serial NOT NULL UNIQUE,
	-- Key related to table "Outfit"
	id_outfit int NOT NULL,
	-- Key related to table "User"
	id_user_app int NOT NULL,
	tap_star_datetime timestamp,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_post_star_history)
) WITHOUT OIDS;


ALTER SEQUENCE post_star_history_id_post_star_history_SEQ INCREMENT 1 RESTART 1;


-- hashtagテーブルとの結合条件はN:Nとする
CREATE TABLE rel_outfit_hashtag
(
	id_rel_outfit_hashtag serial NOT NULL UNIQUE,
	id_outfit int NOT NULL,
	id_hashtag int NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_rel_outfit_hashtag),
	CONSTRAINT key_outfit_hashtag UNIQUE (id_outfit, id_hashtag)
) WITHOUT OIDS;


-- サービスを利用しているユーザの情報を格納する
CREATE TABLE user_app
(
	id_user_app serial NOT NULL UNIQUE,
	uuid char(36) NOT NULL UNIQUE,
	-- 文字数制限はユーザ登録要件未確定のため仮の値です。
	user_name varchar(20),
	user_profile_image_url varchar(256),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_app)
) WITHOUT OIDS;


ALTER SEQUENCE user_app_id_user_app_SEQ INCREMENT 1 RESTART 1;


CREATE TABLE user_follow
(
	id_user_follow serial NOT NULL UNIQUE,
	id_user_app int NOT NULL UNIQUE,
	id_following_user int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_follow)
) WITHOUT OIDS;


CREATE TABLE user_item_report
(
	id_user_item_report serial NOT NULL,
	id_user_app int NOT NULL UNIQUE,
	id_item int NOT NULL,
	-- 選択結果（mst_selectionの表示に対応）
	code_order int DEFAULT -1 NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_item_report)
) WITHOUT OIDS;


CREATE TABLE user_option_explore
(
	id_user_option_explore serial NOT NULL UNIQUE,
	id_user_app int NOT NULL UNIQUE,
	option_location varchar,
	option_recency int,
	option_gender int,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_option_explore)
) WITHOUT OIDS;


CREATE TABLE user_outfit_report
(
	id_user_outfit_report int NOT NULL,
	id_user_app int NOT NULL,
	id_outfit int NOT NULL,
	-- 選択結果（mst_selectionの表示に対応）
	code_order int DEFAULT -1 NOT NULL,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_outfit_report)
) WITHOUT OIDS;


CREATE TABLE user_ranking
(
	id_user_ranking serial NOT NULL UNIQUE,
	id_user_app int NOT NULL UNIQUE,
	gender varchar(5),
	star_count_at_season int,
	star_count_at_day int,
	star_count_at_week int,
	star_count_at_all int,
	follower_count int DEFAULT 0 NOT NULL,
	-- 桁数確認中
	location_continent varchar,
	location_country varchar,
	location_city varchar,
	location_prefecture varchar,
	location_neighbourhood varchar,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_user_ranking)
) WITHOUT OIDS;


-- ユーザがVisualSearchをした履歴情報を格納する
CREATE TABLE visual_search_history
(
	id_visual_search_history serial NOT NULL UNIQUE,
	-- Key related to table "User"
	id_user_app int NOT NULL,
	source_image_url varchar(256),
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY (id_visual_search_history)
) WITHOUT OIDS;


ALTER SEQUENCE visual_search_history_id_visual_search_history_SEQ INCREMENT 1 RESTART 1;

/* Create Foreign Keys */

ALTER TABLE collection_item
	ADD FOREIGN KEY (id_collection)
	REFERENCES collection (id_collection)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE collection_outfit
	ADD FOREIGN KEY (id_collection)
	REFERENCES collection (id_collection)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE rel_outfit_hashtag
	ADD FOREIGN KEY (id_hashtag)
	REFERENCES hashtag (id_hashtag)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE collection_item
	ADD FOREIGN KEY (id_item)
	REFERENCES item (id_item)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE detection_item
	ADD FOREIGN KEY (id_item)
	REFERENCES item (id_item)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE item_star_count
	ADD FOREIGN KEY (id_item)
	REFERENCES item (id_item)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE post_item_star_history
	ADD FOREIGN KEY (id_item)
	REFERENCES item (id_item)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_item_report
	ADD FOREIGN KEY (id_item)
	REFERENCES item (id_item)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE collection_outfit
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE detection_item
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE outfit_comment_count
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE outfit_ranking
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE outfit_star_count
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE post_comment_history
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE post_star_history
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE rel_outfit_hashtag
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_outfit_report
	ADD FOREIGN KEY (id_outfit)
	REFERENCES outfit (id_outfit)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE outfit
	ADD FOREIGN KEY (id_location)
	REFERENCES outfit_location (id_location)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE collection
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE outfit
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE post_comment_history
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE post_item_star_history
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE post_star_history
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_item_report
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_option_explore
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_outfit_report
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_ranking
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE visual_search_history
	ADD FOREIGN KEY (id_user_app)
	REFERENCES user_app (id_user_app)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;



/* Comments */

COMMENT ON COLUMN collection.private_flag IS 'False : Collection is public,
True: Collection is private.';
COMMENT ON COLUMN detection_item.id_outfit IS 'Key related to table "Outfit"';
COMMENT ON COLUMN detection_item.id_item IS 'Key related to table "Item"';
COMMENT ON COLUMN mst_selection.code_type IS 'コード区分';
COMMENT ON TABLE outfit IS 'ユーザがアップロードした着こなし情報を格納する';
COMMENT ON COLUMN outfit.id_user_app IS 'Key related to table "User"';
COMMENT ON COLUMN outfit.caption IS '文字数制限はフロント要件を確認中のため仮の値です。';
COMMENT ON COLUMN outfit.private_flag IS 'False : Outfit is public,
True: Outfit is private.';
COMMENT ON COLUMN outfit_ranking.location_continent IS '桁数確認中';
COMMENT ON TABLE outfit_star_count IS '着こなしごとのいいねされた数を格納する';
COMMENT ON COLUMN outfit_star_count.id_outfit IS 'Key related to table "Outfit"';
COMMENT ON TABLE post_item_star_history IS 'ユーザが着こなしにいいねをした履歴を格納';
COMMENT ON TABLE post_star_history IS 'ユーザが着こなしにいいねをした履歴を格納';
COMMENT ON COLUMN post_star_history.id_outfit IS 'Key related to table "Outfit"';
COMMENT ON COLUMN post_star_history.id_user_app IS 'Key related to table "User"';
COMMENT ON TABLE rel_outfit_hashtag IS 'hashtagテーブルとの結合条件はN:Nとする';
COMMENT ON TABLE user_app IS 'サービスを利用しているユーザの情報を格納する';
COMMENT ON COLUMN user_app.user_name IS '文字数制限はユーザ登録要件未確定のため仮の値です。';
COMMENT ON COLUMN user_item_report.code_order IS '選択結果（mst_selectionの表示に対応）';
COMMENT ON COLUMN user_outfit_report.code_order IS '選択結果（mst_selectionの表示に対応）';
COMMENT ON COLUMN user_ranking.location_continent IS '桁数確認中';
COMMENT ON TABLE visual_search_history IS 'ユーザがVisualSearchをした履歴情報を格納する';
COMMENT ON COLUMN visual_search_history.id_user_app IS 'Key related to table "User"';
