-- Fix column names if they were created with wrong naming
-- Check if you_tube_url exists and rename it to youtube_url
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_name = 'videos' AND column_name = 'you_tube_url'
    ) THEN
        ALTER TABLE videos RENAME COLUMN you_tube_url TO youtube_url;
    END IF;
    
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_name = 'videos' AND column_name = 'you_tube_id'
    ) THEN
        ALTER TABLE videos RENAME COLUMN you_tube_id TO youtube_id;
    END IF;
END $$;
