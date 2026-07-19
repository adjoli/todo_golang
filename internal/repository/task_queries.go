package repository

const sqlSelectTasks = `
SELECT
	id,
	title,
	completed,
	created_at
FROM tasks
`

const sqlOrderTasks = `
ORDER BY id
`