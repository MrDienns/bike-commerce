-- phpMyAdmin SQL Dump
-- version 4.8.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 03, 2019 at 03:52 PM
-- Server version: 10.1.31-MariaDB
-- PHP Version: 7.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `vanderBinckesdb`
--

-- --------------------------------------------------------

--
-- Table structure for table `accessoire`
--

CREATE TABLE `accessoire` (
  `accessoirenummer` int(11) NOT NULL,
  `naam` varchar(50) NOT NULL,
  `huurprijs` decimal(3,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `accessoire`
--

INSERT INTO `accessoire` (`accessoirenummer`, `naam`, `huurprijs`) VALUES
(1, 'regendak', '2.50'),
(2, 'zonnedak', '2.00'),
(3, ' babystoeltje', '3.00'),
(4, 'smart phone houder', '1.00'),
(5, 'kaarthouder', '1.00'),
(6, 'helm', '1.50');

-- --------------------------------------------------------

--
-- Table structure for table `bakfiets`
--

CREATE TABLE `bakfiets` (
  `bakfietsnummer` int(11) NOT NULL,
  `naam` varchar(50) NOT NULL,
  `type` varchar(10) NOT NULL,
  `huurprijs` decimal(4,2) NOT NULL,
  `aantal` int(11) NOT NULL,
  `aantal_verhuurd` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `bakfiets`
--

INSERT INTO `bakfiets` (`bakfietsnummer`, `naam`, `type`, `huurprijs`, `aantal`, `aantal_verhuurd`) VALUES
(1, 'Cargo Bike Green', 'CBG1', '20.00', 10, 2),
(2, 'Cargo Bike Electric', 'CBE1', '40.00', 10, 2);

-- --------------------------------------------------------

--
-- Table structure for table `klant`
--

CREATE TABLE `klant` (
  `klantnummer` int(11) NOT NULL,
  `naam` varchar(45) DEFAULT NULL,
  `voornaam` varchar(15) DEFAULT NULL,
  `postcode` varchar(6) NOT NULL,
  `huisnummer` int(11) NOT NULL,
  `huisnummer_toevoeging` varchar(5) DEFAULT NULL,
  `opmerkingen` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `klant`
--

INSERT INTO `klant` (`klantnummer`, `naam`, `voornaam`, `postcode`, `huisnummer`, `huisnummer_toevoeging`, `opmerkingen`) VALUES
(1, 'Sharp', 'Leo', '1999XB', 10, 'a', NULL),
(2, 'Long', 'George', '1999XB', 84, NULL, NULL),
(3, 'Caouette', 'Coby', '1019BY', 60, NULL, ''),
(4, 'Donatelli', 'Sean', '1019BY', 97, NULL, NULL),
(5, 'Guyer', 'Rachael', '1019BY', 3, NULL, NULL),
(6, 'Perilloux', 'Ike', '6931NX', 76, NULL, NULL),
(7, 'Linhart', 'Ciara', '6931NX', 13, NULL, NULL),
(8, 'Francis', 'Oscar', '2691UQ', 27, NULL, NULL),
(9, 'Stannard', 'Gillian', '6591FR', 77, NULL, NULL),
(10, 'Zapetis', 'Rasmus', '6591FR', 94, NULL, NULL),
(11, 'Anderson', 'Mariska', '6591FR', 54, NULL, NULL),
(12, 'Kuehn', 'Jurre', '2532VL', 26, NULL, ''),
(13, 'Gonzalez', 'Luis', '2532VM', 96, NULL, NULL),
(14, 'Stevens', 'Maja', '7176NU', 66, NULL, NULL),
(15, 'Sterrett', 'Scotty', '7176NU', 25, NULL, NULL),
(16, 'Schubert', 'Charlotte', '7668WA', 16, NULL, NULL),
(17, 'Beckbau', 'Sophia', '7668WA', 64, NULL, NULL),
(18, 'Goodman', 'Rachael', '6173XD', 37, NULL, NULL),
(19, 'Love', 'Luis', '6173XD', 84, NULL, NULL),
(20, 'Hedgecock', 'Sophia', '6173XD', 28, NULL, NULL),
(21, 'Brennan', 'Lara', '3440JV', 64, NULL, NULL),
(22, 'Pengilly', 'Jurre', '3440JV', 52, NULL, ''),
(23, 'Noteboom', 'Vanessa', '7290ZN', 35, NULL, NULL),
(24, 'Daley', 'Barbara', '1952FB', 93, NULL, NULL),
(25, 'Bruno', 'Cloe', '1952FB', 79, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `medewerker`
--

CREATE TABLE `medewerker` (
  `medewerkernummer` int(11) NOT NULL,
  `achternaam` varchar(45) DEFAULT NULL,
  `voornaam` varchar(20) DEFAULT NULL,
  `datum_in_dienst` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `medewerker`
--

INSERT INTO `medewerker` (`medewerkernummer`, `achternaam`, `voornaam`, `datum_in_dienst`) VALUES
(1, 'Jansen', 'Bas', '2018-05-21'),
(2, 'Rademakers', 'Vincent', '2017-09-01'),
(3, 'Heiden, van der', 'Karel', '2018-05-30'),
(4, 'Kraymans', 'Irene', '2017-09-01'),
(5, 'Boer, de', 'Francine', '2018-07-01'),
(6, 'Velzenmaker', 'Jaap', '2017-11-01');

-- --------------------------------------------------------

--
-- Table structure for table `verhuur`
--

CREATE TABLE `verhuur` (
  `verhuurnummer` int(11) NOT NULL,
  `verhuurdatum` date NOT NULL,
  `bakfietsnummer` int(11) NOT NULL,
  `aantal_dagen` int(11) NOT NULL,
  `huurprijstotaal` decimal(5,2) NOT NULL,
  `klantnummer` int(11) NOT NULL,
  `verhuurder` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `verhuur`
--

INSERT INTO `verhuur` (`verhuurnummer`, `verhuurdatum`, `bakfietsnummer`, `aantal_dagen`, `huurprijstotaal`, `klantnummer`, `verhuurder`) VALUES
(1, '2019-12-02', 1, 5, '127.50', 2, 1),
(2, '2019-12-01', 1, 3, '80.00', 4, 3),
(3, '2019-11-30', 2, 7, '140.00', 10, 3);

-- --------------------------------------------------------

--
-- Table structure for table `verhuuraccessoire`
--

CREATE TABLE `verhuuraccessoire` (
  `verhuurnummer` int(11) NOT NULL,
  `accessoirenummer` int(11) NOT NULL,
  `aantal` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `verhuuraccessoire`
--

INSERT INTO `verhuuraccessoire` (`verhuurnummer`, `accessoirenummer`, `aantal`) VALUES
(1, 1, 1),
(1, 6, 2);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `accessoire`
--
ALTER TABLE `accessoire`
  ADD PRIMARY KEY (`accessoirenummer`);

--
-- Indexes for table `bakfiets`
--
ALTER TABLE `bakfiets`
  ADD PRIMARY KEY (`bakfietsnummer`);

--
-- Indexes for table `klant`
--
ALTER TABLE `klant`
  ADD PRIMARY KEY (`klantnummer`);

--
-- Indexes for table `medewerker`
--
ALTER TABLE `medewerker`
  ADD PRIMARY KEY (`medewerkernummer`);

--
-- Indexes for table `verhuur`
--
ALTER TABLE `verhuur`
  ADD PRIMARY KEY (`verhuurnummer`),
  ADD KEY `FK_verhuur_klant` (`klantnummer`),
  ADD KEY `FK_verhuur_medewerker` (`verhuurder`),
  ADD KEY `FK_verhuur_bakfiets` (`bakfietsnummer`);

--
-- Indexes for table `verhuuraccessoire`
--
ALTER TABLE `verhuuraccessoire`
  ADD PRIMARY KEY (`verhuurnummer`,`accessoirenummer`),
  ADD KEY `FK_verhuuracc_accessoire` (`accessoirenummer`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `accessoire`
--
ALTER TABLE `accessoire`
  MODIFY `accessoirenummer` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `bakfiets`
--
ALTER TABLE `bakfiets`
  MODIFY `bakfietsnummer` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `verhuur`
--
ALTER TABLE `verhuur`
  MODIFY `verhuurnummer` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `verhuur`
--
ALTER TABLE `verhuur`
  ADD CONSTRAINT `FK_verhuur_bakfiets` FOREIGN KEY (`bakfietsnummer`) REFERENCES `bakfiets` (`bakfietsnummer`),
  ADD CONSTRAINT `FK_verhuur_klant` FOREIGN KEY (`klantnummer`) REFERENCES `klant` (`klantnummer`),
  ADD CONSTRAINT `FK_verhuur_medewerker` FOREIGN KEY (`verhuurder`) REFERENCES `medewerker` (`medewerkernummer`);

--
-- Constraints for table `verhuuraccessoire`
--
ALTER TABLE `verhuuraccessoire`
  ADD CONSTRAINT `FK_verhuuracc_accessoire` FOREIGN KEY (`accessoirenummer`) REFERENCES `accessoire` (`accessoirenummer`),
  ADD CONSTRAINT `FK_verhuuracc_verhuur` FOREIGN KEY (`verhuurnummer`) REFERENCES `verhuur` (`verhuurnummer`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
