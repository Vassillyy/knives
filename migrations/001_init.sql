CREATE TABLE IF NOT EXISTS knives (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    price INTEGER NOT NULL DEFAULT 0,
    material TEXT,
    blade_length NUMERIC,
    handle TEXT,
    brand TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS knife_photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    knife_id UUID NOT NULL REFERENCES knives(id),
    s3_key TEXT NOT NULL,
    filename TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
