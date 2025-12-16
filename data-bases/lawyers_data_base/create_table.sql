CREATE TABLE lawyers (
	id SERIAL PRIMARY KEY,
	name TEXT,
	city TEXT,
	wins INT,
	loses INT
);
INSERT INTO lawyers (name, city, wins, loses) VALUES ('Ronaldo', 'Moscow', 24, 9), ('Chicago', 'Santiago', 814, 98), ('Renat', 'Omsk', 228, 52), ('Misha', 'Omsk', 0, 100), ('Leha', 'Madrid', 412, 312)