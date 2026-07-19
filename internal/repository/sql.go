package repository

const sqlInsertTask = `
INSERT INTO tasks (
	title,
	completed,
	created_at
)
VALUES (?, ?, ?);
`

const sqlFindTaskByID = `
SELECT
	id,
	title,
	completed,
	created_at
FROM tasks
WHERE id = ?;
`

const sqlUpdateTask = `
UPDATE tasks
SET
	title = ?,
	completed = ?
WHERE id = ?;
`

const sqlDeleteTask = `
DELETE FROM tasks
WHERE id = ?;
`
