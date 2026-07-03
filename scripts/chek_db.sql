-- Проверка: количество записей в каждой таблице
SELECT 'categories' AS table_name, COUNT(*) AS count FROM categories
UNION ALL
SELECT 'faq_articles', COUNT(*) FROM faq_articles
UNION ALL
SELECT 'services', COUNT(*) FROM services
UNION ALL
SELECT 'support_tickets', COUNT(*) FROM support_tickets
UNION ALL
SELECT 'v_unresolved_tickets (view)', COUNT(*) FROM v_unresolved_tickets;

-- Просмотр нерешенных обращений (с количеством дней)
SELECT * FROM v_unresolved_tickets;


