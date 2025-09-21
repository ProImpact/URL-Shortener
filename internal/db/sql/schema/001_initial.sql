-- +goose Up
-- +goose StatementBegin

CREATE TABLE "url" (
    "id" TEXT NOT NULL,
    "original" TEXT NOT NULL UNIQUE,
    "shortened" TEXT NOT NULL UNIQUE,
    "clicks" INT NOT NULL DEFAULT 0,
    "created" TIMESTAMP NOT NULL,
    PRIMARY KEY("id")
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE "url";
-- +goose StatementEnd
