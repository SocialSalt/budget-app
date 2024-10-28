CREATE TABLE IF NOT EXISTS transactions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date DATETIME,
	company TEXT,
	category TEXT,
	amount INTEGER,
	account_number TEXT,
	institution TEXT,
	full_description TEXT,
	date_added DATETIME
);
CREATE TABLE IF NOT EXISTS budgets (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	category TEXT,
	[group] TEXT,
	year INTEGER,
	month BLOB,
	amount INTEGER
);
CREATE TABLE IF NOT EXISTS balances (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date DATETIME,
	account TEXT,
	account_number TEXT,
	account_id TEXT,
	balance_id TEXT,
	institution TEXT,
	account_type TEXT,
	class TEXT,
	date_added DATETIME
);
