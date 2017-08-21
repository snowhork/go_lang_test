
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE stages (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    user_id int NOT NULL,
    created_at datetime,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE stages;
