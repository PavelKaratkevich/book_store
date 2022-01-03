CREATE TABLE IF NOT EXISTS books_store (
    ID serial PRIMARY key not null,
    Title text not null,
    Authors text[] not null,
    Year char(4) not null
)
