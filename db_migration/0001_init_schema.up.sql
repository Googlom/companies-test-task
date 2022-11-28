CREATE TABLE companies
(
    id           uuid DEFAULT uuid_generate_v4(),
    name         VARCHAR(15) UNIQUE NOT NULL,
    description  VARCHAR(3000),
    employees    INT                NOT NULL,
    registered   BOOLEAN            NOT NULL,
    company_type INT                NOT NULL,
    PRIMARY KEY (id)
);
