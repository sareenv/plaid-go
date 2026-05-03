-- +goose UP
CREATE TABLE plaid_items (
                             id BIGSERIAL PRIMARY KEY,
                             item_id TEXT UNIQUE,
                             access_token TEXT,
                             created_at TIMESTAMP,
                             updated_at TIMESTAMP,
                             deleted_at TIMESTAMP,
                             user_id BIGINT,
                             CONSTRAINT fk_user
                                 FOREIGN KEY (user_id)
                                     REFERENCES users(id)
                                     ON DELETE SET NULL
                                     ON UPDATE CASCADE
);
-- +goose DOWN
drop table plaid_items;