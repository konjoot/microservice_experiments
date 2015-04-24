puts Config.database_url
Sequel.connect(Config.database_url, max_connections: Config.db_pool)
