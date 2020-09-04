create table prayer
(
	id INTEGER
		constraint prayer_pk
			primary key autoincrement,
	prayer_text TEXT,
	approved INTEGER DEFAULT 0,
	public INTEGER DEFAULT 0,
	created DATE,
	state INTEGER DEFAULT 0,
	original_test TEXT
);
