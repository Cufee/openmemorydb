module byvko.dev/repo/openmemorydb

go 1.17

replace aftermath.link/repo/logs => ../logs

require (
	aftermath.link/repo/logs v0.0.0-00010101000000-000000000000
	github.com/boltdb/bolt v1.3.1
	github.com/gofiber/fiber/v2 v2.23.0
	github.com/google/uuid v1.3.0
)

require (
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.31.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)
