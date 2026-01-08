BEGIN;

-- Remove added indexes
DROP INDEX IF EXISTS idx_events_start_time;
DROP INDEX IF EXISTS idx_events_category;

-- Drop constraints and columns introduced in the upgrade
ALTER TABLE events
    DROP CONSTRAINT IF EXISTS events_end_after_start,
    DROP COLUMN IF EXISTS start_time,
    DROP COLUMN IF EXISTS end_time,
    DROP COLUMN IF EXISTS recurrence,
    DROP COLUMN IF EXISTS recurrence_until,
    DROP COLUMN IF EXISTS category,
    DROP COLUMN IF EXISTS color,
    DROP COLUMN IF EXISTS timezone;

COMMIT;



