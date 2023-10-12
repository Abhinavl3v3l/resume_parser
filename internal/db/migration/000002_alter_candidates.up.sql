ALTER TABLE "candidates"
DROP COLUMN "resume_file";

ALTER TABLE "candidates"
ADD COLUMN "email" TEXT NOT NULL UNIQUE;

