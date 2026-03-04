<script setup lang="ts">
// Dimension 4: Database Management

const openPhpMyAdmin = () => {
  const host = window.location.hostname
  window.open(`http://${host}:8081`, "_blank")
}
</script>

<template>
  <div>
    <APageHeader title="มิติที่ 4: Database Management" />

    <ARow :gutter="[16, 16]">

      <!-- MySQL Server -->
      <ACol :xs="24" :md="12">
        <ACard title="🗄️ MySQL Server" :bordered="false">

          <ADivider orientation="left">Performance</ADivider>

          <ASpace direction="vertical" style="width:100%">

            <ACard size="small">
              <ATag color="blue">innodb_buffer_pool_size</ATag>
              <div style="margin-top:4px">
                แนะนำตั้งค่า <b>50–70%</b> ของ RAM เพื่อเพิ่ม performance
              </div>
            </ACard>

            <ACard size="small">
              <ATag color="blue">max_connections</ATag>
              <div style="margin-top:4px">
                ปรับตามจำนวน concurrent users (เช่น 150–300)
              </div>
            </ACard>

            <ACard size="small">
              <ATag color="orange">slow_query_log = ON</ATag>
              <div style="margin-top:4px">
                เปิดเพื่อวิเคราะห์ SQL ที่ทำงานช้า
              </div>
            </ACard>

          </ASpace>

          <ADivider orientation="left">Security</ADivider>

          <ACard size="small">
            <ASpace>
              <ATag color="red">❌ root</ATag>
              <ATag color="green">
                จำกัดสิทธิ์: SELECT, INSERT, UPDATE, DELETE
              </ATag>
            </ASpace>
            <div style="margin-top:4px">
              Web application ไม่ควรใช้ root
            </div>
          </ACard>

        </ACard>
      </ACol>

      <!-- phpMyAdmin -->
      <ACol :xs="24" :md="12">
        <ACard title="🛠 Database Tools" :bordered="false">

          <AButton type="primary" block style="margin-bottom:12px" @click="openPhpMyAdmin">
            เปิด phpMyAdmin
          </AButton>

          <AList size="small">

            <AListItem>
              <AListItemMeta description="เปลี่ยน URL ทางเข้า เช่น /db-manage-app">
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
              <AListItemMeta description="จำกัด IP ที่เข้าถึง phpMyAdmin">
                <template #title>
                  <ATag color="orange">IP Restriction</ATag>
                </template>
              </AListItemMeta>
            </AListItem>

            <AListItem>
              <AListItemMeta description="Port 3306 ห้ามเปิด Public">
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
              <p style="margin-top:8px;color:#888">
                Backup อัตโนมัติ
              </p>
            </ACol>

            <ACol :xs="24" :md="8">
              <AStatistic title="Retention" value="7 / 4 / 6" suffix="วัน/สัปดาห์/เดือน" />
            </ACol>

            <ACol :xs="24" :md="8">
              <AStatistic title="Test Restore" value="ทุกไตรมาส" />
              <p style="margin-top:8px;color:#888">
                ตรวจสอบการ restore เป็นระยะ
              </p>
            </ACol>

          </ARow>

        </ACard>
      </ACol>

    </ARow>
  </div>
</template>