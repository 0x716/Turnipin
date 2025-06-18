# 🧠 Turnipin

> A fully local, offline-capable plagiarism checker for academic documents.  
> Built with Go + Gin + SqlC + SQLite on the backend, and Svelte + Tailwind on the frontend.

---

## 💡 Features

- Upload reference documents (PDF / Word)
- Upload documents for similarity checking
- Uses N-gram + Jaccard + Levenshtein distance for similarity scoring
- Store raw files and extracted text locally
- Built-in secure login system (single-user only)
- Realtime progress indicator for checks
- Plagiarism result report with highlighted segments

---

## 🛠 Tech Stack

- Backend: **Go + Gin + SqlC + SQLite**
- Frontend: **Svelte + Tailwind CSS v3**
- File Extraction: PDF/Word parsing
- Similarity: N-gram → Jaccard Filter → Levenshtein Precision Match

---

## 📂 Folder Structure

```
backend/
├── api/ # Gin API handlers
├── service/ # Business logic
├── repository/ # SQLc + repository pattern
├── db/ # schema.sql, query.sql, sqlc.yaml
├── util/ # Extractors, tokenizers, matchers
├── files/ # Uploaded original/reference docs
frontend/
├── src/ # Svelte components, routes
└── public/ # Static assets
```

---

## 📅 1-Week Development Plan (Turnipin Plagiarism Checker)

### Day 1: Project Initialization
- [x] Create GitHub repository and README description
- [x] Setup basic frontend/backend directory structure
- [x] Configure `sqlc.yaml` and generate query code
- [x] Create `schema.sql` and `query.sql`
- [ ] Initialize config module (using Viper)
- [ ] Implement repository layer and model interfaces

---

### Day 2: Text Processing Module
- [ ] Develop PDF / Word text extraction (extractor)
- [ ] Implement text cleaning and tokenization (preprocessor)
- [ ] Create N-Gram extraction and shingle utilities
- [ ] Write Levenshtein similarity calculation module
- [ ] Write corresponding unit tests

---

### Day 3: File Upload and Storage
- [ ] Implement upload APIs (`/api/reference/upload`, `/api/check/upload`)
- [ ] Extract text → generate shingles → store in DB
- [ ] Save original files (PDF/Word) to `files/` directory
- [ ] Return JSON info including ID, status, filename, etc.

---

### Day 4: Comparison and Report Generation
- [ ] Implement main comparison service logic
- [ ] Develop combined Jaccard + Levenshtein matching logic
- [ ] Format reports and provide `/api/check/report/{id}`
- [ ] Highlight precise matching segments with reference texts

---

### Day 5: Login System and Access Control
- [ ] Single-user login functionality (JWT authentication)
- [ ] Frontend login page and localStorage token storage
- [ ] Middleware to verify token presence on requests

---

### Day 6: Frontend UI Development (Svelte + Tailwind)
- [ ] Homepage: list reference documents, upload sections, file preview
- [ ] Plagiarism check page: upload area, submit button, progress bar
- [ ] Report page: display comparison results with highlighted segments
- [ ] Add UI feedback and loading states

---

### Day 7: Integration Testing and Packaging
- [ ] Write integration test cases covering the entire flow
- [ ] Handle error responses and API integration issues
- [ ] Complete build: Go backend + frontend Vite build
- [ ] Prepare final USB-friendly deployment package (including SQLite + file folders)

---

## 📄 License

This project is licensed under the [MIT License](./LICENSE).

© 2025 Jun Lin. All rights reserved.
