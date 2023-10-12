ALTER TABLE "candidates"
ADD COLUMN "resume_file" TEXT NOT NULL;

ALTER TABLE "candidates"
DROP COLUMN "email";
