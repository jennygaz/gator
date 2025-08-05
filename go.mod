module github.com/jennygaz/gator

go 1.24.0

replace github.com/jennygaz/gator v0.0.0 => ../gator

require github.com/jennygaz/gator/internal/config v0.0.0-20250803025117-3133927e77db

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
)
