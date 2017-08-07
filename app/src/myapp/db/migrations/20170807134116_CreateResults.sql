
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE results (
    id int NOT NULL AUTO_INCREMENT,
    user_id int NOT NULL,
    status bool DEFAULT false,
    created_at datetime,
    updated_at datetime,
    PRIMARY KEY(id)
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE results;