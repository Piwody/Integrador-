# Sistema de Gestión de Imágenes

Este proyecto representa una solución integral de software e infraestructura para el almacenamiento, procesamiento y administración segura de recursos multimedia en un entorno corporativo. El sistema combina el desarrollo de servicios distribuidos bajo el modelo **cliente-servidor** con el diseño de una infraestructura de red física y lógica segmentada mediante **VLANs**.

---

## 1. ¿Qué es este Sistema?

Es un ecosistema de software distribuido de alta eficiencia compuesto por dos microservicios principales que cooperan de forma aislada:
1. **Módulo de Autenticación y Control de Acceso:** Desarrollado en Python con FastAPI, encargado del gobierno de identidades digitales.
2. **Módulo Core de Procesamiento Multimedia:** Desarrollado en Go, diseñado bajo principios de Arquitectura Hexagonal para la manipulación orientada al rendimiento de archivos de imágenes.

A nivel de infraestructura, el sistema no corre en una red plana tradicional; está dentro de una red organizacional compartimentada que utiliza routing e switching avanzado para optimizar el tráfico de datos.

---

## 2. ¿Para qué funciona? (Propósito y Utilidad)

En el ámbito empresarial, las imágenes (credenciales, auditorías, planos, capturas de datos) constituyen activos de información críticos. Este sistema funciona para:
* **Garantizar la Confidencialidad:** Evita que usuarios no autorizados o atacantes en la red intercepten credenciales o accedan a imágenes privadas.
* **Optimizar la Disponibilidad:** Utiliza Go en el Core de archivos para responder con latencias mínimas y alta concurrencia al procesar cargas pesadas de imágenes.
* **Establecer Políticas de Zero-Trust :** Un usuario o dispositivo en la red no puede interactuar con el almacenamiento multimedia si no ha sido explícitamente validado y firmado por el servicio de identidad.

---

## 3. ¿Cómo funciona el Sistema? (Mecánica de Operación)

El funcionamiento del sistema se divide en dos fases críticas que cruzan las fronteras del software y de la red:

### A. El Ciclo de Autenticación y Cifrado
1. **Registro:** El usuario envía sus datos desde la **VLAN 10 (Clientes)**. La API de FastAPI en la **VLAN 20** intercepta la petición. Antes de tocar el disco, la contraseña pasa por un motor criptográfico basado en el estándar AES, el cual transforma el texto plano en bloques cifrados no legibles, almacenándolos finalmente en la base de datos relacional **SQLite** (gestionada mediante **SQLModel, Dbeaver**).
2. **Login:** El cliente envía sus credenciales. El backend recupera el bloque cifrado de la base de datos, realiza la validación correspondiente y, si los datos coinciden, emite un token de sesión seguro que autoriza al cliente a operar.

### B. El Ciclo de Gestión de Imágenes e Inter-VLAN Routing
1. **Petición de Carga:** El cliente (VLAN 10), ya autenticado, realiza una petición de tipo `MultipartForm` al servidor de Go, el cual reside de manera aislada en la **VLAN 30**.
2. **Validación perimetral y de aplicación:** Los switches y routers de la infraestructura controlan que el tráfico viaje exclusivamente por los puertos autorizados. El servidor de Go recibe el archivo físico, valida sus metadatos (tamaño, tipo de formato de imagen), mapea el ID del propietario y guarda el archivo en el sistema de almacenamiento persistente.

---

## 4. Conocimientos y Pilares Científicos Aplicados

Para la construcción de este proyecto, el equipo integra disciplinas avanzadas de la ingeniería de software y la telemática:

###  Ingeniería de Software y Arquitectura Limpia
* **Arquitectura Hexagonal (Ports & Adapters):** Aplicada en el servicio de Go. Permite desacoplar las reglas del negocio (dominio e imágenes) de los factores externos (servidores HTTP, drivers de almacenamiento o bases de datos). Esto garantiza la mantenibilidad a largo plazo.
* **Modelo Cliente-Servidor Distribuido:** Separación de responsabilidades donde las capas de presentación, lógica y datos se comunican mediante protocolos de aplicación estandarizados (HTTP/REST).

###  Criptografía y Seguridad de la Información
* **Cifrado Simétrico por Bloques (AES):** Aplicación de algoritmos matemáticos seguros para proteger la persistencia de datos sensibles, cumpliendo con estándares internacionales de seguridad industrial.
* **Políticas de Privacidad del Diseño (Privacy by Design):** Estructuración de flujos lógicos donde los datos de identidad digital nunca se exponen ni viajan en texto plano a través de la red.

###  Redes de Datos y Conmutación Corporativa
* **Segmentación (VLANs):** Diseño lógico de redes para aislar entornos lógicos de servidores corporativos de los entornos de usuarios finales, disminuyendo la superficie de ataque del sistema.
* **Enrutamiento Inter-VLAN y ACLs:** Configuración de directrices de enrutamiento de control de acceso para regular estrictamente qué segmentos de la red tienen permitido comunicarse entre sí.

###  Impacto Social, Ético y Cultural
* **Inclusión Tecnológica:** El uso de tecnologías eficientes y de bajo consumo de recursos (Go + SQLite) permite el despliegue del sistema en hardware heredado o de bajo costo, promoviendo la accesibilidad tecnológica.
* **Ética y Gobernanza de Datos:** El sistema implementa un esquema de responsabilidad digital donde cada recurso multimedia está estrictamente ligado a una identidad, previniendo el uso indebido o malintencionado de la infraestructura de almacenamiento.

---

## 5. Stack Tecnológico Utilizado

* **Lenguajes de Programación:** Go (Golang), Python.
* **Frameworks & ORMs:** FastAPI, SQLModel.
* **Motor de Base de Datos:** SQLite.
* **Criptografía:** Librerías nativas de seguridad (Estándar AES).
* **Simulación de Infraestructura:** Cisco Packet Tracer / Entornos GNS3.
* **Control de Versiones:** Git & GitHub.

  
