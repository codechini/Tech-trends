# AGENTS.md - Developer & AI Guidance for Tech Velocity Engine

This document defines architecture boundaries, conventions, and constraints for human developers and AI assistants working on this repository.

---

## 1. Project Mission & Constraints
`Tech Velocity Engine` is a high-performance, single-binary Go monolith designed to run on free-tier cloud infrastructure.
- **Primary Goal:** Scrape, tokenize, score, and visualize tech topic velocity.
- **Memory Footprint Target:** < 50MB RAM at runtime.
- **Budget Constraint:** $0 hosting cost limit (Fly.io/Render + Supabase Free Tier, OCI(Oracle Cloud Infrastructure)).

---

## 2. Directory Structure & Boundary Rules

```text
tech-velocity-monolith/
├── cmd/app/               # Application entrypoint & routing setup ONLY
├── internal/
│   ├── config/            # Env loading & validation
│   ├── database/          # Postgres pool setup & sqlc generated queries
│   ├── ingester/          # Scraping logic, N-gram extraction, tokenizers
│   ├── metrics/           # Velocity algorithms & Prometheus definitions
│   ├── testutils/         # Mocks, fixtures, and TDD helper functions
│   └── web/               # Handlers, templ templates, and static assets
├── migrations/            # SQL migrations for Supabase
├── AGENTS.md              # Repository rules for AI coders
└── Makefile               # Task runner (test, build, lint)

---

```
## 3. How the Spike Detection Logic Works

1. **Unigrams & Bigrams:** Pure keyword matching often misses contextual topics. The `ExtractKeywords` function tracks both single words (`cat`) and word pairs (`cat food`, `quantum computing`).
2. **Deduplication per Article:** Keywords are incremented once per article so a single post repeating the word "cat" 50 times doesn't trigger a false positive.
3. **Z-Score / Relative Spike Metric:**

$$\text{Expected Recent Count} = \left(\frac{\text{Historical Baseline Count}}{38 \text{ days}}\right) \times 7 \text{ days}$$

If the keyword count in the last 7 days is higher than $\text{Expected Recent Count} \times \text{Multiplier}$ (and meets your minimum threshold of 3–5 articles), it is flagged as an emerging trend.
