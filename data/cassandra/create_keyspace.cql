CREATE KEYSPACE drinklogs_service WITH REPLICATION = {
    'class': 'SimpleStrategy',
    'replication_factor': 1
};