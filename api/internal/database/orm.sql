CREATE TABLE `ingress`
(
    `id`           varchar(36) PRIMARY KEY,
    `display_name` varchar(255) NOT NULL,
    `config`       json         not null
);

CREATE TABLE IF NOT EXISTS `ingress_log`
(
    `id`           varchar(36) PRIMARY KEY,
    `ingress_id`   varchar(36)  NOT NULL,
    `status`       varchar(255) NOT NULL,
    `created_at`   datetime     NOT NULL,
    `request_hash` varchar(255) NOT NULL,
    `request`      blob         NOT NULL,
    `response`     blob         NOT NULL
);

CREATE TABLE `egress`
(
    `id`           varchar(36) PRIMARY KEY,
    `display_name` varchar(255) NOT NULL,
    `config`       json         not null
);
INSERT INTO `egress` (`id`, `display_name`, `config`)
VALUES ('00000000-0000-0000-0000-000000000001', 'Ingress list', '{"handler": "list"}'),
       ('00000000-0000-0000-0000-000000000002', 'Ingress item', '{"handler": "item"}');

CREATE TABLE `resource_group`
(
    `id`           varchar(36) PRIMARY KEY,
    `display_name` varchar(255) NOT NULL,
    `config`       json         not null
);
INSERT INTO `resource_group` (`id`, `display_name`, `config`)
VALUES ('00000000-0000-0000-0000-000000000000', 'djeurnie core 1.0', '{}');

CREATE TABLE IF NOT EXISTS `api_routes`
(
    `id`                varchar(36) PRIMARY KEY,
    `resource_group_id` varchar(36)  NOT NULL,
    `path`              varchar(255) NOT NULL,
    `method`            varchar(255) NOT NULL,
    `config`            json         NOT NULL
);
INSERT INTO `api_routes` (id, resource_group_id, path, method, config)
VALUES ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000000', '/ingress', 'GET', '{"Target": "egress", "TargetId": "00000000-0000-0000-0000-000000000001"}'),
       ('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000000', '/ingress/:id', 'GET', '{"Target": "egress", "TargetId": "00000000-0000-0000-0000-000000000002"}');

