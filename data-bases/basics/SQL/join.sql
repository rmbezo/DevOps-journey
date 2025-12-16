SELECT u.name, o.product
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;