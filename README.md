# Final-SDTE
## Planning and Design
รายละเอียดเกี่ยวกับ
* ออกแบบ
* Task Management
* Acceptance Testing
* Testing

สามารถเข้าดูได้ที่ [Wiki](https://github.com/SilverSky9/final-SDTE/wiki/Feat:-ค้นหาภาค-สาขา-คณะที่เปิดรับสมัคร-สำหรับหลักสูตร-ปริญญาตรี)


## Deployment
สามารถเข้าดู Jenkins Pipeline ได้ที่ [128.199.128.171](http://128.199.128.171:8085/) <br>
```
User : admin
Pass : 1234
```
### ขั้นตอนการ Build
Run คำสั่งต่อไปนี้เพื่อ Build docker image สำหรับ frontend และ backend
``` 
docker compose -f docker-compose-build.yml build
```

### ขั้นตอนการ Deploy
Run คำสั่งต่อไปนี้เพื่อ Down docker container ก่อนที่จะ Up ขึ้นมาใหม่ สำหรับ frontend และ backend
``` 
docker compose -f docker-compose-deploy.yml down 
```

Run คำสั่งต่อไปนี้เพื่อ Up docker container ขึ้นมา สำหรับ frontend และ backend
``` 
docker compose -f docker-compose-deploy.yml up -d
```

## API
List API ที่สามารถเรียกใช้งานได้เมื่อทำการ Deploy service เรียบร้อยหมดแล้ว

### Check health service
request
```
    http://localhost:3050/health
```
response
```
    "Go fiber is good!"
```

### Get all faculty
request
```
    http://localhost:3050/faculty/all
```
response
```json
    [
    {
        "faculty_id": 1,
        "faculty_name": "วิศวกรรมศาสตร์"
    },
    {
        "faculty_id": 2,
        "faculty_name": "บริหาร"
    },
    {
        "faculty_id": 3,
        "faculty_name": "ทัตแพทยศาสตร์"
    },
    {
        "faculty_id": 4,
        "faculty_name": "อุสาหกรรมการบิน"
    },
    {
        "faculty_id": 5,
        "faculty_name": "สถาปัตยกรรมศาสตร์"
    },
    {
        "faculty_id": 6,
        "faculty_name": "นวัตกรรมขั้นสูง"
    }
]
```

### Get all round open
request
```
    http://localhost:3050/round/all
```
response
```json
    [
    {
        "round_id": 1,
        "round_name": "(1/2565) รอบ Dircet Admisstion 3 (international Program)"
    },
    {
        "round_id": 2,
        "round_name": "(1/2565) รอบ Dircet Admisstion 2 (international student Program)"
    },
    {
        "round_id": 3,
        "round_name": "(1/2565) รอบ Dircet Admisstion 2 (international Program)"
    },
    {
        "round_id": 4,
        "round_name": "(1/2565) รอบ 2 โครงการอาชีวะพรีเมี่ยม"
    }
]
```

### Get all data to display
request
```
    http://localhost:3050/data/all
```
response
```json
    [
    {
        "major_id": 1,
        "major_name": "หลักสูตร : Chemical Engineering (International Program)",
        "major": "คณะ วิศวกรรมศาสตร์",
        "major_des": "หลักสูตร : Chemical Engineering (International Program) / Chemical Engineering (International Program) ",
        "open_date": "วันที่เปิดรับสมัคร : 3 พฤษภาคม 2022 - 19 พฤษภาคม 2022"
    },
    {
        "major_id": 2,
        "major_name": "หลักสูตร : Electrical Engineering (International Program) | แขนง : ไฟฟ้ากำลัง (นานาชาติ)",
        "major": "คณะ วิศวกรรมศาสตร์",
        "major_des": "หลักสูตร : Chemical Engineering (International Program) / Chemical Engineering (International Program) ",
        "open_date": "วันที่เปิดรับสมัคร : 3 พฤษภาคม 2022 - 19 พฤษภาคม 2022"
    },
    {
        "major_id": 3,
        "major_name": "หลักสูตร : การเป็นผู้ประกอบการระดับโลก (หลักสูตรนานาชาติ)",
        "major": "คณะ บริหารธุรกิจ",
        "major_des": "หลักสูตร : การเป็นผู้ประกอบการระดับโลก (หลักสูตรนานาชาติ) / Business Administration Program in Global Entrepreneurship (International Prog",
        "open_date": "วันที่เปิดรับสมัคร : 5 เมษายน 2022 - 23 พฤษภาคม 2022"
    }
]
```

