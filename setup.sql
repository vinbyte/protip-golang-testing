SET timezone = 'Asia/Jakarta';

-- setup users dan data
CREATE TABLE public.users (
	id serial NOT NULL,
	"name" varchar(255) NOT NULL,
	email varchar(255) not null,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);
INSERT INTO users (id, "name", email) VALUES(1, 'Spongebob', 'spongebob@bikinibottom.com');
INSERT INTO users (id, "name", email) VALUES(2, 'Patrick', 'patrick@bikinibottom.com');
