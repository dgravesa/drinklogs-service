CREATE TABLE drinklogs (
    uid bigint,
    time timeuuid,
    amount double,
    PRIMARY KEY (uid, time)
);