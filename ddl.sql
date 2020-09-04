create table prayer
(
	id INTEGER
		constraint prayer_pk
			primary key autoincrement,
	prayer_text TEXT,
	approved int,
	public int,
	created DATE,
	state int,
	original_test int
);
