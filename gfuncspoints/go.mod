module github.com/host1812/gtutor-practical/gfuncspoints

go 1.19

require (
	gjson v1.0.0
  gjson/objects v1.0.0
)

replace (
	gjson v1.0.0 => ../gjson
  gjson/objects v1.0.0 => ../gjson/objects
)