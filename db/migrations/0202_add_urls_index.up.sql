ALTER TABLE `urls`
	ADD UNIQUE INDEX `urls_short` (`short`);
	
ALTER TABLE `urls`
	ADD INDEX `urls_hits` (`hits`);