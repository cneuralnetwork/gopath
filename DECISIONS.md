# Decisions

- The design artifact URL was inaccessible in this environment. The exact color, type, spacing, and layout values supplied in the prompt are the design source of truth.
- Go is absent from PATH and the network cannot resolve go.dev. The application shows installation instructions when Go is missing. Compilation and Go tests must be run once a toolchain is installed.
- Progress uses an atomic JSON file in the configured data directory. This is the requested fallback when SQLite cannot be added or verified; the environment cannot fetch modernc.org/sqlite.
- YAML metadata files use JSON syntax, which is a valid subset of YAML 1.2, to avoid an unverified parser dependency.
- The server serves a static fallback page if the frontend has not been built; `make build` requires the frontend build.
- The current course files are generated as an initial editable draft. Their counts meet the requested structure, but the exercises and projects need topic-specific implementations and review before they can be presented as a complete curriculum.
- The frontend uses a vendored Markdown parser because the offline package cache does not contain react-markdown. The installed local dependencies were available from another workspace, and `web/dist` is committed for a single-binary build once Go is installed.
