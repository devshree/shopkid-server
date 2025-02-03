-- Insert sample products
INSERT INTO products (name, description, price, category, age_range, stock, image) VALUES
    ('Wooden Building Blocks', 'Classic wooden blocks for creative play', 29.99, 'toys', '3-8', 50, 'https://images.unsplash.com/photo-1587654780291-39c9404d746b?w=500'),
    ('Kids Art Easel', 'Double-sided art easel with chalkboard', 45.99, 'toys', '4-12', 30, 'https://images.unsplash.com/photo-1617791160505-6f00504e3519?w=500'),
    ('Dinosaur T-Shirt', 'Cotton t-shirt with dinosaur print', 15.99, 'clothes', '4-6', 100, 'https://images.unsplash.com/photo-1529374255404-311a2a4f1fd9?w=500'),
    ('Rainbow Dress', 'Colorful summer dress', 24.99, 'clothes', '3-7', 75, 'https://images.unsplash.com/photo-1567113463300-102a7eb3cb26?w=500'),
    ('STEM Robot Kit', 'Educational robot building kit', 39.99, 'toys', '8-12', 40, 'https://images.unsplash.com/photo-1535378917042-10a22c95931a?w=500'),
    ('Space Pajamas', 'Glow-in-dark space themed PJs', 19.99, 'clothes', '5-8', 60, 'https://images.unsplash.com/photo-1595461135849-bf08893fdc2c?w=500'),
    ('Musical Xylophone', 'Colorful wooden xylophone', 22.99, 'toys', '2-6', 45, 'https://images.unsplash.com/photo-1520523839897-bd0b52f945a0?w=500'),
    ('Winter Jacket', 'Warm waterproof winter jacket', 49.99, 'clothes', '6-10', 80, 'https://images.unsplash.com/photo-1604644401890-0bd678c83788?w=500'),
    ('Magnetic Tiles', 'Magnetic building tiles set', 34.99, 'toys', '3-10', 55, 'https://images.unsplash.com/photo-1596461404969-9ae70f2830c1?w=500'),
    ('Sports Shoes', 'Comfortable athletic shoes', 29.99, 'clothes', '7-12', 90, 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=500');

-- Insert sample users with roles
INSERT INTO users (name, email, password, role) VALUES
    ('John Doe', 'john@example.com', 'hashed_password_1', 'buyer'),
    ('Jane Smith', 'jane@example.com', 'hashed_password_2', 'admin'),
    ('Mike Johnson', 'mike@example.com', 'hashed_password_3', 'purchaser');

-- Insert sample cart items
INSERT INTO cart_items (user_id, product_id, quantity, price) VALUES
    (1, 1, 2, 59.98),
    (2, 2, 1, 45.99),
    (3, 3, 3, 47.97);

-- Insert sample orders
INSERT INTO orders (user_id, total_amount, status) VALUES
    (1, 59.98, 'pending'),
    (2, 45.99, 'completed'),
    (3, 47.97, 'pending');

-- Insert sample order items
INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
    (1, 1, 2, 29.99),
    (2, 2, 1, 45.99),
    (3, 3, 3, 15.99);

-- Insert sample login history
INSERT INTO login_history (user_id, status) VALUES
    ('1', 'success'),
    ('2', 'error'),
    ('1', 'success');
