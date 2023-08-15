CREATE TABLE `ingress` (
    `id` varchar(36) PRIMARY KEY,
    `tenant_id` varchar(36) NOT NULL,
    `display_name` varchar(255) NOT NULL
);