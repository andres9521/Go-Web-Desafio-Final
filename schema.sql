DROP TABLE IF EXISTS `turnos`;

CREATE TABLE `turnos` (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `descripcion` varchar(45) DEFAULT NULL,
                                `pacienteId` int DEFAULT NULL,
                                `dentistaId` int DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                KEY `patient_idx` (`patient`),
                                KEY `dentist_idx` (`dentist`),
                                CONSTRAINT `dentista` FOREIGN KEY (`dentistaId`) REFERENCES `dentistas` (`id`),
                                CONSTRAINT `paciente` FOREIGN KEY (`pacienteId`) REFERENCES `pacientes` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


DROP TABLE IF EXISTS `dentistas`;

CREATE TABLE `dentistas` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `nombre` varchar(45) DEFAULT NULL,
                            `apellido` varchar(45) DEFAULT NULL,
                            `matricula` varchar(45) DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `id_UNIQUE` (`id`),
                            UNIQUE KEY `matricula_UNIQUE` (`matricula`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `pacientes`;

CREATE TABLE `pacientes` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `nombre` varchar(45) DEFAULT NULL,
                            `apellido` varchar(45) DEFAULT NULL,
                            `domicilio` varchar(45) DEFAULT NULL,
                            `dni` varchar(45) DEFAULT NULL,
                            `fechaDeAlta` varchar(45) DEFAULT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `dentistas` (`nombre`, `apellido`, `matricula`) VALUES
                                                          ('Juan', 'García', '1234567A'),
                                                          ('María', 'Martínez', '2345678B'),
                                                          ('Carlos', 'González', '3456789C');

INSERT INTO `pacientes` (`nombre`, `apellido`, `dni`, `domicilio`, `fechaDeAlta`) VALUES
                                                                         ('Laura', 'Gómez', '12345678A', 'Calle Mayor 1', '1990-10-15'),
                                                                         ('Pedro', 'Hernández', '23456789B', 'Plaza España 2', '1985-07-22'),
                                                                         ('Ana', 'Pérez', '34567890C', 'Avenida Libertad 3', '1995-04-03');

INSERT INTO `appointments` (`description`, `pacienteId`, `dentistaId`) VALUES
                                                                             ('Limpieza dental', 1, 1),
                                                                             ('Extracción muela del juicio', 2, 2),
                                                                             ('Empaste dental', 3, 1);
