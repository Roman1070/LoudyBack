CREATE TABLE IF NOT EXISTS artists(
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    bio TEXT,
    cover TEXT,
    likes_count INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS artist_name_idx ON artists(name);

CREATE TABLE IF NOT EXISTS artists_albums(
    id SERIAL PRIMARY KEY,
    artist_id INTEGER NOT NULL,
    album_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS artist_idx ON artists_albums(artist_id);
CREATE INDEX IF NOT EXISTS album_idx ON artists_albums(album_id);

CREATE TABLE IF NOT EXISTS albums(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    tracks_ids INTEGER[] NOT NULL,
    cover TEXT,
    release_date date
);

CREATE INDEX IF NOT EXISTS albums_name_idx ON albums(name);

CREATE TABLE IF NOT EXISTS tracks(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    file TEXT NOT NULL,
    album_id INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS tracks_name_idx ON tracks(name);