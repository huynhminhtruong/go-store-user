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

### **Relationship between user-service and auth-service**
1. **Đăng ký**: Người dùng gọi API của User Service để tạo tài khoản mới
2. **Login**: Sau khi đăng ký, người dùng đăng nhập qua Authentication-Authorization Service và nhận được token
3. **Sử dụng token**: Khi gọi API từ các dịch vụ khác (vd: Order Service để lấy danh sách đơn hàng), token sẽ được gửi kèm trong request headers. Các service sẽ dùng Authentication-Authorization Service để xác thực token và quyết định cấp quyền truy cập

### **Ưu điểm của thiết kế này**
- **Scalability**: User Service và Authentication-Authorization Service tách biệt giúp hệ thống mở rộng dễ dàng khi số lượng người dùng tăng
- **Security**: Mật khẩu được lưu trữ an toàn bằng cách mã hóa, và token có thời hạn cụ thể
- **Flexibility**: Khi cần, bạn có thể thêm third-party auth (như Google OAuth) hoặc mở rộng JWT với các quyền chi tiết hơn

Thiết kế này giúp bạn có một nền tảng bảo mật, linh hoạt và dễ dàng mở rộng trong tương lai
