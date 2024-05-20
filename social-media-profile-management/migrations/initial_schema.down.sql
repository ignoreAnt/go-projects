-- Drop triggers
DELIMITER //

DROP TRIGGER IF EXISTS before_update_user;
//
DROP TRIGGER IF EXISTS before_update_user_profile;
//
DROP TRIGGER IF EXISTS before_update_picture;
//
DROP TRIGGER IF EXISTS before_update_work_detail;
//
DROP TRIGGER IF EXISTS before_update_education_detail;
//
DROP TRIGGER IF EXISTS before_update_life_event;
//

DELIMITER ;

-- Drop tables in reverse order to avoid foreign key constraints issues
DROP TABLE IF EXISTS life_events;
DROP TABLE IF EXISTS education_details;
DROP TABLE IF EXISTS work_details;
DROP TABLE IF EXISTS pictures;
DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS users;
