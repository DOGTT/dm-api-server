
SELECT setval(pg_get_serial_sequence('user_info', 'id'), 200000001);
SELECT setval(pg_get_serial_sequence('pet_info', 'id'), 100000001);


INSERT INTO pofp_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (1, '探险', 50, '#FF0000', '2024-05-06 10:00:00', '2024-05-06 10:00:00');
INSERT INTO pofp_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (2, '小憩', 30, '#00FF00', '2024-05-06 10:05:00', '2024-05-06 10:05:00');
INSERT INTO pofp_type_info (id, name, coverage_radius, theme_color, created_at, updated_at) 
VALUES (3, '溜溜', 20, '#0000FF', '2024-05-06 10:10:00', '2024-05-06 10:10:00');
