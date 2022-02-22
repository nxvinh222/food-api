### Nếu chúng ta có nhiều hơn 1 module, làm sao để giao tiếp với nhau
- Tại tầng business ta thực hiện gọi tới store của các module khác
### Giả sử module "Restaurant" cần data số lượt like từ module "Like Restaurant" thì sẽ truy xuất như thế nào?
- Tại business của Restaurant ta gọi store của RestaurantLike
