create table prayer
(
	id INTEGER
		constraint prayer_pk
			primary key autoincrement,
	prayer_text TEXT,
	approved INTEGER DEFAULT 0,
	public INTEGER DEFAULT 0,
	created DATE DEFAULT CURRENT_DATE,
	state INTEGER DEFAULT 0,
	original_test TEXT
);

create table authorization (
    auth_key TEXT primary key,
    type INTEGER DEFAULT 0
);