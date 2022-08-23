--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-1.pgdg110+1)
-- Dumped by pg_dump version 14.3

-- Started on 2022-08-21 18:31:15

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 16385)
-- Name: breed; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.breed (
    id integer DEFAULT 1 NOT NULL,
    breed_name character varying(200) NOT NULL,
    category_id integer NOT NULL
);


ALTER TABLE public.breed OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16390)
-- Name: category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category (
    id integer DEFAULT 1 NOT NULL,
    category_name character varying(200) NOT NULL
);


ALTER TABLE public.category OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16395)
-- Name: location; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.location (
    id integer DEFAULT 1 NOT NULL,
    location_name character varying(200) NOT NULL
);


ALTER TABLE public.location OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16400)
-- Name: pet; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pet (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    age integer NOT NULL,
    image_url character varying(500) NOT NULL,
    description character varying(500),
    breed_id integer NOT NULL,
    category_id integer NOT NULL,
    location_id integer NOT NULL
);


ALTER TABLE public.pet OWNER TO postgres;

--
-- TOC entry 3182 (class 2606 OID 16389)
-- Name: breed breed_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.breed
    ADD CONSTRAINT breed_pkey PRIMARY KEY (id);


--
-- TOC entry 3184 (class 2606 OID 16394)
-- Name: category category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- TOC entry 3186 (class 2606 OID 16399)
-- Name: location location_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.location
    ADD CONSTRAINT location_pkey PRIMARY KEY (id);


-- Completed on 2022-08-21 18:31:15

--
-- PostgreSQL database dump complete
--

