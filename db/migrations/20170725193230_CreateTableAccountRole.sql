
-- +goose Up
CREATE TABLE account_role
(
  user_id INTEGER NOT NULL,
  role_id INTEGER NOT NULL,
  grant_date TIMESTAMP WITHOUT TIME ZONE,
  PRIMARY KEY (user_id, role_id),
  CONSTRAINT account_role_role_id_fkey FOREIGN KEY (role_id)
    REFERENCES role (role_id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT account_role_user_id_fkey FOREIGN KEY (user_id)
    REFERENCES account (user_id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
);


-- +goose Down
DROP TABLE IF EXISTS account_role;
