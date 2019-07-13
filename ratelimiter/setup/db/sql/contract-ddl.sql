CREATE TABLE contract (
    contractid  bigserial PRIMARY KEY,
    clientname        varchar(100),
    clientgroup       varchar(100),
    allowedlimit integer,
    windowinminutes integer
);