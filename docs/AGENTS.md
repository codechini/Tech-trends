# AGENTS.md - Developer & AI Guidance for Tech Velocity Engine

This document defines architecture boundaries, conventions, and constraints for human developers and AI assistants working on this repository.

---

## 1. Project Mission & Constraints
`Tech Velocity Engine` is a high-performance, single-binary Go monolith designed to run on free-tier cloud infrastructure.
- **Primary Goal:** Scrape, tokenize, score, and visualize tech topic velocity.
- **Memory Footprint Target:** < 50MB RAM at runtime.
- **Budget Constraint:** $0 hosting cost limit (Fly.io/Render + Supabase Free Tier).

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

