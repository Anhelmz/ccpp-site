BEGIN;

-- Add richer scheduling fields for events
ALTER TABLE events
    ADD COLUMN IF NOT EXISTS start_time         TIMESTAMPTZ,
    ADD COLUMN IF NOT EXISTS end_time           TIMESTAMPTZ,
    ADD COLUMN IF NOT EXISTS recurrence         TEXT DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS recurrence_until   TIMESTAMPTZ,
    ADD COLUMN IF NOT EXISTS category           TEXT,
    ADD COLUMN IF NOT EXISTS color              TEXT,
    ADD COLUMN IF NOT EXISTS timezone           TEXT DEFAULT 'UTC';

-- Backfill start/end times from legacy date column if present
UPDATE events
SET start_time = date
WHERE start_time IS NULL AND date IS NOT NULL;

UPDATE events
SET end_time = start_time + INTERVAL '1 hour'
WHERE end_time IS NULL AND start_time IS NOT NULL;

-- Ensure timezone has a value
UPDATE events
SET timezone = 'UTC'
WHERE timezone IS NULL OR timezone = '';

-- Enforce non-nullable start/end times now that they are backfilled
ALTER TABLE events
    ALTER COLUMN start_time SET NOT NULL,
    ALTER COLUMN end_time   SET NOT NULL;

-- Guard against inverted time ranges
ALTER TABLE events
    ADD CONSTRAINT events_end_after_start CHECK (end_time >= start_time);

-- Helpful indexes for queries
CREATE INDEX IF NOT EXISTS idx_events_start_time ON events (start_time);
CREATE INDEX IF NOT EXISTS idx_events_category ON events (category);

COMMIT;




