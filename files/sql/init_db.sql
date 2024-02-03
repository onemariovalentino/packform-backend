CREATE TYPE product_kind AS ENUM ('Corrugated Box','Hand Sanitizer');

/* -- tbl_companies -- */
CREATE TABLE tbl_companies (
    company_id bigint NOT NULL,
    company_name character varying(100) NOT NULL
);
CREATE SEQUENCE tbl_companies_company_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE tbl_companies_company_id_seq OWNED BY tbl_companies.company_id;
ALTER TABLE ONLY tbl_companies ALTER COLUMN company_id SET DEFAULT nextval('tbl_companies_company_id_seq'::regclass);
ALTER TABLE ONLY tbl_companies
    ADD CONSTRAINT tbl_companies_pkey PRIMARY KEY (company_id);

/* -- tbl_customers -- */
CREATE TABLE tbl_customers (
    user_id character varying(30) NOT NULL,
    login character varying(30) NOT NULL,
    password character varying(100) NOT NULL,
    name character varying(60) NOT NULL,
    company_id bigint NOT NULL,
    credit_cards character varying(60) NOT NULL
);
ALTER TABLE ONLY tbl_customers
    ADD CONSTRAINT tbl_customers_pkey PRIMARY KEY (user_id);
ALTER TABLE ONLY tbl_customers
    ADD CONSTRAINT fk_tbl_companies_customer FOREIGN KEY (company_id) REFERENCES tbl_companies(company_id);

/* -- tbl_orders -- */
CREATE TABLE tbl_orders (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    order_name character varying(100) NOT NULL,
    customer_id character varying(30) NOT NULL
);
CREATE SEQUENCE tbl_orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE tbl_orders_id_seq OWNED BY tbl_orders.id;
ALTER TABLE ONLY tbl_orders ALTER COLUMN id SET DEFAULT nextval('tbl_orders_id_seq'::regclass);
ALTER TABLE ONLY tbl_orders
    ADD CONSTRAINT tbl_orders_pkey PRIMARY KEY (id);
ALTER TABLE ONLY tbl_orders
    ADD CONSTRAINT fk_tbl_customers_order FOREIGN KEY (customer_id) REFERENCES tbl_customers(user_id);

/* -- tbl_order_items --*/
CREATE TABLE tbl_order_items (
    id bigint NOT NULL,
    order_id bigint,
    price_per_unit numeric(12,4),
    quantity bigint,
    product product_kind
);
CREATE SEQUENCE tbl_order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE tbl_order_items_id_seq OWNED BY tbl_order_items.id;
ALTER TABLE ONLY tbl_order_items ALTER COLUMN id SET DEFAULT nextval('tbl_order_items_id_seq'::regclass);
ALTER TABLE ONLY tbl_order_items
    ADD CONSTRAINT tbl_order_items_pkey PRIMARY KEY (id);
ALTER TABLE ONLY tbl_order_items
    ADD CONSTRAINT fk_tbl_orders_order_item FOREIGN KEY (order_id) REFERENCES tbl_orders(id);

/* -- tbl_order_item_deliveries --*/
CREATE TABLE tbl_order_item_deliveries (
    id bigint NOT NULL,
    order_item_id bigint,
    delivered_quantity bigint
);
CREATE SEQUENCE tbl_order_item_deliveries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE tbl_order_item_deliveries_id_seq OWNED BY tbl_order_item_deliveries.id;
ALTER TABLE ONLY tbl_order_item_deliveries ALTER COLUMN id SET DEFAULT nextval('tbl_order_item_deliveries_id_seq'::regclass);
ALTER TABLE ONLY tbl_order_item_deliveries
    ADD CONSTRAINT tbl_order_item_deliveries_pkey PRIMARY KEY (id);
ALTER TABLE ONLY tbl_order_item_deliveries
    ADD CONSTRAINT fk_tbl_order_items_order_item_delivery FOREIGN KEY (order_item_id) REFERENCES tbl_order_items(id);


