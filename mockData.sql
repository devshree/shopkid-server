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

-- Insert sample users
INSERT INTO users (name, email, password) VALUES
    ('John Doe', 'john@example.com', 'hashed_password_1'),
    ('Jane Smith', 'jane@example.com', 'hashed_password_2'),
    ('Mike Johnson', 'mike@example.com', 'hashed_password_3'),
    ('Sarah Wilson', 'sarah@example.com', 'hashed_password_4'),
    ('Tom Brown', 'tom@example.com', 'hashed_password_5');

-- Insert sample cart items
INSERT INTO cart_items (product_id, quantity, price) VALUES
    (1, 2, 59.98),
    (3, 1, 15.99),
    (5, 1, 39.99),
    (7, 3, 68.97),
    (9, 2, 69.98);

-- Insert sample orders
INSERT INTO orders (user_id, product_id, quantity, price) VALUES
    (1, 2, 1, 45.99),
    (2, 4, 2, 49.98),
    (3, 6, 1, 19.99),
    (4, 8, 1, 49.99),
    (5, 10, 2, 59.98);
