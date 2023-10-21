module github.com/jluong102/projectguardian/angel

replace github.com/jluong102/projectguardian/logger => ../../../shared/logger
replace github.com/jluong102/projectguardian/permissions => ../../../shared/permissions

go 1.15

require github.com/jluong102/projectguardian/logger v0.0.0-00010101000000-000000000000
