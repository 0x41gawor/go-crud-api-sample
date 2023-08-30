CREATE DATABASE IF NOT EXISTS `go_crud_api_sample_db`;
USE `go_crud_api_sample_db`;

DROP TABLE IF EXISTS `continents`;
DROP TABLE IF EXISTS `countries`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
`id` int NOT NULL auto_increment,
`login` varchar(45) NOT NULL unique,
`password` varchar(60) NOT NULL,
primary key (`id`)
) engine=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE `continents` (
`id` int NOT NULL auto_increment,
`name` varchar(45) NOT NULL unique,
`population` float NOT NULL, # [bilions]
`gdp` float NOT NULL, # [trilions USD] #PPP
`gdp_per_capita` float NOT NULL, # [k USD] 
primary key (`id`)
) engine=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE `countries` (
`id` int NOT NULL auto_increment,
`name` varchar(45) NOT NULL unique,
`continent_id` int NOT NULL,
`population` float NOT NULL,
`gdp` float NOT NULL,
`gdp_per_capita` float NOT NULL,
`top5cities` varchar(45) NOT NULL,

primary key (`id`),

KEY `FK_continent_idx` (`continent_id`),
CONSTRAINT `FK_continent`
FOREIGN KEY (`continent_id`)
REFERENCES `continents` (`id`)

on delete no action on update no action
) engine=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;


INSERT INTO `continents` VALUES
	(0, 'FF', 78.83, 500, 2);
    
INSERT INTO `countries` VALUES
	(0, 'France', 1, 78.83, 500, 2, 'Paris, Monaco, Marseille, Toulouse, Lyon');
    
select * from `users`;
select * from `continents`;
select * from `countries`;