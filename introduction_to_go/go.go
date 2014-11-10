i := pivot(s)
go sort(s[:i])
go sort(s[i:])
