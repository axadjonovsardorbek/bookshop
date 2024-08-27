DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'writing_type_enum') THEN
        DROP TYPE writing_type_enum CASCADE;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
        DROP TYPE role_enum CASCADE;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'working_styles_enum') THEN
        DROP TYPE working_styles_enum CASCADE;
    END IF;
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'working_types_enum') THEN
        DROP TYPE working_types_enum CASCADE;
    END IF;
END $$;

DROP TABLE IF EXISTS vacancies CASCADE;
DROP TABLE IF EXISTS books CASCADE;
DROP TABLE IF EXISTS languages CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS translators CASCADE;
DROP TABLE IF EXISTS publishers CASCADE;
DROP TABLE IF EXISTS authors CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS tokens CASCADE;
