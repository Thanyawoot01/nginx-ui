# LAMP/LEMP Stack Management System

ระบบบริหารจัดการ LAMP/LEMP Stack แบบครบวงจร มุ่งเน้นความปลอดภัย ประสิทธิภาพ และการจัดการสิทธิ์ตามหลัก Least Privilege

---

## 5 มิติการจัดการ

| มิติ | หัวข้อ | เครื่องมือ |
|------|--------|-----------|
| 1 | **User & Access Control** | Linux Groups, sudoers, SGID |
| 2 | **OS & Infrastructure** | htop, UFW, unattended-upgrades |
| 3 | **Web Server Management** | NGINX Server Blocks, Logs, Terminal |
| 4 | **Database Management** | MySQL tuning, phpMyAdmin security |
| 5 | **Security & Backup** | SSH Key, Fail2Ban, mysqldump |

---

## บทบาทและสิทธิ์ (Roles & Permissions)

| Role | หน้าที่ | สิทธิ์ OS |
|------|--------|-----------|
| **System Admin** | ดูแล OS, Network, Security, Services | Full Root (`sudo`) |
| **Software Dev** | Deploy โค้ด, ดู Log, restart Web Server | `/var/www/html`, read logs |
| **DB Admin (DBA)** | ปรับจูน MySQL, Backup/Restore | MySQL service/config |

---

## การใช้งาน

### ติดตั้งด้วย Docker

```bash
docker run -dit \
  --name=nginx-ui \
  --restart=always \
  -e TZ=Asia/Bangkok \
  -v /mnt/user/appdata/nginx:/etc/nginx \
  -v /mnt/user/appdata/nginx-ui:/etc/nginx-ui \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -p 8080:80 -p 8443:443 \
  uozi/nginx-ui:latest
```

เข้าใช้งาน: `http://<server_ip>:8080/install`

### Build จาก Source

```bash
# Frontend
cd app && pnpm install && pnpm build

# Backend
go generate
go build -tags=jsoniter -o nginx-ui -v main.go
./nginx-ui -config app.ini
```

---

## หลักการความปลอดภัย

- **Least Privilege** — ให้สิทธิ์เท่าที่จำเป็น
- **HTTPS บังคับ** — SSL/TLS ทุก endpoint
- **SSH Key Only** — ปิด PasswordAuthentication
- **Firewall** — เปิดเฉพาะ port 80, 443 (public) และ 22 (จำกัด IP)
- **Backup อัตโนมัติ** — รายวัน + test restore ทุกไตรมาส

## License

GNU Affero General Public License v3.0 — ดูรายละเอียดในไฟล์ [LICENSE](LICENSE)
