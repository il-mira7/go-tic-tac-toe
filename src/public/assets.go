package public

import "embed"

// index.html встраивается в бинарник: файл не нужен рядом с процессом (важно для docker).
//
//go:embed index.html
var Root embed.FS
