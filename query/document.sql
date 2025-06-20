-- name: InsertDocument :exec
INSERT INTO template_documents (title, filename) VALUES (?, ?);

-- name: GetDocumentByID :one
SELECT * FROM template_documents WHERE id = ?;

-- name: ListDocuments :many
SELECT * FROM template_documents;