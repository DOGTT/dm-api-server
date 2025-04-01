
SELECT setval(pg_get_serial_sequence('user_info', 'id'), 100000000);
SELECT setval(pg_get_serial_sequence('pet_info', 'id'), 500000000);


INSERT INTO channel_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (1, '跑跑', 50, '#FF0000', '2024-05-06 10:00:00', '2024-05-06 10:00:00');
INSERT INTO channel_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (2, '小憩', 5, '#00FF00', '2024-05-06 10:05:00', '2024-05-06 10:05:00');
INSERT INTO channel_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (3, '露营', 500, '#0000FF', '2024-05-06 10:10:00', '2024-05-06 10:10:00');
INSERT INTO channel_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (4, '远足', 500, '#0000FF', '2024-05-06 10:10:00', '2024-05-06 10:10:00');
INSERT INTO channel_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (5, '探险', 500, '#0000FF', '2024-05-06 10:10:00', '2024-05-06 10:10:00');
