# Final-SDTE
## Planning and Design
รายละเอียดเกี่ยวกับ
* ออกแบบ
* Task Management
* Acceptance Testing
* Testing

สามารถเข้าดูได้ที่ [Wiki](https://github.com/SilverSky9/final-SDTE/wiki/Feat:-ค้นหาภาค-สาขา-คณะที่เปิดรับสมัคร-สำหรับหลักสูตร-ปริญญาตรี)


## Deployment
### ขั้นตอนการ Build
Run คำสั่งต่อไปนี้เพื่อ ฺBuild docker image สำหรับ frontend และ backend <br>
``` docker compose -f docker-compose-build.yml build ```

### ขั้นตอนการ Deploy
Run คำสั่งต่อไปนี้เพื่อ Down docker container ก่อนที่จะ Up ขึ่นมาใหม่ สำหรับ frontend และ backend <br>
``` docker compose -f docker-compose-deploy.yml down ```

Run คำสั่งต่อไปนี้เพื่อ Up docker container ขึ้นมา สำหรับ frontend และ backend <br>
``` docker compose -f docker-compose-deploy.yml up -d ```

