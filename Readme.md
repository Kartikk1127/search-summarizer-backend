# 🔎 Search Summarizer Backend

This is the backend service for the **Search Summarizer Chrome Extension**.  
It fetches webpage content, summarizes it into **5–6 bullet points** using an LLM, and returns the result to the frontend extension.

---

## 🚀 Features
- REST API endpoint for summarization
- Fetches and cleans webpage HTML
- Summarizes content into concise bullet points
- Handles errors gracefully (`Summary unavailable`)
- Ready for deployment (VM or serverless)

---

## 🛠 Tech Stack
- **Language**: Go
- **Framework**: Standard `net/http` (REST API)
- **Dependencies**:
    - [`goquery`](https://github.com/PuerkitoBio/goquery) or `go-readability` → extract main text from HTML
    - Go HTTP client → fetch webpage content
    - LLM API client (OpenAI/Anthropic via HTTP)
    - *(Optional)* Redis / in-memory caching

---

## 📡 API Design

### Endpoint
`POST /summarize`

### Request
```json
{
  "url": "https://example.com/article"
}
```

### Response(Success)
```json
{
"summary": [
    "Bullet point 1",
    "Bullet point 2",
    "Bullet point 3",
    "Bullet point 4",
    "Bullet point 5"
  ]
}
```

### Response(Failure)
```json
{
  "summary": null,
  "error": "unavailable"
}
```

## 🔄 Processing Pipeline

1. Receive url from frontend.
2. Fetch webpage HTML.
3. Extract main content only (ignore navbars, ads, boilerplate).
4. Clean and truncate text (to fit LLM token limits).
5. Call LLM API with prompt:
   1. Summarize the following webpage into 5–6 concise bullet points:
      [TEXT HERE]
6. Return JSON response with bullet points.

## ⚠️ Error Handling

1. Webpage fetch fails → return "Summary unavailable"
2. Content extraction fails → return "Summary unavailable"
3. LLM API call fails → return "Summary unavailable"

## Optimizations (Future)

1. Caching: Save {url → summary} to avoid repeat summarizations
2. Rate Limiting: Prevent abuse when made public
3. Batch Mode: Summarize multiple URLs in one request

## ☁️ Deployment
### Phase 1 (Personal Use)

1. Deploy Go backend on a cloud VM (DigitalOcean / AWS EC2 / GCP)
2. Enable CORS for extension requests
3. Hardcode your personal API key (safe for private use)

### Phase 2 (Public Release)

1. Move to serverless (AWS Lambda / GCP Cloud Run)
2. Add API key authentication for users
3. Add rate limits + monitoring

## 📌 Summary

1. This backend acts as the brain of the Search Summarizer project:
2. It fetches, extracts, and summarizes webpage content.
3. Designed for scalability (VM → serverless).
4. Built with Go for simplicity and performance.