

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE SEQUENCE public.account_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customer_id_seq OWNER TO postgres;
ALTER TABLE public.account_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id integer DEFAULT nextval('public.customer_id_seq'::regclass) NOT NULL,
    customer_id character varying(255),
    name character varying(255),
    date_of_birth date,
    city character varying(255),
    zipcode character varying(10),
    status integer DEFAULT 1
);

CREATE TABLE public.accounts (
    id integer DEFAULT nextval('public.account_id_seq'::regclass) NOT NULL,
    account_id character varying(255),
    customer_id character varying(255),
    opening_date timestamp without time zone,
    account_type character varying(50),
    status integer DEFAULT 1,
    amount numeric(10, 2)
);


ALTER TABLE public.customers OWNER TO postgres;
ALTER TABLE public.accounts OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customer_id_seq', 1, true);
SELECT pg_catalog.setval('public.account_id_seq', 1, true);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);



