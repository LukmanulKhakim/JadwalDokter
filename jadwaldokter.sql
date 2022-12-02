-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: localhost    Database: jadwaldokter
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `dokters`
--

DROP TABLE IF EXISTS `dokters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dokters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nama_dokter` longtext,
  `poli_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_dokters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dokters`
--

LOCK TABLES `dokters` WRITE;
/*!40000 ALTER TABLE `dokters` DISABLE KEYS */;
INSERT INTO `dokters` VALUES (1,'2022-12-01 17:01:16.783','2022-12-01 17:01:16.783',NULL,'dr. Aida Musyarrofah, Sp.OG',1),(2,'2022-12-01 17:01:32.235','2022-12-01 17:01:32.235',NULL,' dr. Halida Nelasari, Sp.OG(K)',1),(3,'2022-12-01 17:01:48.517','2022-12-01 17:01:48.517',NULL,'dr. Kusuma Andriana, Sp.OG',1),(4,'2022-12-01 17:02:03.668','2022-12-01 17:02:03.668',NULL,'dr. Moch. Maroef, Sp.OG',1),(5,'2022-12-01 17:02:27.698','2022-12-01 17:02:27.698',NULL,'dr. Dicky Faturrachman, Sp.A., M.Biomed',2),(6,'2022-12-01 17:02:43.847','2022-12-01 17:02:43.847',NULL,'Dr. dr Renny Suwarnity Sp.A(K)',2),(7,'2022-12-01 17:03:01.950','2022-12-01 17:03:01.950',NULL,'dr. Hawin Nurdiana, Sp.A., M.Kes',2),(8,'2022-12-02 08:32:42.432','2022-12-02 08:32:42.432',NULL,'dr. Husnul Asariati, Sp.A., M.Biomed',2),(9,'2022-12-02 08:32:58.583','2022-12-02 08:32:58.583',NULL,'dr. Pertiwi Febriana, M.Sc., Sp.A',2),(10,'2022-12-02 08:33:21.312','2022-12-02 08:33:21.312',NULL,'Dini Fidyanti Devi, M.Psi., Psikolog',3),(11,'2022-12-02 08:33:39.738','2022-12-02 08:33:39.738',NULL,'Wulida Azmiyya. E. R., M.Psi., Psikolog',3),(12,'2022-12-02 08:33:56.203','2022-12-02 08:33:56.203',NULL,'dr. Marintik Ilahi, Sp.KJ',4),(13,'2022-12-02 08:34:08.760','2022-12-02 08:34:08.760',NULL,'dr. Achmad Rifai, Sp.PD, FINASIM',4),(14,'2022-12-02 08:35:16.451','2022-12-02 08:35:16.451',NULL,'dr. Ardhi Bustami, Sp.PD',5),(15,'2022-12-02 08:35:31.190','2022-12-02 08:35:31.190',NULL,'dr. Gerry Permadi, Sp.PD',5),(16,'2022-12-02 08:35:46.709','2022-12-02 08:35:46.709',NULL,'dr. Isbandiyah, Sp.PD',5),(17,'2022-12-02 08:35:58.724','2022-12-02 08:35:58.724',NULL,'dr. Nina Nur Arifah, Sp.PD',5),(18,'2022-12-02 08:36:13.210','2022-12-02 08:36:13.210',NULL,'Dr.dr. Meddy Setiawan Sp.PD',5);
/*!40000 ALTER TABLE `dokters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jadwals`
--

DROP TABLE IF EXISTS `jadwals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jadwals` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `hari` longtext,
  `jam` longtext,
  `dokter_id` bigint unsigned DEFAULT NULL,
  `poli_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_jadwals_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jadwals`
--

LOCK TABLES `jadwals` WRITE;
/*!40000 ALTER TABLE `jadwals` DISABLE KEYS */;
INSERT INTO `jadwals` VALUES (1,'2022-12-01 18:40:49.381','2022-12-01 18:40:49.381',NULL,'senin','14.00-16.00',2,1),(2,'2022-12-01 18:41:26.098','2022-12-01 18:41:26.098',NULL,'senin','10.00-12.00',3,1),(3,'2022-12-01 23:30:13.123','2022-12-01 23:30:13.123',NULL,'selasa','18.00-21.00',1,1),(4,'2022-12-01 23:31:10.853','2022-12-01 23:31:10.853',NULL,'selasa','09.00-11.00',2,1),(5,'2022-12-02 08:40:07.973','2022-12-02 08:40:07.973',NULL,'senin','10:00-11:30',5,2),(6,'2022-12-02 08:40:34.039','2022-12-02 08:40:34.039',NULL,'senin','15:00-17:00',7,2),(7,'2022-12-02 08:40:52.837','2022-12-02 08:40:52.837',NULL,'senin','18:00-19:30',8,2),(8,'2022-12-02 08:41:20.690','2022-12-02 08:41:20.690',NULL,'senin','12:30-14:30',9,2),(9,'2022-12-02 08:42:10.116','2022-12-02 08:42:10.116',NULL,'selasa','10:00-11:30 ',5,2),(10,'2022-12-02 08:42:26.086','2022-12-02 08:42:26.086',NULL,'selasa','18:00-19:30 ',5,2),(11,'2022-12-02 08:42:58.186','2022-12-02 08:42:58.186',NULL,'selasa','09:00-11:00 ',6,2),(12,'2022-12-02 08:43:17.948','2022-12-02 08:43:17.948',NULL,'selasa','15:00-17:00 ',8,2),(13,'2022-12-02 08:43:48.843','2022-12-02 08:43:48.843',NULL,'selasa','11:30-13:30 ',9,2),(14,'2022-12-02 08:44:19.557','2022-12-02 08:44:19.557',NULL,'rabu','18:00-21:00 ',1,1),(15,'2022-12-02 08:44:50.127','2022-12-02 08:44:50.127',NULL,'rabu','10:00-12:00 ',3,1),(16,'2022-12-02 08:45:24.284','2022-12-02 08:45:24.284',NULL,'kamis','08:00-12:00 ',1,1),(17,'2022-12-02 08:46:01.828','2022-12-02 08:46:01.828',NULL,'kamis','13:00-15:00',2,1),(18,'2022-12-02 08:46:46.463','2022-12-02 08:46:46.463',NULL,'kamis','09:00-11:00',6,2);
/*!40000 ALTER TABLE `jadwals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `polis`
--

DROP TABLE IF EXISTS `polis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `polis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nama_poli` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_polis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `polis`
--

LOCK TABLES `polis` WRITE;
/*!40000 ALTER TABLE `polis` DISABLE KEYS */;
INSERT INTO `polis` VALUES (1,'2022-12-01 11:21:34.273','2022-12-01 11:21:34.273',NULL,'Spesialis Kebidanan dan Kandungan /Obstetrician & Gynecologist-Dokter Obsgyn'),(2,'2022-12-01 11:22:08.503','2022-12-01 11:22:08.503',NULL,'Spesialis Anak / Pediatrician'),(3,'2022-12-01 11:23:10.962','2022-12-01 11:23:10.962',NULL,'Klinik Psikologi / Psikolog'),(4,'2022-12-02 08:28:46.837','2022-12-02 08:28:46.837',NULL,'Spesialis Kesehatan Jiwa / Psychiatrist'),(5,'2022-12-02 08:29:04.285','2022-12-02 08:29:04.285',NULL,'Spesialis Penyakit Dalam / Internist'),(6,'2022-12-02 08:29:17.541','2022-12-02 08:29:17.541',NULL,'Spesialis Jantung dan Pembuluh Darah / Cardiologist'),(7,'2022-12-02 08:29:35.721','2022-12-02 08:29:35.721',NULL,'Spesialis Paru / Pulmonologist'),(8,'2022-12-02 08:29:55.312','2022-12-02 08:29:55.312',NULL,'Spesialis Saraf / Neurologist'),(9,'2022-12-02 08:30:27.305','2022-12-02 08:30:27.305',NULL,'Gigi / Dentist umum'),(10,'2022-12-02 08:30:43.402','2022-12-02 08:30:43.402',NULL,'Spesialis Jaringan Penyangga Gigi / Periodontist'),(11,'2022-12-02 08:30:55.841','2022-12-02 08:30:55.841',NULL,'Spesialis Penyakit Mulut / Oral Medicine');
/*!40000 ALTER TABLE `polis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'jadwaldokter'
--

--
-- Dumping routines for database 'jadwaldokter'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-02  9:30:27
