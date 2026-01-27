CREATE TABLE IF NOT EXISTS events_calendar (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    details TEXT,
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    location TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_events_calendar_start_date ON events_calendar (start_date);
CREATE INDEX IF NOT EXISTS idx_events_calendar_end_date ON events_calendar (end_date);
CREATE INDEX IF NOT EXISTS idx_events_calendar_deleted_at ON events_calendar (deleted_at);
