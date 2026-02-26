<script setup lang="ts">
// Dimension 5: Server Hardening
// Covers SSH security, Fail2Ban, UFW firewall rules
</script>

<template>
  <div>
    <APageHeader title="มิติที่ 5: Server Hardening" />

    <ARow :gutter="[16, 16]">
      <!-- SSH Security -->
      <ACol :xs="24" :md="12">
        <ACard title="🔑 SSH Security" :bordered="false">
          <AAlert
            message="ปิดการเข้าสู่ระบบด้วย Password — บังคับใช้ SSH Key เท่านั้น"
            type="warning"
            show-icon
            style="margin-bottom: 12px;"
          />
          <ADescriptions bordered :column="1" size="small">
            <ADescriptionsItem label="PermitRootLogin">
              <ATag color="red">no</ATag>
              ปิดไม่ให้ Root ล็อกอินโดยตรง
            </ADescriptionsItem>
            <ADescriptionsItem label="PasswordAuthentication">
              <ATag color="red">no</ATag>
              บังคับใช้ SSH Key
            </ADescriptionsItem>
            <ADescriptionsItem label="Port">
              <ATag color="orange">เปลี่ยนจาก 22</ATag>
              แนะนำให้เปลี่ยนพอร์ต SSH เพื่อลด brute-force
            </ADescriptionsItem>
          </ADescriptions>
        </ACard>
      </ACol>

      <!-- Fail2Ban -->
      <ACol :xs="24" :md="12">
        <ACard title="🛡️ Intrusion Prevention (Fail2Ban)" :bordered="false">
          <AAlert
            message="ติดตั้ง Fail2Ban เพื่อบล็อก IP ที่พยายาม brute-force SSH หรือส่ง Request ผิดปกติ"
            type="info"
            show-icon
            style="margin-bottom: 12px;"
          />
          <AList size="small" :bordered="false">
            <AListItem>
              <AListItemMeta description="บล็อก IP ที่ login ผิดเกินจำนวนที่กำหนด">
                <template #title>
                  SSH Jail
                </template>
              </AListItemMeta>
            </AListItem>
            <AListItem>
              <AListItemMeta description="ป้องกัน Request ผิดปกติบน Nginx">
                <template #title>
                  Nginx Jail
                </template>
              </AListItemMeta>
            </AListItem>
          </AList>
        </ACard>
      </ACol>

      <!-- UFW Firewall -->
      <ACol :xs="24">
        <ACard title="🔥 UFW Firewall Rules" :bordered="false">
          <ATable
            :data-source="firewallRules"
            :columns="firewallColumns"
            :pagination="false"
            size="small"
            :bordered="true"
          />
        </ACard>
      </ACol>
    </ARow>
  </div>
</template>

<script lang="ts">
const firewallRules = [
  { key: '1', port: '80 (HTTP)', access: 'สาธารณะ', status: 'อนุญาต', note: 'Web traffic' },
  { key: '2', port: '443 (HTTPS)', access: 'สาธารณะ', status: 'อนุญาต', note: 'Secure web traffic' },
  { key: '3', port: '22 (SSH)', access: 'จำกัด', status: 'แนะนำเปลี่ยนพอร์ต', note: 'เข้าผ่าน VPN/IP ที่อนุญาตเท่านั้น' },
  { key: '4', port: '3306 (MySQL)', access: 'ปิด Public', status: '❌ ห้ามเปิด', note: 'เข้าผ่าน Localhost หรือ VPN เท่านั้น' },
]

const firewallColumns = [
  { title: 'Port', dataIndex: 'port', key: 'port' },
  { title: 'การเข้าถึง', dataIndex: 'access', key: 'access' },
  { title: 'สถานะ', dataIndex: 'status', key: 'status' },
  { title: 'หมายเหตุ', dataIndex: 'note', key: 'note' },
]
</script>
