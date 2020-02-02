USE drinklogs_service;

CREATE TABLE drinklogs (
    uid varchar,
    time timeuuid,
    amount double,
    PRIMARY KEY (uid, time)
);