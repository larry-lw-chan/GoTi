-- +goose Up
-- +goose StatementBegin
INSERT INTO users (username, email, password, created_at, updated_at) 
VALUES('Testy', 'test1@gmail.com', '123456', UNIXEPOCH(), UNIXEPOCH());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id = 0;
-- +goose StatementEnd
