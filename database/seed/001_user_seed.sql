-- +goose Up
-- +goose StatementBegin
INSERT INTO users (username, email, password, created_at, updated_at) 
VALUES('test', 'test@test.com', '$2a$04$Re8k56ax4pDhpqxUbbOLdO6PtmUBFc6Lmv7/Pld9cJY/VMDIsNPfW', UNIXEPOCH(), UNIXEPOCH());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id = 1;
-- +goose StatementEnd
