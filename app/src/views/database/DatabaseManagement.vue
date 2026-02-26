<script setup lang="ts">
// Dimension 4: Database Management
// Covers MySQL security, performance tuning, and phpMyAdmin hardening
</script>

<template>
  <div>
    <APageHeader title="มิติที่ 4: Database Management" />

    <ARow :gutter="[16, 16]">
      <!-- MySQL Server -->
      <ACol :xs="24" :md="12">
        <ACard title="🗄️ MySQL Server" :bordered="false">
          <ADescriptions bordered :column="1" size="small">
            <ADescriptionsItem label="Buffer Pool Size">
              <ATag color="blue">innodb_buffer_pool_size</ATag>
              ตั้งค่า 50–70% ของ RAM ทั้งหมด
            </ADescriptionsItem>
            <ADescriptionsItem label="Max Connections">
              <ATag color="blue">max_connections</ATag>
              ปรับตามจำนวน concurrent users
            </ADescriptionsItem>
            <ADescriptionsItem label="Slow Query Log">
              <ATag color="orange">slow_query_log = ON</ATag>
              เปิดให้ Dev วิเคราะห์และปรับปรุง SQL
            </ADescriptionsItem>
            <ADescriptionsItem label="Web App User">
              <ATag color="red">❌ root</ATag>
              <ATag color="green">✅ จำกัดสิทธิ์: SELECT, INSERT, UPDATE, DELETE</ATag>
            </ADescriptionsItem>
          </ADescriptions>
        </ACard>
      </ACol>

      <!-- phpMyAdmin Security -->
      <ACol :xs="24" :md="12">
        <ACard title="🔐 phpMyAdmin Security" :bordered="false">
          <AList size="small" :bordered="false">
            <AListItem>
              <AListItemMeta
                description="เปลี่ยน URL ทางเข้าให้คาดเดายาก เช่น /db-manage-app"
              >
                <template #title>
                  <ATag color="purple">URL Obfuscation</ATag>
                </template>
              </AListItemMeta>
            </AListItem>
            <AListItem>
              <AListItemMeta description="ปิดการล็อกอินด้วย root (AllowRoot = false)">
                <template #title>
                  <ATag color="red">Access Control</ATag>
                </template>
              </AListItemMeta>
            </AListItem>
            <AListItem>
              <AListItemMeta description="จำกัด IP ที่เข้าถึงได้ ใช้ .htpasswd ครอบ หรือเปิด 2FA">
                <template #title>
                  <ATag color="orange">IP Restriction / 2FA</ATag>
                </template>
              </AListItemMeta>
            </AListItem>
            <AListItem>
              <AListItemMeta description="Port 3306 ห้ามเปิด Public — เข้าผ่าน VPN หรือ Localhost เท่านั้น">
                <template #title>
                  <ATag color="red">Port 3306</ATag>
                </template>
              </AListItemMeta>
            </AListItem>
          </AList>
        </ACard>
      </ACol>

      <!-- Backup -->
      <ACol :xs="24">
        <ACard title="💾 Database Backup & Replication" :bordered="false">
          <ARow :gutter="16">
            <ACol :xs="24" :md="8">
              <AStatistic title="Automation" value="mysqldump" suffix="ทุกวัน" />
              <p style="margin-top: 8px; color: #888;">ส่งขึ้น Cloud Storage อัตโนมัติ</p>
            </ACol>
            <ACol :xs="24" :md="8">
              <AStatistic title="Retention" value="7 / 4 / 6" suffix="วัน/สัปดาห์/เดือน" />
            </ACol>
            <ACol :xs="24" :md="8">
              <AStatistic title="Test Restore" value="ทุกไตรมาส" />
              <p style="margin-top: 8px; color: #888;">MySQL Replication (Master-Slave) สำหรับระบบ Critical</p>
            </ACol>
          </ARow>
        </ACard>
      </ACol>
    </ARow>
  </div>
</template>
