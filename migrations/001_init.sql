- CREATE DATABASE college_faq;
-- \c college_faq;

-- 1. Таблица категорий
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- 2. Таблица FAQ статей
CREATE TABLE faq_articles (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Таблица служб (навигатор)
CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    room VARCHAR(20),
    floor INTEGER,
    phone VARCHAR(20),
    manager VARCHAR(100),
    schedule_weekdays VARCHAR(100),
    schedule_saturday VARCHAR(100)
);

-- 4. Таблица обращений
CREATE TABLE support_tickets (
    id SERIAL PRIMARY KEY,
    student_name VARCHAR(100) NOT NULL,
    course INTEGER NOT NULL CHECK (course BETWEEN 1 AND 4),
    question TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'Новый' CHECK (status IN ('Новый', 'Решен')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 5. Представление для нерешенных обращений
CREATE VIEW v_unresolved_tickets AS
SELECT 
    id,
    student_name,
    course,
    question,
    created_at,
    EXTRACT(DAY FROM (CURRENT_TIMESTAMP - created_at)) AS days_passed
FROM support_tickets
WHERE status = 'Новый'
ORDER BY created_at ASC;

-- Комментарии:
COMMENT ON TABLE categories IS 'Категории вопросов';
COMMENT ON TABLE faq_articles IS 'Статьи с ответами на частые вопросы';
COMMENT ON TABLE services IS 'Справочник служб и кабинетов колледжа';
COMMENT ON TABLE support_tickets IS 'Журнал обращений студентов';
COMMENT ON VIEW v_unresolved_tickets IS 'Новые обращения с количеством дней с момента подачи';
