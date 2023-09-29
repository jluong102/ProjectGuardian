module github.com/jluong102/projectguardian

replace github.com/jluong102/projectguardian/logger => ./src/shared/logger

go 1.19

require github.com/jluong102/projectguardian/logger v0.0.0-00010101000000-000000000000 // indirect
