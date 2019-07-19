CREATE TABLE contract (
    contractid  bigserial UNIQUE,
    clientname        varchar(100),
    clientgroup       varchar(100),
    allowedlimit integer,
    windowinminutes integer,
    PRIMARY KEY(clientname, clientgroup)
);

CREATE TABLE api (
    id bigserial UNIQUE,
    api varchar(2400),
    contractid bigserial references contract(contractid),
    clientgroup varchar(100),
    PRIMARY KEY(contractid, api)
);