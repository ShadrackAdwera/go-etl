CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expires_at" timestamptz NOT NULL
);

CREATE TABLE "files" (
  "id" bigserial PRIMARY KEY,
  "file_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by_id" bigint NOT NULL
);

CREATE TABLE "data" (
  "id" bigserial PRIMARY KEY,
  "home_scored" integer NOT NULL,
  "away_scored" integer NOT NULL,
  "home_team" varchar NOT NULL,
  "away_team" varchar NOT NULL,
  "match_day" timestamptz NOT NULL DEFAULT (now()),
  "referee" varchar NOT NULL,
  "winner" varchar NOT NULL,
  "season" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by_id" bigint NOT NULL,
  "file_id" bigint NOT NULL
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "files" ADD FOREIGN KEY ("created_by_id") REFERENCES "users" ("id");

ALTER TABLE "data" ADD FOREIGN KEY ("created_by_id") REFERENCES "users" ("id");

ALTER TABLE "data" ADD FOREIGN KEY ("file_id") REFERENCES "files" ("id");