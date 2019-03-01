### Report trello tools ###

- Chức năng 1: Xem các card đang được tạo trong 1 tuần trên list review-me và done.
    -> Chạy router "/b/cards/review/:id".
    Chỉ có 3 card được tạo trong vòng 1 tuần trở lại đây.

- Chức năng 2: Xem các card đã bị thay đổi ngày hết hạn trên bảng.
    -> Chạy router "/b/cards/changeddue/:id".



** Chức năng phụ:
    - C1: Save data (có data mới sử dụng được chức năng 2): Lưu dữ liệu vào database.
        Chỉ chạy 1 lần duy nhất khi khởi động hệ thống.
        -> Chạy router "/b/cards/save/:id"
        Do đã save data từ trước nên không chạy lại nữa.
    
    - C2: Update date: Cập nhật dữ liệu card lên database.
        -> Chạy router "/b/cards/update/:id"
        Update ngày hết hạn lên database.


## Xong ##