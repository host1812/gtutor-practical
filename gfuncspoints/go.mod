module github.com/host1812/gtutor-practical/gfuncspoints

go 1.19

require github.com/host1812/gtutor-practical/gjson/objects v1.0.0

replace (
	github.com/host1812/gtutor-practical/gjson v1.0.0 => ../gjson
	github.com/host1812/gtutor-practical/gjson/objects v1.0.0 => ../gjson/objects
)
