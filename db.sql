CREATE TABLE faculty (
    faculty_id int,
    faculty_name varchar(255)
);

INSERT INTO faculty (faculty_id, faculty_name)
VALUES (1, 'วิศวกรรมศาสตร์');
INSERT INTO faculty (faculty_id, faculty_name)
VALUES (2, 'บริหาร');
INSERT INTO faculty (faculty_id, faculty_name)
VALUES (3, 'ทัตแพทยศาสตร์');
INSERT INTO faculty (faculty_id, faculty_name)
VALUES (4, 'อุสาหกรรมการบิน');
INSERT INTO faculty (faculty_id, faculty_name)
VALUES (5, 'สถาปัตยกรรมศาสตร์');
INSERT INTO faculty (faculty_id, faculty_name)
VALUES (6, 'นวัตกรรมขั้นสูง');


CREATE TABLE round_ava (
    round_id int,
    round_name varchar(255)
);
INSERT INTO round_ava (round_id, round_name)
VALUES (1, '(1/2565) รอบ Dircet Admisstion 3 (international Program)');
INSERT INTO round_ava (round_id, round_name)
VALUES (2, '(1/2565) รอบ Dircet Admisstion 2 (international student Program)');
INSERT INTO round_ava (round_id, round_name)
VALUES (3, '(1/2565) รอบ Dircet Admisstion 2 (international Program)');
INSERT INTO round_ava (round_id, round_name)
VALUES (4, '(1/2565) รอบ 2 โครงการอาชีวะพรีเมี่ยม');


CREATE TABLE major_ava (
    major_id int,
    major_name varchar(255),
    major varchar(255),
    major_des varchar(255),
    open_date varchar(255)
);
INSERT INTO major_ava (major_id, major_name, major, major_des, open_date)
VALUES (1, 'หลักสูตร : Chemical Engineering (International Program)', 'คณะ วิศวกรรมศาสตร์', 'หลักสูตร : Chemical Engineering (International Program) / Chemical Engineering (International Program) ', 'วันที่เปิดรับสมัคร : 3 พฤษภาคม 2022 - 19 พฤษภาคม 2022');
INSERT INTO major_ava (major_id, major_name, major, major_des, open_date)
VALUES (2, 'หลักสูตร : Electrical Engineering (International Program) | แขนง : ไฟฟ้ากำลัง (นานาชาติ)', 'คณะ วิศวกรรมศาสตร์', 'หลักสูตร : Chemical Engineering (International Program) / Chemical Engineering (International Program) ', 'วันที่เปิดรับสมัคร : 3 พฤษภาคม 2022 - 19 พฤษภาคม 2022');
INSERT INTO major_ava (major_id, major_name, major, major_des, open_date)
VALUES (3, 'หลักสูตร : การเป็นผู้ประกอบการระดับโลก (หลักสูตรนานาชาติ)', 'คณะ บริหารธุรกิจ', 'หลักสูตร : การเป็นผู้ประกอบการระดับโลก (หลักสูตรนานาชาติ) / Business Administration Program in Global Entrepreneurship (International Prog', 'วันที่เปิดรับสมัคร : 5 เมษายน 2022 - 23 พฤษภาคม 2022');