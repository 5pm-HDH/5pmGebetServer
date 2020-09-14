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

INSERT INTO authorization(auth_key, type) VALUES ('test_master_key', 0);
INSERT INTO authorization(auth_key, type) VALUES ('test_pray_key', 1);
INSERT INTO authorization(auth_key, type) VALUES ('test_view_key', 2);

