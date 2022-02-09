# GoEx10
### Vì sao không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud.
- Khi database trở nên lớn hơn, dung lượng ảnh/file có thể trở nên vô cùng lớn.
- Lưu ảnh trong chính service làm cho việc quản lý, sắp xếp chúng trở nên khó khăn và phức tạp hơn các cloud thiết kế chuyên dụng cho việc lưu trữ.
- Service nên là stateless, vì vậy nên nó không nên chứa file.
- Dùng các cloud chuyên dụng để lưu ảnh thao tác rất sướng.
### Vì sao không chứa binary ảnh vào DB?
- Dung lượng ảnh có thể trở nên rất lớn, sẽ tốn nhiều chi phí để xử lý chúng(mà cũng không hiệu quả). Vì vậy nên để database tập trung vào việc của nó như tìm kiếm, sắp xếp, xử lý các mối quan hệ,..v.v Và để việc xử lý ảnh cho các cloud chuyên dụng.
- Dung lượng database có thể trở nên rất lớn nên việc backup thường xuyên trở nên khó khăn.
- Trong cơ sở dữ liệu quan hệ, dữ liệu kiểu binary sẽ không thể tạo index hay các mối quan hệ.
- Nếu lưu ảnh dưới dạng binary, service sẽ lại phải tốn thêm công sức để chuyển nó về các dạng file như jpg, png.
