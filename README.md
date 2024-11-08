# Business Logic

### 1. **User Service**
   - **Chức năng chính**:
     - **Đăng ký tài khoản**: Khi người dùng chưa có tài khoản, họ có thể tạo tài khoản mới qua API của User Service. Service này lưu trữ thông tin cơ bản của người dùng như username, email, và hashed password
     - **Quản lý thông tin người dùng**: Cập nhật thông tin cá nhân hoặc thay đổi mật khẩu
   - **Công nghệ và cách triển khai**:
     - **Cơ sở dữ liệu**: Sử dụng PostgreSQL hoặc MongoDB (nếu cần lưu trữ động) để lưu trữ thông tin người dùng
     - **Mã hóa mật khẩu**: Sử dụng thư viện mã hóa như `bcrypt` để hash mật khẩu trước khi lưu vào database
     - **API**: Thiết kế các API RESTful cho đăng ký và quản lý tài khoản
   - **Triển khai**: Đưa User Service vào file `docker-compose.yml` để chạy cùng với các service khác trong hệ thống

### 2. **Relationship between user-service and auth-service**
1. **Đăng ký**: Người dùng gọi API của User Service để tạo tài khoản mới
2. **Login**: Sau khi đăng ký, người dùng đăng nhập qua Authentication-Authorization Service và nhận được token
3. **Sử dụng token**: Khi gọi API từ các dịch vụ khác (vd: Order Service để lấy danh sách đơn hàng), token sẽ được gửi kèm trong request headers. Các service sẽ dùng Authentication-Authorization Service để xác thực token và quyết định cấp quyền truy cập

### 3. **Ưu điểm của thiết kế này**
- **Scalability**: User Service và Authentication-Authorization Service tách biệt giúp hệ thống mở rộng dễ dàng khi số lượng người dùng tăng
- **Security**: Mật khẩu được lưu trữ an toàn bằng cách mã hóa, và token có thời hạn cụ thể
- **Flexibility**: Khi cần, bạn có thể thêm third-party auth (như Google OAuth) hoặc mở rộng JWT với các quyền chi tiết hơn

# Cấu hình các Service và Route trong Kong

1. **Tạo Service** trong Kong cho từng dịch vụ backend (gRPC gateway endpoint)

   Ví dụ: nếu bạn có một dịch vụ `BookService` và `grpc-gateway` đã tạo các endpoint HTTP tại `http://user:8083`, bạn có thể cấu hình như sau:

   ```bash
   curl -i -X POST http://localhost:8001/services \
     --data "name=user-service" \
     --data "url=http://user:8083"
   ```
   Lấy danh sách các service đã tạo:

   ```bash
   curl -i -X GET http://localhost:8001/services
   ```

2. **Tạo Route** để xác định đường dẫn cho từng endpoint HTTP mà bạn muốn nhận từ Kong.

   ```bash
   curl -i -X POST http://localhost:8001/services/user-service/routes \
     --data "paths[]=/v1/users" \
     --data "strip_path=false"
   ```

   Thao tác này sẽ cấu hình một route `/v1/users` trong Kong. Khi có một request đến `http://kong-host:8000/v1/users`, Kong sẽ chuyển tiếp tới `http://<grpc-gateway>:8083/v1/users`

   Lấy danh sách các routes của 1 service cụ thể:

   ```bash
   curl -i -X GET http://localhost:8001/services/${SERVICE_NAME}/routes
   ```

3. **Tùy chỉnh các route** để xác định các endpoint cụ thể nếu bạn có nhiều phương thức trong `UserService` (như `Create`, `GetUser`, `ListUsers`)

   Ví dụ: để định tuyến `GET /v1/users/{user_id}` tới `grpc-gateway`:
   ```bash
   curl -i -X POST http://localhost:8001/services/user-service/routes \
     --data "paths[]=/v1/users/{user_id}" \
     --data "methods[]=GET"
   ```

4. Xóa route theo ID

   ```bash
   curl -i -X DELETE http://localhost:8001/routes/${ROUTE_ID}
   ```
