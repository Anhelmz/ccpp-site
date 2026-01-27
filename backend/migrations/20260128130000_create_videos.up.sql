CREATE TABLE IF NOT EXISTS videos (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    youtube_id VARCHAR(11) NOT NULL,
    youtube_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_videos_youtube_id ON videos (youtube_id);
CREATE INDEX IF NOT EXISTS idx_videos_deleted_at ON videos (deleted_at);
