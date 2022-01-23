CREATE TABLE `restaurants` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `restaurants_categories` (
	`category_id` INT NOT NULL,
	`restaurant_id` INT NOT NULL,
	`name` VARCHAR(255),
	PRIMARY KEY (`category_id`,`restaurant_id`)
);

CREATE TABLE `categories` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `food` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `food_categories copy` (
	`category_id` INT NOT NULL,
	`food_id` INT NOT NULL,
	`name` VARCHAR(255),
	PRIMARY KEY (`category_id`,`food_id`)
);

CREATE TABLE `users` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `users_restaurants` (
	`user_id` INT NOT NULL,
	`restaurant_id` INT NOT NULL,
	`like` BOOLEAN NOT NULL DEFAULT false,
	`rating` FLOAT,
	PRIMARY KEY (`user_id`,`restaurant_id`)
);

CREATE TABLE `users_food` (
	`user_id` INT NOT NULL,
	`food_id` INT NOT NULL,
	`like` BOOLEAN NOT NULL DEFAULT false,
	`rating` FLOAT,
	PRIMARY KEY (`user_id`,`food_id`)
);

ALTER TABLE `restaurants_categories` ADD CONSTRAINT `restaurants_categories_fk0` FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`);

ALTER TABLE `restaurants_categories` ADD CONSTRAINT `restaurants_categories_fk1` FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants`(`id`);

ALTER TABLE `food_categories copy` ADD CONSTRAINT `food_categories copy_fk0` FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`);

ALTER TABLE `food_categories copy` ADD CONSTRAINT `food_categories copy_fk1` FOREIGN KEY (`food_id`) REFERENCES `food`(`id`);

ALTER TABLE `users_restaurants` ADD CONSTRAINT `users_restaurants_fk0` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

ALTER TABLE `users_restaurants` ADD CONSTRAINT `users_restaurants_fk1` FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants`(`id`);

ALTER TABLE `users_food` ADD CONSTRAINT `users_food_fk0` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

ALTER TABLE `users_food` ADD CONSTRAINT `users_food_fk1` FOREIGN KEY (`food_id`) REFERENCES `food`(`id`);









