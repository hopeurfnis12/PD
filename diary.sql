-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Май 02 2022 г., 00:02
-- Версия сервера: 8.0.24
-- Версия PHP: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `diary`
--

-- --------------------------------------------------------

--
-- Структура таблицы `subjects`
--

CREATE TABLE `subjects` (
  `id` int UNSIGNED NOT NULL,
  `subject` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `subjects`
--

INSERT INTO `subjects` (`id`, `subject`) VALUES
(29, 'Информатика'),
(30, 'Математика'),
(31, 'Физика');

-- --------------------------------------------------------

--
-- Структура таблицы `todo_list`
--

CREATE TABLE `todo_list` (
  `id` int UNSIGNED NOT NULL,
  `subject_id` int UNSIGNED NOT NULL,
  `todo` text,
  `do` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `todo_list`
--

INSERT INTO `todo_list` (`id`, `subject_id`, `todo`, `do`) VALUES
(83, 29, 'Изучить JS', 1),
(85, 29, 'Изучить C#', 1),
(86, 29, 'Изучить C++', 1),
(87, 29, 'Изучить Java', 0),
(88, 30, 'Решить упражнение 21', 1),
(89, 30, 'Решить упражнение 29', 1),
(92, 29, 'Изучить Golang', 1),
(93, 29, 'Сделать проект по Golang\r\n', 1),
(94, 29, 'Сделать курсовую', 0);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `subjects`
--
ALTER TABLE `subjects`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `todo_list`
--
ALTER TABLE `todo_list`
  ADD PRIMARY KEY (`id`),
  ADD KEY `subject_id` (`subject_id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `subjects`
--
ALTER TABLE `subjects`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT для таблицы `todo_list`
--
ALTER TABLE `todo_list`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=95;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `todo_list`
--
ALTER TABLE `todo_list`
  ADD CONSTRAINT `todo_list_ibfk_1` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
