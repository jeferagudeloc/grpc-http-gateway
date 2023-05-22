-- Insert fake data into the "order" table
INSERT INTO http_gateway_grpc.`order` (`id`, `order_number`, `creation_date`, `updation_date`, `status`)
VALUES
  ('76045a01-93e4-45c9-8670-8ebe780056a2 ', 'ORD001', '2023-05-20', '2023-05-21', 'Pending'),
  ('8433937f-18a2-4501-ad58-2abc1707aa93 ', 'ORD002', '2023-05-19', '2023-05-20', 'Completed');

-- Insert fake data into the "profile" table
INSERT INTO http_gateway_grpc.`profile` (`id`, `name`, `type`)
VALUES
  ('1ba3f566-cf9a-4814-99c9-6ceca164c1e2 ', 'Profile 1', 'Type A'),
  ('553bfd79-f007-4b9f-a9c3-c926c8588881 ', 'Profile 2', 'Type B');

-- Insert fake data into the "user" table
INSERT INTO http_gateway_grpc.`user` (`id`, `name`, `lastname`, `email`, `profile`, `status`, `profile_id`)
VALUES
  ('64a498d9-09b3-4b0d-98ed-cc796e6457c5 ', 'John', 'Doe', 'john@example.com', 'User Profile', 'Active', '1ba3f566-cf9a-4814-99c9-6ceca164c1e2 '),
  ('e5a1893f-8ff9-4082-966a-ba6824d8fe50 ', 'Jane', 'Smith', 'jane@example.com', 'Admin Profile', 'Inactive', '553bfd79-f007-4b9f-a9c3-c926c8588881 ');
