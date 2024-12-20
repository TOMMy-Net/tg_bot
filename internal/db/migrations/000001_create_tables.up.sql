CREATE TABLE IF NOT EXISTS application(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS support (
	"id" VARCHAR(255) NOT NULL UNIQUE,
	"user_id" INTEGER NOT NULL,
	"problem" VARCHAR(255) NOT NULL,
	PRIMARY KEY("id")
);
