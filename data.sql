--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3 (Debian 10.3-1.pgdg90+1)
-- Dumped by pg_dump version 10.3 (Debian 10.3-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: comment_stars; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.comment_stars (
    id integer NOT NULL,
    user_id integer NOT NULL,
    comment_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: comment_stars_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.comment_stars_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: comment_stars_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.comment_stars_id_seq OWNED BY public.comment_stars.id;


--
-- Name: comments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.comments (
    id integer NOT NULL,
    post_id integer NOT NULL,
    user_id integer NOT NULL,
    content text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- Name: hashtags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hashtags (
    id integer NOT NULL,
    key_word character varying,
    post_count integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: hashtags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.hashtags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: hashtags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.hashtags_id_seq OWNED BY public.hashtags.id;


--
-- Name: images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.images (
    id integer NOT NULL,
    post_id integer NOT NULL,
    name character varying(272),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.images_id_seq OWNED BY public.images.id;


--
-- Name: locations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.locations (
    id integer NOT NULL,
    place_id character varying,
    post_count integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: locations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.locations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.locations_id_seq OWNED BY public.locations.id;


--
-- Name: masters; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.masters (
    id integer NOT NULL,
    owner_id integer,
    owner_type text,
    user_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: masters_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.masters_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: masters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.masters_id_seq OWNED BY public.masters.id;


--
-- Name: matches; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.matches (
    id integer NOT NULL,
    tournament_id integer,
    description text,
    start_date timestamp without time zone NOT NULL,
    team1_id integer NOT NULL,
    team2_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    team1_goals integer,
    team2_goals integer
);


--
-- Name: matches_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.matches_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: matches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.matches_id_seq OWNED BY public.matches.id;


--
-- Name: post_hashtags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.post_hashtags (
    id integer NOT NULL,
    post_id integer,
    hashtag_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: post_hashtags_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.post_hashtags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: post_hashtags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.post_hashtags_id_seq OWNED BY public.post_hashtags.id;


--
-- Name: post_stars; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.post_stars (
    id integer NOT NULL,
    user_id integer NOT NULL,
    post_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: post_stars_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.post_stars_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: post_stars_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.post_stars_id_seq OWNED BY public.post_stars.id;


--
-- Name: posts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.posts (
    id integer NOT NULL,
    user_id integer NOT NULL,
    caption text,
    location_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    type character varying(20) DEFAULT 'status'::character varying NOT NULL
);


--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.posts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: star_counts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.star_counts (
    id integer NOT NULL,
    owner_id integer,
    owner_type text,
    quantity integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: star_counts_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.star_counts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: star_counts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.star_counts_id_seq OWNED BY public.star_counts.id;


--
-- Name: team_players; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.team_players (
    id integer NOT NULL,
    user_id integer NOT NULL,
    team_id integer NOT NULL,
    "position" character varying(10) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: team_players_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.team_players_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: team_players_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.team_players_id_seq OWNED BY public.team_players.id;


--
-- Name: teams; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.teams (
    id integer NOT NULL,
    name character varying(256) NOT NULL,
    description text,
    max_members integer DEFAULT 16,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: teams_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.teams_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.teams_id_seq OWNED BY public.teams.id;


--
-- Name: tournament_teams; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tournament_teams (
    id integer NOT NULL,
    tournament_id integer NOT NULL,
    team_id integer NOT NULL,
    score integer DEFAULT 0,
    "group" character varying(256),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: tournament_teams_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.tournament_teams_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tournament_teams_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.tournament_teams_id_seq OWNED BY public.tournament_teams.id;


--
-- Name: tournaments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tournaments (
    id integer NOT NULL,
    name character varying(256) NOT NULL,
    description text,
    start_date timestamp without time zone NOT NULL,
    end_date timestamp without time zone NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: tournaments_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.tournaments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: tournaments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.tournaments_id_seq OWNED BY public.tournaments.id;


--
-- Name: user_follows; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_follows (
    id integer NOT NULL,
    user_id integer NOT NULL,
    follow_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: user_follows_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_follows_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_follows_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_follows_id_seq OWNED BY public.user_follows.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    user_name character varying(32),
    email character varying(256),
    password character varying(256),
    first_name character varying(256),
    last_name character varying(256),
    city character varying(256),
    country character varying(256),
    about character varying(500),
    quote character varying(256),
    birthday date,
    role character varying(50) DEFAULT 'user'::character varying,
    score integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: videos; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.videos (
    id integer NOT NULL,
    post_id integer NOT NULL,
    name character varying(272),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


--
-- Name: videos_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.videos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: videos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.videos_id_seq OWNED BY public.videos.id;


--
-- Name: comment_stars id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comment_stars ALTER COLUMN id SET DEFAULT nextval('public.comment_stars_id_seq'::regclass);


--
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- Name: hashtags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hashtags ALTER COLUMN id SET DEFAULT nextval('public.hashtags_id_seq'::regclass);


--
-- Name: images id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.images ALTER COLUMN id SET DEFAULT nextval('public.images_id_seq'::regclass);


--
-- Name: locations id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.locations ALTER COLUMN id SET DEFAULT nextval('public.locations_id_seq'::regclass);


--
-- Name: masters id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.masters ALTER COLUMN id SET DEFAULT nextval('public.masters_id_seq'::regclass);


--
-- Name: matches id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matches ALTER COLUMN id SET DEFAULT nextval('public.matches_id_seq'::regclass);


--
-- Name: post_hashtags id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_hashtags ALTER COLUMN id SET DEFAULT nextval('public.post_hashtags_id_seq'::regclass);


--
-- Name: post_stars id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_stars ALTER COLUMN id SET DEFAULT nextval('public.post_stars_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Name: star_counts id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.star_counts ALTER COLUMN id SET DEFAULT nextval('public.star_counts_id_seq'::regclass);


--
-- Name: team_players id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.team_players ALTER COLUMN id SET DEFAULT nextval('public.team_players_id_seq'::regclass);


--
-- Name: teams id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.teams ALTER COLUMN id SET DEFAULT nextval('public.teams_id_seq'::regclass);


--
-- Name: tournament_teams id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tournament_teams ALTER COLUMN id SET DEFAULT nextval('public.tournament_teams_id_seq'::regclass);


--
-- Name: tournaments id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tournaments ALTER COLUMN id SET DEFAULT nextval('public.tournaments_id_seq'::regclass);


--
-- Name: user_follows id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_follows ALTER COLUMN id SET DEFAULT nextval('public.user_follows_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: videos id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.videos ALTER COLUMN id SET DEFAULT nextval('public.videos_id_seq'::regclass);


--
-- Data for Name: comment_stars; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.comment_stars (id, user_id, comment_id, created_at, updated_at, deleted_at) FROM stdin;
5	5	76	2018-05-20 19:48:54.486214	2018-05-20 19:48:54.486214	2018-05-20 19:49:06.052153
3	5	74	2018-05-20 19:37:39.275885	2018-05-20 19:51:59.309933	2018-05-20 19:51:59.684341
4	5	73	2018-05-20 19:37:40.146787	2018-05-20 19:52:11.075067	2018-05-20 19:52:18.094015
6	5	77	2018-05-21 01:13:08.506402	2018-05-21 01:13:08.506402	\N
7	1	84	2018-05-23 04:01:00.789582	2018-05-23 04:01:00.789582	\N
9	1	86	2018-05-23 04:01:22.949468	2018-05-23 04:01:22.949468	\N
10	1	80	2018-05-23 04:01:33.259737	2018-05-23 04:01:33.259737	\N
11	1	78	2018-05-23 04:01:43.969172	2018-05-23 04:01:43.969172	\N
12	1	90	2018-05-23 04:01:47.167177	2018-05-23 04:01:47.167177	\N
13	1	82	2018-05-26 17:11:46.311034	2018-05-26 17:11:46.311034	\N
8	1	85	2018-05-23 04:01:13.811051	2018-05-26 17:11:49.546932	2018-05-26 17:11:49.71672
\.


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.comments (id, post_id, user_id, content, created_at, updated_at, deleted_at) FROM stdin;
73	30	5	comment 	2018-05-20 19:33:57.87844	2018-05-20 19:33:57.87844	2018-05-20 19:56:07.273911
74	30	5	haha	2018-05-20 19:34:39.666363	2018-05-20 19:34:39.666363	2018-05-20 19:57:25.581712
75	30	5	abv	2018-05-20 19:39:20.357535	2018-05-20 19:39:20.357535	2018-05-20 20:00:34.460562
76	30	5	bvvv	2018-05-20 19:44:12.453791	2018-05-20 19:44:12.453791	2018-05-20 20:00:36.288269
77	31	5	test	2018-05-21 01:13:06.435558	2018-05-21 01:13:06.435558	\N
78	33	5	test	2018-05-21 01:47:42.306402	2018-05-21 01:47:42.306402	\N
80	33	1	asdasdzc	2018-05-23 03:18:27.600605	2018-05-23 03:18:27.600605	2018-05-26 17:11:44.318621
81	33	1	asdasdda	2018-05-23 03:18:38.47496	2018-05-23 03:18:38.47496	2018-05-26 17:11:45.24345
82	33	1	asdasdadsad	2018-05-23 03:19:14.473902	2018-05-23 03:19:14.473902	2018-05-26 17:11:47.083026
83	33	1	abc	2018-05-23 03:27:10.857277	2018-05-23 03:27:10.857277	2018-05-26 17:11:48.20123
84	33	1	asd	2018-05-23 03:29:06.335177	2018-05-23 03:29:06.335177	2018-05-26 17:11:48.871718
85	33	1	hahaha	2018-05-23 03:54:03.989754	2018-05-23 03:54:03.989754	2018-05-26 17:11:50.557264
86	33	1	hahaha	2018-05-23 03:56:40.815569	2018-05-23 03:56:40.815569	2018-05-26 17:11:50.733957
90	33	1	sao ha	2018-05-23 04:00:44.13575	2018-05-23 04:00:44.13575	2018-05-26 17:11:51.718662
79	32	1	test	2018-05-21 01:47:55.350768	2018-05-21 01:47:55.350768	2018-05-26 17:12:01.414882
87	32	1	hahaha	2018-05-23 03:56:48.979803	2018-05-23 03:56:48.979803	2018-05-26 17:12:03.18235
88	32	1	???	2018-05-23 03:58:18.095955	2018-05-23 03:58:18.095955	2018-05-26 17:12:04.01991
89	32	1	hahaha	2018-05-23 04:00:40.6291	2018-05-23 04:00:40.6291	2018-05-26 17:12:04.806473
91	34	1	haha	2018-05-27 16:03:51.540244	2018-05-27 16:03:51.540244	\N
92	46	1	test	2018-05-29 06:10:38.856036	2018-05-29 06:10:38.856036	\N
\.


--
-- Data for Name: hashtags; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.hashtags (id, key_word, post_count, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.images (id, post_id, name, created_at, updated_at, deleted_at) FROM stdin;
1	1	9a1a748024c50f75864a8162e822c2de1526043611.jpg	2018-05-11 13:00:41.815605	2018-05-11 13:00:41.815605	\N
3	5	0a43823a1429d30db0ff2c2c228241911526814047.png	2018-05-20 11:00:52.017524	2018-05-20 11:00:52.017524	\N
4	5	5bea826a3d776d934c74ad1f1c5d435f1526814047.png	2018-05-20 11:00:52.021636	2018-05-20 11:00:52.021636	\N
5	6	c71e4146b71c21a3fc7036df363dc39a1526814123.png	2018-05-20 11:02:07.94427	2018-05-20 11:02:07.94427	\N
6	7	c71e4146b71c21a3fc7036df363dc39a1526814123.png	2018-05-20 11:02:09.30692	2018-05-20 11:02:09.30692	\N
7	8	c71e4146b71c21a3fc7036df363dc39a1526814123.png	2018-05-20 11:02:25.381219	2018-05-20 11:02:25.381219	\N
8	9	dddaf631c478b57302b059537470ebcc1526814227.png	2018-05-20 11:03:50.93468	2018-05-20 11:03:50.93468	\N
9	10	fa3a0118ccf98a6011f377ffa790b78c1526814270.png	2018-05-20 11:04:37.933272	2018-05-20 11:04:37.933272	\N
10	22	f5b5d462ddc7e28c4c2ad137015382b11526815112.png	2018-05-20 11:18:36.767525	2018-05-20 11:18:36.767525	\N
13	25	e8423b86dabda32f2b171e7fd12ef39e1526818107.png	2018-05-20 12:08:36.060268	2018-05-20 12:08:36.060268	2018-05-20 17:03:28.921561
14	25	84a0a83dcbf1a4d52ecef5deff3cea2a1526818107.png	2018-05-20 12:08:36.062413	2018-05-20 12:08:36.062413	2018-05-20 17:03:28.921561
12	24	d1059a0adf3da7c0d8148ed0bf5b3d471526817161.png	2018-05-20 11:52:46.897501	2018-05-20 11:52:46.897501	2018-05-20 17:03:49.018682
11	23	946eaff07d100890cdaafd2fa0ac76611526816119.png	2018-05-20 11:35:27.358974	2018-05-20 11:35:27.358974	2018-05-20 17:03:57.978858
2	2	9a1a748024c50f75864a8162e822c2de1526043611.jpg	2018-05-11 16:36:58.651166	2018-05-11 16:36:58.651166	2018-05-20 17:06:39.972059
15	28	473dcf7e31f1d2fcc97d2c53db901aac1526840175.png	2018-05-20 18:16:19.457087	2018-05-20 18:16:19.457087	2018-05-20 18:35:58.256969
16	32	538d8b495064671eabe6f0c920ae21a21526866927.png	2018-05-21 01:42:27.158015	2018-05-21 01:42:27.158015	2018-05-26 17:18:44.123891
17	32	a8ea9a5f4119a77461cfb1ffe705f9a61526866927.png	2018-05-21 01:42:27.160833	2018-05-21 01:42:27.160833	2018-05-26 17:18:44.123891
18	32	5276474353f4b5f03ab609e66ff8b3fe1526866927.png	2018-05-21 01:42:27.162344	2018-05-21 01:42:27.162344	2018-05-26 17:18:44.123891
\.


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.locations (id, place_id, post_count, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: masters; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.masters (id, owner_id, owner_type, user_id, created_at, updated_at, deleted_at) FROM stdin;
2	2	teams	1	2018-05-11 14:38:21.361972	2018-05-11 14:38:21.361972	\N
5	1	tournaments	1	2018-05-11 15:05:48.276021	2018-05-11 15:05:48.276021	\N
6	1	matches	1	2018-05-11 15:08:57.075898	2018-05-11 15:08:57.075898	\N
7	2	matches	1	2018-05-11 15:09:19.411308	2018-05-11 15:09:19.411308	\N
3	3	teams	1	2018-05-11 14:38:43.066454	2018-05-11 14:38:43.066454	\N
4	4	teams	1	2018-05-11 14:39:08.077581	2018-05-11 14:39:08.077581	\N
1	1	teams	1	2018-05-11 14:37:54.715426	2018-05-11 14:37:54.715426	\N
\.


--
-- Data for Name: matches; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.matches (id, tournament_id, description, start_date, team1_id, team2_id, created_at, updated_at, deleted_at, team1_goals, team2_goals) FROM stdin;
2	1	Match 1 group B	2018-05-05 15:04:05	2	4	2018-05-11 15:09:19.409343	2018-05-11 15:09:19.409343	\N	1	0
1	1	Match 1 group A	2018-05-05 15:04:05	1	3	2018-05-11 15:08:57.070876	2018-05-11 15:08:57.070876	\N	1	1
\.


--
-- Data for Name: post_hashtags; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.post_hashtags (id, post_id, hashtag_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: post_stars; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.post_stars (id, user_id, post_id, created_at, updated_at, deleted_at) FROM stdin;
5	5	30	2018-05-20 19:48:46.075967	2018-05-20 19:52:58.986871	\N
6	5	31	2018-05-21 01:13:09.979111	2018-05-21 01:13:09.979111	\N
7	5	33	2018-05-21 01:47:50.380517	2018-05-21 01:47:50.380517	\N
9	1	33	2018-05-23 03:27:31.096557	2018-05-23 04:01:30.830004	\N
8	1	32	2018-05-21 01:52:53.668761	2018-05-23 04:01:49.551723	\N
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.posts (id, user_id, caption, location_id, created_at, updated_at, deleted_at, type) FROM stdin;
27	1	test	\N	2018-05-20 18:15:56.951368	2018-05-20 18:15:56.951368	2018-05-20 18:35:51.98347	talent-wanted
29	1	haha	\N	2018-05-20 18:20:52.196081	2018-05-20 18:20:52.196081	2018-05-20 18:35:55.050789	status
28	1	test test1	\N	2018-05-20 18:16:19.4534	2018-05-20 18:16:19.4534	2018-05-20 18:35:58.253523	status
30	5	tìm gì cho sáng nay?	\N	2018-05-20 19:11:16.050016	2018-05-20 19:11:16.050016	2018-05-20 20:00:42.296715	status
31	5	test	\N	2018-05-21 01:12:02.934336	2018-05-21 01:12:02.934336	2018-05-21 01:14:19.320406	status
33	1	tin thu mon ...... 	\N	2018-05-21 01:47:21.741603	2018-05-21 01:47:21.741603	2018-05-26 17:18:36.10482	talent-wanted
32	5	test	\N	2018-05-21 01:42:27.155065	2018-05-21 01:42:27.155065	2018-05-26 17:18:44.120648	status
34	1	this test\n	\N	2018-05-27 15:52:48.935549	2018-05-27 15:52:48.935549	2018-05-29 03:45:18.264744	status
35	1	wtf	\N	2018-05-29 05:48:07.935064	2018-05-29 05:48:07.935064	2018-05-29 05:48:43.978153	status
36	1	test	\N	2018-05-29 05:49:43.143629	2018-05-29 05:49:43.143629	2018-05-29 05:49:54.767863	status
38	1	asdasdsad	\N	2018-05-29 05:52:04.344845	2018-05-29 05:52:04.344845	2018-05-29 05:53:42.128145	status
37	1	asdsadas	\N	2018-05-29 05:50:39.421701	2018-05-29 05:50:39.421701	2018-05-29 05:53:45.248282	status
45	1	asdaszxcgegrw	\N	2018-05-29 06:07:37.815173	2018-05-29 06:07:37.815173	2018-05-29 06:10:47.857908	status
44	1	asdasdzxczx	\N	2018-05-29 06:05:38.181566	2018-05-29 06:05:38.181566	2018-05-29 06:10:51.628513	status
43	1	sadsadzxcxz\n	\N	2018-05-29 06:04:22.975841	2018-05-29 06:04:22.975841	2018-05-29 06:10:53.909439	status
42	1	asdasdzxczx	\N	2018-05-29 06:02:41.231968	2018-05-29 06:02:41.231968	2018-05-29 06:10:57.011001	status
41	1	asdasdas	\N	2018-05-29 06:02:16.169694	2018-05-29 06:02:16.169694	2018-05-29 06:10:59.240213	status
40	1	asdasd	\N	2018-05-29 06:00:15.698258	2018-05-29 06:00:15.698258	2018-05-29 06:11:02.371773	status
39	1	asdads	\N	2018-05-29 05:57:07.407755	2018-05-29 05:57:07.407755	2018-05-29 06:11:04.902423	status
46	1	asdasdxzcxzcwrewr2313	\N	2018-05-29 06:10:35.351979	2018-05-29 06:10:35.351979	2018-05-29 06:11:07.507322	status
\.


--
-- Data for Name: star_counts; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.star_counts (id, owner_id, owner_type, quantity, created_at, updated_at, deleted_at) FROM stdin;
99	27	posts	0	2018-05-20 18:15:56.955743	2018-05-20 18:15:56.955743	\N
100	28	posts	0	2018-05-20 18:16:19.455363	2018-05-20 18:16:19.455363	\N
101	29	posts	0	2018-05-20 18:20:52.198266	2018-05-20 18:20:52.198266	\N
105	75	comments	0	2018-05-20 19:39:20.360948	2018-05-20 19:39:20.360948	\N
106	76	comments	0	2018-05-20 19:44:12.455762	2018-05-20 19:49:06.045652	\N
104	74	comments	0	2018-05-20 19:34:39.670741	2018-05-20 19:51:59.683318	\N
103	73	comments	0	2018-05-20 19:33:57.882935	2018-05-20 19:52:18.092891	\N
102	30	posts	1	2018-05-20 19:11:16.054689	2018-05-20 19:52:58.990654	\N
108	77	comments	1	2018-05-21 01:13:06.445557	2018-05-21 01:13:08.512531	\N
107	31	posts	1	2018-05-21 01:12:02.938418	2018-05-21 01:13:09.983492	\N
112	79	comments	0	2018-05-21 01:47:55.352175	2018-05-21 01:47:55.352175	\N
114	81	comments	0	2018-05-23 03:18:38.4762	2018-05-23 03:18:38.4762	\N
116	83	comments	0	2018-05-23 03:27:10.858456	2018-05-23 03:27:10.858456	\N
120	87	comments	0	2018-05-23 03:56:48.980952	2018-05-23 03:56:48.980952	\N
121	88	comments	0	2018-05-23 03:58:18.097304	2018-05-23 03:58:18.097304	\N
122	89	comments	0	2018-05-23 04:00:40.630138	2018-05-23 04:00:40.630138	\N
117	84	comments	1	2018-05-23 03:29:06.337287	2018-05-23 04:01:00.798799	\N
119	86	comments	1	2018-05-23 03:56:40.816557	2018-05-23 04:01:22.954691	\N
110	33	posts	2	2018-05-21 01:47:21.744771	2018-05-23 04:01:30.834074	\N
113	80	comments	1	2018-05-23 03:18:27.60609	2018-05-23 04:01:33.269009	\N
111	78	comments	1	2018-05-21 01:47:42.308205	2018-05-23 04:01:43.973214	\N
123	90	comments	1	2018-05-23 04:00:44.136905	2018-05-23 04:01:47.174699	\N
109	32	posts	1	2018-05-21 01:42:27.156646	2018-05-23 04:01:49.554684	\N
115	82	comments	1	2018-05-23 03:19:14.476035	2018-05-26 17:11:46.321393	\N
118	85	comments	0	2018-05-23 03:54:03.991753	2018-05-26 17:11:49.715572	\N
124	34	posts	0	2018-05-27 15:52:48.942957	2018-05-27 15:52:48.942957	\N
125	91	comments	0	2018-05-27 16:03:51.545498	2018-05-27 16:03:51.545498	\N
126	35	posts	0	2018-05-29 05:48:07.939232	2018-05-29 05:48:07.939232	\N
127	36	posts	0	2018-05-29 05:49:43.146449	2018-05-29 05:49:43.146449	\N
128	37	posts	0	2018-05-29 05:50:39.423129	2018-05-29 05:50:39.423129	\N
129	38	posts	0	2018-05-29 05:52:04.346956	2018-05-29 05:52:04.346956	\N
130	39	posts	0	2018-05-29 05:57:07.409423	2018-05-29 05:57:07.409423	\N
131	40	posts	0	2018-05-29 06:00:15.700667	2018-05-29 06:00:15.700667	\N
132	41	posts	0	2018-05-29 06:02:16.171642	2018-05-29 06:02:16.171642	\N
133	42	posts	0	2018-05-29 06:02:41.234967	2018-05-29 06:02:41.234967	\N
134	43	posts	0	2018-05-29 06:04:22.977985	2018-05-29 06:04:22.977985	\N
135	44	posts	0	2018-05-29 06:05:38.182826	2018-05-29 06:05:38.182826	\N
136	45	posts	0	2018-05-29 06:07:37.81614	2018-05-29 06:07:37.81614	\N
137	46	posts	0	2018-05-29 06:10:35.352844	2018-05-29 06:10:35.352844	\N
138	92	comments	0	2018-05-29 06:10:38.866443	2018-05-29 06:10:38.866443	\N
\.


--
-- Data for Name: team_players; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.team_players (id, user_id, team_id, "position", created_at, updated_at, deleted_at) FROM stdin;
6	6	2	gk	2018-05-11 14:38:21.36442	2018-05-11 14:38:21.364421	\N
7	7	2	any	2018-05-11 14:38:21.364423	2018-05-11 14:38:21.364424	\N
8	8	2	any	2018-05-11 14:38:21.364425	2018-05-11 14:38:21.364425	\N
9	9	2	any	2018-05-11 14:38:21.364426	2018-05-11 14:38:21.364426	\N
10	10	2	any	2018-05-11 14:38:21.364427	2018-05-11 14:38:21.364428	\N
11	11	3	gk	2018-05-11 14:38:43.067574	2018-05-11 14:38:43.067574	\N
12	12	3	any	2018-05-11 14:38:43.067576	2018-05-11 14:38:43.067576	\N
13	13	3	any	2018-05-11 14:38:43.067577	2018-05-11 14:38:43.067577	\N
14	14	3	any	2018-05-11 14:38:43.067578	2018-05-11 14:38:43.067578	\N
15	15	3	any	2018-05-11 14:38:43.067579	2018-05-11 14:38:43.067579	\N
25	21	3	any	2018-05-11 14:53:20.860453	2018-05-11 14:53:20.860453	\N
16	16	4	gk	2018-05-11 14:39:08.078659	2018-05-11 14:39:08.078659	\N
17	17	4	any	2018-05-11 14:39:08.078661	2018-05-11 14:39:08.078661	\N
18	18	4	any	2018-05-11 14:39:08.078662	2018-05-11 14:39:08.078662	\N
19	19	4	any	2018-05-11 14:39:08.078663	2018-05-11 14:39:08.078663	\N
20	20	4	any	2018-05-11 14:39:08.078663	2018-05-11 14:39:08.078664	\N
1	1	1	gk	2018-05-11 14:37:54.719037	2018-05-11 14:37:54.719037	\N
2	2	1	any	2018-05-11 14:37:54.719039	2018-05-11 14:37:54.719039	\N
3	3	1	any	2018-05-11 14:37:54.719041	2018-05-11 14:37:54.719041	\N
4	4	1	any	2018-05-11 14:37:54.719042	2018-05-11 14:37:54.719042	\N
5	5	1	any	2018-05-11 14:37:54.719043	2018-05-11 14:37:54.719043	\N
\.


--
-- Data for Name: teams; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.teams (id, name, description, max_members, created_at, updated_at, deleted_at) FROM stdin;
2	13T1	Doi Lop 13T1	0	2018-05-11 14:38:21.360363	2018-05-11 14:38:21.360363	\N
3	13T3	Doi Lop 13T3	0	2018-05-11 14:38:43.065243	2018-05-11 14:53:20.858507	\N
4	13T4	Doi Lop 13T4	0	2018-05-11 14:39:08.076243	2018-05-11 14:39:08.076243	\N
1	13T2	Doi Lop 13T2	0	2018-05-11 14:37:54.713321	2018-05-11 14:37:54.713321	\N
\.


--
-- Data for Name: tournament_teams; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.tournament_teams (id, tournament_id, team_id, score, "group", created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: tournaments; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.tournaments (id, name, description, start_date, end_date, created_at, updated_at, deleted_at) FROM stdin;
1	BKDN_IT 2018	Giai Khoa IT	2018-05-01 15:04:05	2018-06-01 15:04:05	2018-05-11 15:05:48.271787	2018-05-11 15:05:48.271787	\N
\.


--
-- Data for Name: user_follows; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_follows (id, user_id, follow_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, user_name, email, password, first_name, last_name, city, country, about, quote, birthday, role, score, created_at, updated_at, deleted_at) FROM stdin;
3	luongvien	luongvien@gmail.com	$2a$10$zROMjt8DHz916XJ8vgXaH.7w1sY0Lcl7vnzTYUdU.4CUPR5KucivC	Viễn	Lương					\N	user	0	2018-05-11 14:18:04.3854	2018-05-11 14:18:04.3854	\N
5	khacan	khacan@gmail.com	$2a$10$CC2mwlkDQV/xPTWqPZQ3veuzlIIIixYY1u7rA0iyZOxSu4Q7CgZby	Khắc	Ẩn					\N	user	0	2018-05-11 14:19:09.021253	2018-05-11 14:19:09.021253	\N
7	minhhuy	mihhuy@gmail.com	$2a$10$Xdhilu16xx/.lD0EzErKDe9ILn0gjsSPbNc7DMAFaE6pqZUblxbg6	Minh	Huy					\N	user	0	2018-05-11 14:19:56.848607	2018-05-11 14:19:56.848607	\N
8	vanvu	vanvu@gmail.com	$2a$10$bJmfUjOitc9xyD6WRSmaveagYpgA/2B5wVZnnqOonewztDQVzuULG	Văn	Vũ					\N	user	0	2018-05-11 14:20:43.712643	2018-05-11 14:20:43.712643	\N
10	phucminh	phucminh@gmail.com	$2a$10$A9oGUuzEoNigIURgCjIHkO0Il44Q4CVt7.u3iMa5p614mnwIH/ec6	Phúc	Minh					\N	user	0	2018-05-11 14:22:35.900456	2018-05-11 14:22:35.900456	\N
11	ductanh	ductanh@gmail.com	$2a$10$f6GgPl8Gnn/fbsSViuTrd.kV6KrlEmybTcLljaTkUK2xaZACnJQNK	Đức	Tánh					\N	user	0	2018-05-11 14:23:03.662766	2018-05-11 14:23:03.662766	\N
12	vankhanh	vankhanh@gmail.com	$2a$10$YRNI.8fAVSua4x0xcdUQXOFRmW/lq37ALEtZj82LlP0Gy9SVSdRRK	Văn	Khánh					\N	user	0	2018-05-11 14:23:32.65618	2018-05-11 14:23:32.65618	\N
17	vanderossi	vanderossi@gmail.com	$2a$10$uOz3dcnyZ9qdMuz8im6qSeZf/UEi8GAr/SPv77P5BvHPMTtmxSZMq	Van 	Derossi					\N	user	0	2018-05-11 14:25:56.676978	2018-05-11 14:25:56.676978	\N
18	laoan	laoan@gmail.com	$2a$10$MkCLEHFNwomhBO7tNaLO7.HXFjoqgsCVL4tuboD9oDqG9ErpLwfjq	Nguyễn	Ẩn					\N	user	0	2018-05-11 14:26:24.748729	2018-05-11 14:26:24.748729	\N
19	lamia	lamia@gmail.com	$2a$10$QIJoEFyqnvoY7iHO4YdHbeJ3eyCzFNfR793TyGYHsV./Q84gEna5O	Nguyễn	Lâm					\N	user	0	2018-05-11 14:26:51.347388	2018-05-11 14:26:51.347388	\N
4	ducdung	ducdung@gmail.com	$2a$10$SGYG7qDLp7IkBUKlF51RJeZsZE4wlr9GTJWxnkU2FLpkrBPenND5m	Dũng	Văn					\N	user	0	2018-05-11 14:18:33.546974	2018-05-11 14:18:33.546974	\N
9	dinhvo	dinhvo@gmail.com	$2a$10$0ZU3iEtQf6oD7JITNYuH5O0L.zq61GEwIPbxwkCCHypXnWveukyLK	Định	Võ					\N	user	0	2018-05-11 14:21:35.757804	2018-05-11 14:21:35.757804	\N
13	vannghia	vannghia@gmail.com	$2a$10$bgYsjYWXB0Th.9Kx7zLqyu9eMwwAw5X0efdsE/Ybagx4t6/b4wc6G	Nghĩa	Văn					\N	user	0	2018-05-11 14:24:31.097495	2018-05-11 14:24:31.097495	\N
14	vananh	vananh@gmail.com	$2a$10$yXlMwMbCA.WaMdYp4FN/zuQM78Gx2MwzDGHFqGH8vIp7SPxQD79Pq	Anh	Nguyễn					\N	user	0	2018-05-11 14:24:47.300927	2018-05-11 14:24:47.300927	\N
15	vanduc	vanduc@gmail.com	$2a$10$Ix71kFgSLjFj.DovhlfCNOIUPEt4PeHfQgPro2NYuAUaWLzzXHuA.	Đức	Phùng					\N	user	0	2018-05-11 14:25:01.668696	2018-05-11 14:25:01.668696	\N
16	vanhoang	vanhoang@gmail.com	$2a$10$vqkIqgRH.CQ6E7434Rt6Bu01AekcZ5n0FYApmNZ7xi97KQn9vNUE2	Hoàng	Trương					\N	user	0	2018-05-11 14:25:14.737086	2018-05-11 14:25:14.737086	\N
20	nguyenhau	hau13t2@gmail.com	$2a$10$GNLs5tWChV0tvj0yDssXm.G/Ax8HBmGI0vrCPGJeKv5neB0Fix0FO	Hậu	Nguyễn					\N	user	0	2018-05-11 14:27:41.565732	2018-05-11 14:27:41.565732	\N
21	toanngo	ngoductoan12c1a@gmail.com	$2a$10$XjyPm2KXW0d2Q/CjYwUQee2Fm3Rz/F5wXZ2nOeaYm5Y0Vgdb8omu6	Toàn	Ngô					\N	user	0	2018-05-11 14:30:04.832053	2018-05-11 14:30:04.832053	\N
1	congthinh	congthinh@gmail.com	$2a$10$qJUAj9MM4VUDEsZBrmSPt.GESoUsw8r8TZ7tPfmRuBCmXoRL10ASq	Thịnh	Hoàng	Da nang	vietnam		vỏ quýt dày có móng tay nhọn, núi cao còn có núi cao hơn. Nên làm người phải biết vị trí mình đang đứng, tham vọng chỉ vừa đủ và không nên quá ảo tưởng về bản thân.	1995-12-21	s_admin	0	2018-05-11 12:46:59.866569	2018-05-19 06:38:14.385099	\N
22	my	my@gmail.com	$2a$10$7/GOb2e5x7SKBKZ7JUR/cOguN.4P63eGcjAWwEmOcHtDJD3Zqkex.	My	Thị			1111111		\N	user	0	2018-05-11 15:21:39.34827	2018-05-12 02:28:44.998424	\N
6	duchuy	duchuy@gmail.com	$2a$10$dN//OulwIeIfFOQw2zb1G.APo7L9cLEWYxCl9t4E8a537Y4s1TtP.	Đức	Huy	da nang				\N	user	0	2018-05-11 14:19:36.063129	2018-05-13 04:23:15.998759	\N
2	ducan	ducan@gmail.com	$2a$10$MiN7MqXlnSe736cOqUh3XutVaEU0UVaPFHqi61L02FfERnwtXFAp.	An	Đức					\N	s_admin	0	2018-05-11 14:17:39.483253	2018-05-11 14:17:39.483253	\N
23	dungvan123	dungvan1@email.com	$2a$10$Bf.8F4Oh.34S9bc6wRR8MeJqOMQUyapm6KMUCs/HbgHsqwXI8LgKG	dung	van					\N	user	0	2018-05-12 02:22:57.462961	2018-05-12 02:22:57.462961	\N
25	my2	my2@gmail.com	$2a$10$5MT7AgzyMFigsitUxvdXXO1IIn4uS.vU.gb8XrtunAILsfT33U6LC	My	Thị					\N	user	0	2018-05-12 15:29:18.379019	2018-05-12 15:29:18.379019	\N
\.


--
-- Data for Name: videos; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.videos (id, post_id, name, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Name: comment_stars_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.comment_stars_id_seq', 13, true);


--
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.comments_id_seq', 92, true);


--
-- Name: hashtags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.hashtags_id_seq', 5, true);


--
-- Name: images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.images_id_seq', 18, true);


--
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.locations_id_seq', 1, false);


--
-- Name: masters_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.masters_id_seq', 7, true);


--
-- Name: matches_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.matches_id_seq', 2, true);


--
-- Name: post_hashtags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.post_hashtags_id_seq', 5, true);


--
-- Name: post_stars_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.post_stars_id_seq', 9, true);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.posts_id_seq', 46, true);


--
-- Name: star_counts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.star_counts_id_seq', 138, true);


--
-- Name: team_players_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.team_players_id_seq', 28, true);


--
-- Name: teams_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.teams_id_seq', 4, true);


--
-- Name: tournament_teams_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.tournament_teams_id_seq', 1, false);


--
-- Name: tournaments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.tournaments_id_seq', 1, true);


--
-- Name: user_follows_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_follows_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_id_seq', 28, true);


--
-- Name: videos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.videos_id_seq', 1, false);


--
-- Name: comment_stars comment_stars_comment_id_user_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comment_stars
    ADD CONSTRAINT comment_stars_comment_id_user_id_unique UNIQUE (comment_id, user_id);


--
-- Name: comment_stars comment_stars_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comment_stars
    ADD CONSTRAINT comment_stars_pkey PRIMARY KEY (id);


--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: hashtags hashtags_key_word_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hashtags
    ADD CONSTRAINT hashtags_key_word_key UNIQUE (key_word);


--
-- Name: hashtags hashtags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hashtags
    ADD CONSTRAINT hashtags_pkey PRIMARY KEY (id);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);


--
-- Name: locations locations_place_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_place_id_key UNIQUE (place_id);


--
-- Name: masters masters_owner_id_owner_type_user_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.masters
    ADD CONSTRAINT masters_owner_id_owner_type_user_id_unique UNIQUE (owner_id, owner_type, user_id);


--
-- Name: masters masters_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.masters
    ADD CONSTRAINT masters_pkey PRIMARY KEY (id);


--
-- Name: matches matches_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_pkey PRIMARY KEY (id);


--
-- Name: post_hashtags post_hashtags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_hashtags
    ADD CONSTRAINT post_hashtags_pkey PRIMARY KEY (id);


--
-- Name: post_hashtags post_hashtags_post_id_hashtag_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_hashtags
    ADD CONSTRAINT post_hashtags_post_id_hashtag_id_unique UNIQUE (post_id, hashtag_id);


--
-- Name: post_stars post_stars_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_stars
    ADD CONSTRAINT post_stars_pkey PRIMARY KEY (id);


--
-- Name: post_stars post_stars_post_id_user_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_stars
    ADD CONSTRAINT post_stars_post_id_user_id_unique UNIQUE (post_id, user_id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: star_counts star_counts_owner_id_owner_type_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.star_counts
    ADD CONSTRAINT star_counts_owner_id_owner_type_unique UNIQUE (owner_id, owner_type);


--
-- Name: star_counts star_counts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.star_counts
    ADD CONSTRAINT star_counts_pkey PRIMARY KEY (id);


--
-- Name: team_players team_players_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.team_players
    ADD CONSTRAINT team_players_pkey PRIMARY KEY (id);


--
-- Name: team_players team_players_user_id_team_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.team_players
    ADD CONSTRAINT team_players_user_id_team_id_unique UNIQUE (user_id, team_id);


--
-- Name: teams teams_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY (id);


--
-- Name: tournament_teams tournament_teams_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tournament_teams
    ADD CONSTRAINT tournament_teams_pkey PRIMARY KEY (id);


--
-- Name: tournament_teams tournament_teams_tournament_id_team_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tournament_teams
    ADD CONSTRAINT tournament_teams_tournament_id_team_id_unique UNIQUE (tournament_id, team_id);


--
-- Name: tournaments tournaments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_pkey PRIMARY KEY (id);


--
-- Name: user_follows user_follows_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_follows
    ADD CONSTRAINT user_follows_pkey PRIMARY KEY (id);


--
-- Name: user_follows user_follows_user_id_follow_id_unique; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_follows
    ADD CONSTRAINT user_follows_user_id_follow_id_unique UNIQUE (user_id, follow_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_user_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_name_key UNIQUE (user_name);


--
-- Name: videos videos_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.videos
    ADD CONSTRAINT videos_pkey PRIMARY KEY (id);


--
-- Name: comment_stars comment_stars_comment_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comment_stars
    ADD CONSTRAINT comment_stars_comment_id_fkey FOREIGN KEY (comment_id) REFERENCES public.comments(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: comment_stars comment_stars_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comment_stars
    ADD CONSTRAINT comment_stars_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: comments comments_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: comments comments_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: masters masters_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.masters
    ADD CONSTRAINT masters_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: matches matches_team1_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_team1_id_fkey FOREIGN KEY (team1_id) REFERENCES public.teams(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: matches matches_team2_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_team2_id_fkey FOREIGN KEY (team2_id) REFERENCES public.teams(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: matches matches_tournament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matches
    ADD CONSTRAINT matches_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: post_hashtags post_hashtags_hashtag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_hashtags
    ADD CONSTRAINT post_hashtags_hashtag_id_fkey FOREIGN KEY (hashtag_id) REFERENCES public.hashtags(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: post_hashtags post_hashtags_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_hashtags
    ADD CONSTRAINT post_hashtags_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: post_stars post_stars_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_stars
    ADD CONSTRAINT post_stars_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: post_stars post_stars_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_stars
    ADD CONSTRAINT post_stars_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: posts posts_location_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_location_id_fkey FOREIGN KEY (location_id) REFERENCES public.locations(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: posts posts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: team_players team_players_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.team_players
    ADD CONSTRAINT team_players_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: team_players team_players_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.team_players
    ADD CONSTRAINT team_players_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: tournament_teams tournament_teams_team_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tournament_teams
    ADD CONSTRAINT tournament_teams_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: tournament_teams tournament_teams_tournament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tournament_teams
    ADD CONSTRAINT tournament_teams_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: user_follows user_follows_follow_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_follows
    ADD CONSTRAINT user_follows_follow_id_fkey FOREIGN KEY (follow_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- Name: user_follows user_follows_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_follows
    ADD CONSTRAINT user_follows_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

