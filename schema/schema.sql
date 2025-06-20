--  template documents table: use to save template filename and title.
--  filename not included path.
--  title is pdf title
CREATE TABLE IF NOT EXISTS template_documents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    filename TEXT NOT NULL,
    upload_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- template chunks table: use to save a part of template text and the text offset
-- (use to tell the frontend highlight the chunk if the chunk is dump)
CREATE TABLE IF NOT EXISTS template_chunks (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- template chunks id
    doc_id INTEGER NOT NULL, -- documents id
    offset_start INTEGER NOT NULL, -- template text chunk offset start
    offset_end INTEGER NOT NULL, -- template text chunk offset end
    text TEXT NOT NULL, -- chunk content
    FOREIGN KEY (doc_id) REFERENCES template_documents(id) ON DELETE CASCADE -- foreign key reference to documents
);

-- index of table of template chunks binding doc_id
CREATE INDEX IF NOT EXISTS idx_template_chunks_doc_id ON template_chunks(doc_id);
