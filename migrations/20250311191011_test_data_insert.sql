-- +goose Up
-- +goose StatementBegin
INSERT INTO alerts(MESSAGE, AUTHOR) VALUES ('Test From Goose', 'AdixPradhan');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM alerts;
-- +goose StatementEnd
