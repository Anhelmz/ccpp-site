BEGIN;

-- Remove mime_type column
ALTER TABLE galleries
    DROP COLUMN IF EXISTS mime_type;

COMMIT;
