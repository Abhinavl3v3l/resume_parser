CREATE TABLE "candidates" (
  "id" SERIAL PRIMARY KEY,
  "resume_file" TEXT NOT NULL,
  "experience_level" INTEGER NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "skills" (
  "id" SERIAL PRIMARY KEY,
  "candidate_id" INT NOT NULL,
  "skill_name" TEXT NOT NULL
);

ALTER TABLE "skills" ADD FOREIGN KEY ("candidate_id") REFERENCES "candidates" ("id");