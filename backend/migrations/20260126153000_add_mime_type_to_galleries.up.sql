BEGIN;

-- Add mime_type column as nullable first
ALTER TABLE galleries
    ADD COLUMN IF NOT EXISTS mime_type TEXT;

-- Backfill mime_type based on filename extension for existing records
UPDATE galleries
SET mime_type = CASE
    WHEN LOWER(filename) LIKE '%.jpg' OR LOWER(filename) LIKE '%.jpeg' THEN 'image/jpeg'
    WHEN LOWER(filename) LIKE '%.png' THEN 'image/png'
    WHEN LOWER(filename) LIKE '%.gif' THEN 'image/gif'
    WHEN LOWER(filename) LIKE '%.webp' THEN 'image/webp'
    ELSE 'image/jpeg' -- Default fallback
END
WHERE mime_type IS NULL;

-- Also check path field for data URIs
UPDATE galleries
SET mime_type = CASE
    WHEN path LIKE 'data:image/jpeg%' THEN 'image/jpeg'
    WHEN path LIKE 'data:image/png%' THEN 'image/png'
    WHEN path LIKE 'data:image/gif%' THEN 'image/gif'
    WHEN path LIKE 'data:image/webp%' THEN 'image/webp'
    ELSE mime_type
END
WHERE mime_type IS NULL OR mime_type = 'image/jpeg';

-- Set default for any remaining NULL values
UPDATE galleries
SET mime_type = 'image/jpeg'
WHERE mime_type IS NULL;

-- Now make it NOT NULL
ALTER TABLE galleries
    ALTER COLUMN mime_type SET NOT NULL;

COMMIT;
