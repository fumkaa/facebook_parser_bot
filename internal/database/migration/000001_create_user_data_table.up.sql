CREATE TABLE IF NOT EXISTS data_filters (
    id INT PRIMARY KEY AUTO_INCREMENT,
    chat_id BIGINT NOT NULL,
    monitoring VARCHAR(500) DEFAULT '' NOT NULL,
    city VARCHAR(300) DEFAULT '' NOT NULL,
    radius VARCHAR(300) DEFAULT '' NOT NULL,
    category VARCHAR(300) DEFAULT '' NOT NULL,
    filter_file VARCHAR(300) DEFAULT '' NOT NULL
)