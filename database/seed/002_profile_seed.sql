-- +goose Up
-- +goose StatementBegin
INSERT INTO profiles (username, name, bio, link, avatar, private, user_id, created_at, updated_at) 
VALUES('testy_one', 'Testy One', 'I am a test user', 'www.test.com', null, 0, 1, UNIXEPOCH(), UNIXEPOCH());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM profiles WHERE id = 1;
-- +goose StatementEnd
