# GoEx6
### Index trong table DB có công dụng gì?
Index trong table DB có công dụng đưa dữ liệu trong trường được đánh Index vào một cấu trúc dữ liệu (BTREE, HASHMAP). Khi nhận được truy vấn, Database sẽ thông qua Index để tìm dữ liệu mong muốn.
### Nếu không có index thì khác biệt trong truy vấn là gì?
Index tuy là sẽ tăng tốc độ truy vấn nhưng có nhược điểm là
- Tốn bộ nhớ, Database cần bộ nhớ để lưu cấu trúc dữ liệu của index.
- Giảm tốc độ của các truy vấn INSERT, UPDATE, DELETE do mỗi lần dữ liệu thay đổi, Database phải đánh lại Index.
Vậy tóm lại nên đánh Index ở các trường dữ liệu hay thực hiện truy vấn đọc với WHERE, hạn chế đánh Index ở các trường dữ liệu thay đổi liên tục.
