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

## 📅 1 週開發計畫（Turnipin Plagiarism Checker）

### Day 1：專案初始化
- [x] 建立 GitHub 倉庫與 README 說明
- [x] 建立 frontend / backend 基本目錄結構
- [x] 設定 sqlc.yaml，產生 query code
- [x] 建立 schema.sql 與 query.sql
- [ ] 初始化 config 模組（使用 Viper）
- [ ] 撰寫 repository 層與 model interface

---

### Day 2：文字處理模組
- [ ] 撰寫 PDF / Word 文字提取（extractor）
- [ ] 撰寫文本清洗與 Token 化（preprocessor）
- [ ] 撰寫 N-Gram 擷取與 Shingle 工具
- [ ] 撰寫 Levenshtein 相似度計算模組
- [ ] 撰寫對應單元測試

---

### Day 3：資料上傳與儲存
- [ ] 上傳 API（/api/reference/upload、/api/check/upload）
- [ ] 提取文字 → 擷取 Shingle → 儲存 DB
- [ ] 儲存原始檔案（PDF/Word）至 `files/` 資料夾
- [ ] 回傳 JSON 資訊含 ID、狀態、檔名等

---

### Day 4：比對與報告產生
- [ ] 比對主流程服務邏輯
- [ ] Jaccard + Levenshtein 混合比對邏輯
- [ ] 報告格式整理並提供 `/api/check/report/{id}`
- [ ] 檢查與參考資料的段落精比對與高亮資訊

---

### Day 5：登入系統與權限保護
- [ ] 單一使用者登入功能（JWT 驗證）
- [ ] 前端登入頁與 localStorage 儲存 Token
- [ ] Middleware 驗證 Token 是否存在

---

### Day 6：前端 UI 開發（Svelte + Tailwind）
- [ ] 首頁：參考資料列表、上傳區塊、查看原始檔案
- [ ] 查重頁：上傳檔案區、提交按鈕、進度條
- [ ] 報告頁：比對結果展示、高亮段落顯示
- [ ] 加入介面提示與 Loading 狀態

---

### Day 7：整合測試與打包
- [ ] 撰寫 integration 測試流程
- [ ] 處理錯誤回傳與 API 整合問題
- [ ] 完整打包：Go build + 前端 Vite build
- [ ] 準備最終 USB friendly 部署版本（含 SQLite + 檔案資料夾）

---

## 📄 License

本專案採用 [MIT License](./LICENSE) 授權。

© 2025 林均 保留所有權利。
