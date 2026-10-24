# Decisions

- The design artifact URL was inaccessible in this environment. The exact color, type, spacing, and layout values supplied in the prompt are the design source of truth.
- Go was initially absent, so the app includes a clear installation page. Go 1.27.3 was then installed with Homebrew and the build, app tests, content verification, and local API smoke checks passed.
- Progress uses an atomic JSON file in the configured data directory. This is the requested fallback to SQLite; it keeps the app dependency free on the Go side.
- YAML metadata files use JSON syntax, which is a valid subset of YAML 1.2, to avoid an unverified parser dependency.
- The server serves a static fallback page if the frontend has not been built; `make build` requires the frontend build.
- The current course files are generated as an initial editable draft. Their counts meet the requested structure, but the exercises and projects need topic-specific implementations and review before they can be presented as a complete curriculum.
- The frontend uses a vendored Markdown parser because the offline package cache does not contain react-markdown. The installed local dependencies were available from another workspace, and `web/dist` is committed for a single-binary build.
