CREATE TABLE contract (
    contractid  bigserial PRIMARY KEY,
    clientname        varchar(100),
    clientgroup       varchar(100),
    allowedlimit integer,
    windowinminutes integer
);

CREATE TABLE api (
    id bigserial PRIMARY KEY,
    api varchar(2400),
    contractid bigserial references contract(contractid),
    clientgroup varchar(100)
);