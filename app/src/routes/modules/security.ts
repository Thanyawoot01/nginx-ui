import type { RouteRecordRaw } from 'vue-router'
import { LockOutlined } from '@ant-design/icons-vue'

export const securityRoutes: RouteRecordRaw[] = [
  {
    path: 'security',
    name: 'Security Hardening',
    component: () => import('@/views/security/SecurityHardening.vue'),
    meta: {
      name: () => $gettext('Server Hardening'),
      icon: LockOutlined,
      roles: ['admin'],
    },
  },
]
