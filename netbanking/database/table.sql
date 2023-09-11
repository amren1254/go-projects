-- CREATE USER netbanking;

SELECT 'CREATE DATABASE netbanking' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'netbanking')\gexec
GRANT ALL PRIVILEGES ON DATABASE netbanking TO netbanking;

CREATE SCHEMA IF NOT EXISTS netbanking
    AUTHORIZATION netbanking;

-- ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA netbanking
-- GRANT ALL ON TABLES TO netbanking;

CREATE TYPE netbanking.user_status AS ENUM
    ('active', 'inactive');

ALTER TYPE netbanking.user_status
    OWNER TO netbanking;

CREATE TYPE netbanking.account_type AS ENUM
    ('savings', 'current');

ALTER TYPE netbanking.account_type
    OWNER TO netbanking;


CREATE TABLE netbanking."user"
(
    id uuid NOT NULL UNIQUE,
    name character varying(50),
    username character varying(50) NOT NULL UNIQUE,
    password character varying(100),
    status netbanking.user_status NOT NULL,
    phone bigint,
    email character varying(100) UNIQUE,
    created_at timestamp with time zone,
    updated_at timestamp with time zone DEFAULT current_timestamp,
    CONSTRAINT pk PRIMARY KEY (id, username, email)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS netbanking."user"
    OWNER to netbanking;


CREATE TABLE netbanking.account
(
    id uuid NOT NULL UNIQUE,
    account_number bigserial NOT NULL UNIQUE,
    account_type netbanking.account_type NOT NULL,
    total_amount bigint,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
    CONSTRAINT account_pk PRIMARY KEY (id, account_number)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS netbanking.account
    OWNER to netbanking;
