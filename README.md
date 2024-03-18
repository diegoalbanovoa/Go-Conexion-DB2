# Go-Conexion-DB2
Un conector Go robusto para interactuar con bases de datos IBM DB2, lo que permite un acceso y manipulación de datos sin problemas en sus aplicaciones Go, este es un prototipo para demostrar la utilidad de go y su versatilidad en el manejo de un gran volumen de datos.. 

Prototipo de funcionalidades

Este proyecto presenta un prototipo desarrollado para probar diferentes funcionalidades en un contexto similar al de un sistema bancario. El objetivo principal es evaluar el rendimiento del lenguaje Go y su interoperabilidad con la base de datos DB2, sentando las bases para la modernización gradual de un módulo específico.

Lógica de negocio

El prototipo se basa en la calificación A, B o C de los usuarios, la cual determina su acceso a los servicios. La calificación se actualiza en la tabla CALIFIRNEW, que abstrae de forma simplificada los datos de un usuario. La actualización se realiza automáticamente utilizando semillas para replicar el experimento con los mismos datos y garantizar la consistencia de los resultados.

Componentes del prototipo

Generación de usuarios: Se generan usuarios aleatorios con una semilla para garantizar la replicabilidad del experimento. Los usuarios se guardan en archivos CSV.
Generación de registros: Cada usuario tiene 8 registros asociados, también generados aleatoriamente y almacenados en archivos CSV.
Creación de registros en la base de datos: Se crea un usuario en la tabla CALIFIR mediante la API.
Creación de registro de cambio de calificación: Se generan nuevas calificaciones para los usuarios y se almacenan en un archivo CSV.
Carga de datos CALIFIRNEW: Se cargan los usuarios y las nuevas calificaciones en la tabla CALIFIRNEW mediante la API.
Actualización de calificación: Se aplica la lógica del negocio original para actualizar la calificación de los usuarios en la tabla CALIFIR.
Implementación de Endpoints en la API en Go

Se han implementado varios endpoints en la API en Go para la interacción con la base de datos DB2:

Endpoint /createCalifir: Inserta nuevos registros en la tabla CALIFIR.
Endpoint /createCalifirnew: Similar a /createCalifir pero para la tabla CALIFIRNEW.
/califirnew: Consulta todos los registros de la tabla CALIFIRNEW.
/update_calfir: Ejecuta consultas SQL para actualizar registros en CALIFIR.
/shutdown: Apaga la API de forma segura.
Elección de Go para la Implementación

Se eligió Go para la implementación de la API por su eficiencia y facilidad de desarrollo. Go ofrece alto rendimiento y manejo eficiente de la concurrencia, lo que lo hace ideal para aplicaciones que requieren un alto procesamiento de datos como la gestión de una base de datos. Además, la simplicidad del código en Go facilita su mantenimiento y evolución.

Consideraciones adicionales

Este prototipo se desarrolló con fines de prueba y no representa un sistema completo. Se pueden realizar mejoras y modificaciones para adaptarlo a diferentes necesidades.
