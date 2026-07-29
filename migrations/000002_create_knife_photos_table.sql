CREATE TABLE knife_photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    knife_id UUID NOT NULL REFERENCES knives(id) ON DELETE CASCADE,
    s3_key TEXT NOT NULL,
    filename TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
