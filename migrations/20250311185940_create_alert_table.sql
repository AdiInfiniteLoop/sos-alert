-- +goose Up
-- +goose StatementBegin
CREATE TABLE alerts (
    id SERIAL PRIMARY KEY NOT NULL ,
    MESSAGE VARCHAR(255) ,
    AUTHOR VARCHAR(50) NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE alerts
-- +goose StatementEnd
