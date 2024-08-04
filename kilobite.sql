-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 04, 2024 at 03:16 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kilobite`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `short_description` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `perks` text DEFAULT NULL,
  `backer_count` int(11) DEFAULT NULL,
  `goal_amount` int(11) DEFAULT NULL,
  `current_amount` int(11) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'Campaign 1', 'short description', 'looooooonggg', 'ini satu, ini dua, dan ini 3', NULL, 10000000, 0, NULL, '2024-07-03 19:54:16', '2024-07-03 19:54:17'),
(2, 1, 'Penggalangan Dana Digital', 'Penggalangan Dana Digital', 'Penggalangan Dana Digital', 'Hadiah satu, dua, dan Tiga', 0, 100000000, 0, 'penggalangan-dana-digital-s-int-1', '2024-07-07 15:03:55', '2024-07-07 15:03:55'),
(3, 2, 'sebuah campaign yang keren', 'Penjelasan yang panjang', 'Penjelasan yang panjang', 'keuntungan satu, kemudian yang dua, dan tiga', 0, 1000000, 0, 'sebuah-campaign-yang-keren-2', '2024-07-07 15:45:25', '2024-07-07 15:45:25'),
(4, 2, 'sebuah campaign yang keren', 'Penjelasan yang panjang', 'Penjelasan yang panjang', 'keuntungan satu, kemudian yang dua, dan tiga', 0, 1000000, 0, 'sebuah-campaign-yang-keren-2', '2024-07-07 15:46:38', '2024-07-07 15:46:38'),
(5, 2, 'sebuah campaign yang keren', 'Penjelasan yang panjang', 'Penjelasan yang panjang', 'keuntungan satu, kemudian yang dua, dan tiga', 0, 1000000, 0, 'sebuah-campaign-yang-keren-2', '2024-07-07 15:47:02', '2024-07-07 15:47:02');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `is_primary` tinyint(4) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'images/2-square-gopher.png', 1, '2024-07-15 20:04:25', '2024-07-15 20:04:25'),
(2, 2, 'images/1-square-gopher.png', 1, '2024-07-15 20:18:03', '2024-07-15 20:18:03');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 50000, 'pending', 'trx01', '2024-07-26 18:05:22', '2024-07-26 18:05:23'),
(2, 1, 1, 10000, 'paid', 'trx01', '2024-07-26 18:06:08', '2024-07-26 18:06:08'),
(3, 3, 2, 20000, 'pending', 'trx02', '2024-07-26 18:45:56', '2024-07-26 18:45:57'),
(4, 2, 2, 10000000, 'pending', '', '2024-08-02 18:38:27', '2024-08-02 18:38:27'),
(5, 2, 2, 10000000, 'pending', '', '2024-08-02 19:25:16', '2024-08-02 19:25:16'),
(6, 2, 1, 1000000, 'pending', '', '2024-08-02 19:27:30', '2024-08-02 19:27:30');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(1, 'Jono', 'ux', 'bayu@gmail.com', '$2a$04$EU4AsEQSD8iiSeUOiZrXhO7Bw.YJOFsg2lpZoWSLV7VjTncqfqRAq', 'images/1-Screenshot 2024-06-26 064245.jpg', 'user', NULL, '2024-07-03 16:52:09', '2024-07-03 18:16:39'),
(2, 'Bang Bayu', 'ux', 'bayuaaji@gmail.com', '$2a$04$lEMofjtHdJhLNGTcO8.dLe15Oy.d.SbSNh.mt4uXTXWOAto4YNWOC', 'images/2-Screenshot 2024-06-26 064245.jpg', 'user', NULL, '2024-07-03 18:20:08', '2024-07-03 18:43:38');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
