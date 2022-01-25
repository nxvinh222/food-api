# GoEx4
Vì sao trong khoá học này các bạn được khuyên không nên dùng khoá ngoại (FK), điểm yếu của khoá ngoại là gì?
## Trả lời
Lý do không sử dụng FK vì mỗi lần insert, update dữ liệu vào bảng, database sẽ phải kiểm tra trường dữ liệu đó tại bảng được reference đến có tồn tại hay không. Điều này dẫn đến việc lãng phí tài nguyên tính toán đối với các dữ liệu không yêu cầu quá cao về tính toàn vẹn.
