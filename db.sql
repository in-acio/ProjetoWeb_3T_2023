-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Tempo de geração: 28-Out-2023 às 14:41
-- Versão do servidor: 10.4.22-MariaDB
-- versão do PHP: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Banco de dados: `prog3`
--

-- --------------------------------------------------------

--
-- Estrutura da tabela `items`
--

CREATE TABLE `items` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(256) DEFAULT NULL,
  `img` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `items`
--

INSERT INTO `items` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `img`) VALUES
(1, '2023-10-28 09:41:02', '2023-10-28 09:41:02', NULL, 'Aguia', 'crOWDrFWwSyJoIIED0E8oLLpreFzOH.jpg'),
(2, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Arara', 'Vkl2yjjFvmO76qDao1NLMzVY6m9YWk.webp'),
(3, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Baleia', '3LmnaSBF3P8HlXazwFsrQD9AIpajMw.jpg'),
(4, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Cachorro', 'GZakWIeH9Vqoic8fEzquoAxCO2AgNA.jpg'),
(5, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Canguru', 'mdg30d1cSwOC8ub2p4jPO0BbvRD9tC.webp'),
(6, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Cavalo-marinho', 'k1mqomyaxfwhrkSup8NRaaY6ZHq5PW.jpg'),
(7, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Cobra', 'dD7k17fGIUn57WlkJSgjQDKn8BeHTq.jpg'),
(8, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Crocodilo', 'HGp7lbfyjPv4Qpf3iccMqJ36Qn6Nuu.jpg'),
(9, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Elefante', 'KmJHyYd41ghNGCQFpvlux0OyX9d7xX.jpg'),
(10, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Escorpiao', '0dozELqDoP5GlBmj0p47mI6y8BcVTY.jpg'),
(11, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Gato', 'RoTfsh0TqOeU8fxjPQpgODzIObtwcu.jpg'),
(12, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Lagarto', 'MghajavpGTJZF5O6FTlr6w9MaNnkfc.jpg'),
(13, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Leao', 'vR2rFDwlEsVpNWQJ07OLLYfGECEYI3.jpg'),
(14, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Lobo-Guara', '3castTYcagZCH28onpcm4JKnpaqk4f.webp'),
(15, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Mosquito', 'OEV9eYQ0A7SPyJ9OcPEyG22LrjxMXA.jpg'),
(16, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Onca', 'ZpV67x2o1b5s3R6VR2Oa8ArqShSuAT.JPG'),
(17, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Peixe', 'nJryPDT8CtSkkhDuKuD5JqdKRJ4IzG.webp'),
(18, '2023-10-28 09:41:03', '2023-10-28 09:41:03', NULL, 'Penguim', 'D4DKbW3PqNQT8UtGAJCTce7j0M5gW0.jpg'),
(19, '2023-10-28 09:41:04', '2023-10-28 09:41:04', NULL, 'Pomba', 'htBv1Hszs6FcHpShB4zffh0aH5N9f9.jpg'),
(20, '2023-10-28 09:41:04', '2023-10-28 09:41:04', NULL, 'Rato', 'OLZNysgUGALvVKA7YLyx4TxS3na5xr.webp'),
(21, '2023-10-28 09:41:04', '2023-10-28 09:41:04', NULL, 'Tartaruga', 'V42HnhUiP6pzoE02Ctkobw4EY5DsO2.jpg'),
(22, '2023-10-28 09:41:04', '2023-10-28 09:41:04', NULL, 'Tatu', 'GvXx2IerVmjJWXzvGp2vJ2M4MSe67D.jpg'),
(23, '2023-10-28 09:41:04', '2023-10-28 09:41:04', NULL, 'Tigre', '2f0kM1WyUe9Ync7VNUMEt6x4RD3XtW.webp'),
(24, '2023-10-28 09:41:04', '2023-10-28 09:41:04', NULL, 'Tucano', 'JQQkxibExv5eMqKjaQXRWwlLXOBb2r.webp');

-- --------------------------------------------------------

--
-- Estrutura da tabela `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(256) DEFAULT NULL,
  `email` varchar(256) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Extraindo dados da tabela `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `password`, `is_admin`) VALUES
(1, '2023-10-28 09:39:04', '2023-10-28 09:39:04', NULL, 'admin', 'admin', '$2a$10$kwe8cxN2OCPUg.td8k/9weqQVrhQlhVPaTo7XH8oW6.2syJ55uXb2', 1);

-- --------------------------------------------------------

--
-- Estrutura da tabela `votes`
--

CREATE TABLE `votes` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `item_id` bigint(20) UNSIGNED DEFAULT NULL,
  `value` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Índices para tabelas despejadas
--

--
-- Índices para tabela `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_items_deleted_at` (`deleted_at`);

--
-- Índices para tabela `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- Índices para tabela `votes`
--
ALTER TABLE `votes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_votes_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT de tabelas despejadas
--

--
-- AUTO_INCREMENT de tabela `items`
--
ALTER TABLE `items`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT de tabela `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT de tabela `votes`
--
ALTER TABLE `votes`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
