# Tổ chức source code trong go (package and module)
```
hoc-golang/
├── go.mod
├── main.go
├── cat/
│   └── cat.go
└── dog/
    └── dog.go
```
- khởi tạo module trong go : go mod init "ten module"
- hàm chạy chứa package là main và function bắt đầu compile và chạy cũng là main
- Để tạo và import các package khác thì lúc import ta dùng  "ten module"/{tên package}
- để import các thư built-in ta dùng import "tên thư viện"
- còn import thư viện ngoài chạy  go get  tên thư viên (vd  go get  gopkg.in/mgo.v2 )
## Lưu ý
- Qui ước trong go : 
  - function, struct, variable muốn import từ module khác được cần viết hoa chữ đầu
  - Nếu viết thường thì chỉ truy xuất được trong package đó thôi