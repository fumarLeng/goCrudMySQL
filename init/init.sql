-- 建立 users 資料表
CREATE TABLE IF NOT EXISTS users (
                                     id INT AUTO_INCREMENT PRIMARY KEY,
                                     name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

-- 插入假資料
INSERT INTO users (name, email, created_at, updated_at) VALUES
                                                            ('Leo', 'leo@example.com', '2025-05-13 12:00:00', '2025-05-13 12:30:00'),
                                                            ('Andy Lau', 'andy.lau@example.com', '2025-05-12 10:00:00', '2025-05-12 10:30:00'),
                                                            ('Eason Chan', 'eason.chan@example.com', '2025-05-11 09:00:00', '2025-05-11 09:30:00'),
                                                            ('Jay Chou', 'jay.chou@example.com', '2025-05-10 08:00:00', '2025-05-10 08:30:00'),
                                                            ('JJ Lin', 'jj.lin@example.com', '2025-05-09 07:00:00', '2025-05-09 07:30:00');