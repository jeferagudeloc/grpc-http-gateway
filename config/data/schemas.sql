-- Create a schema
CREATE SCHEMA IF NOT EXISTS http_gateway_grpc;

-- Create the "order" table in the schema
CREATE TABLE http_gateway_grpc.`order` (
  `id` VARCHAR(50) PRIMARY KEY,
  `order_number` VARCHAR(255) NOT NULL,
  `creation_date` DATE,
  `updation_date` DATE,
  `status` VARCHAR(50)
);

-- Create the "profile" table in the schema
CREATE TABLE http_gateway_grpc.`profile` (
  `id` VARCHAR(50) PRIMARY KEY,
  `name` VARCHAR(50) NOT NULL,
  `type` VARCHAR(50)
);

-- Create the "user" table in the schema
CREATE TABLE http_gateway_grpc.`user` (
  `id` VARCHAR(50) PRIMARY KEY,
  `name` VARCHAR(50) NOT NULL,
  `lastname` VARCHAR(50) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `profile` VARCHAR(50),
  `status` VARCHAR(50),
  `profile_id` VARCHAR(50),
  FOREIGN KEY (`profile_id`) REFERENCES http_gateway_grpc.`profile` (`id`)
);
