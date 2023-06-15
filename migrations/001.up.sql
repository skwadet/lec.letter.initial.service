CREATE TABLE "letters"
(
    "id"              uuid PRIMARY KEY NOT NULL,
    "submission_id"   uuid             NOT NULL,
    "created_at"      timestamp DEFAULT CURRENT_TIMESTAMP,
    "additional_info" varchar DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "objectives"
(
    "id"           uuid PRIMARY KEY NOT NULL,
    "order_number" int              NOT NULL,
    "letter_id"    uuid             NOT NULL,
    "title"        varchar             NOT NULL,
    "created_at"   timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "purposes"
(
    "id"           uuid PRIMARY KEY NOT NULL,
    "order_number" int              NOT NULL,
    "letter_id"    uuid             NOT NULL,
    "title"        varchar             NOT NULL,
    "created_at"   timestamp DEFAULT CURRENT_TIMESTAMP
);