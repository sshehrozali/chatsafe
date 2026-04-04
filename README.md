# Chatsafe

macOS CLI: pack Cursor’s `User` data (`workspaceStorage`, `globalStorage`, `History`, settings) into one timestamped `cursor-backup_*.tar.gz` under `backup/` (created on first run).

```bash
go run ./cmd/chatsafe
go test ./...
```

Close Cursor before backing up.
