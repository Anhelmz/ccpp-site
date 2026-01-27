-- Revert column name changes if needed
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_name = 'videos' AND column_name = 'youtube_url'
    ) THEN
        ALTER TABLE videos RENAME COLUMN youtube_url TO you_tube_url;
    END IF;
    
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_name = 'videos' AND column_name = 'youtube_id'
    ) THEN
        ALTER TABLE videos RENAME COLUMN youtube_id TO you_tube_id;
    END IF;
END $$;
