### Trong trường hợp tạo cột đếm thì làm sao để update cột đó?
- Sử dụng SQL Expression ("<tên cột>+ ?", 1)
### Làm sao để API chính không bị block vì phải update số đếm?
- Sử dụng Goroutine để chạy async job
