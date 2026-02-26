import type { RouteRecordRaw } from 'vue-router'
import { DatabaseOutlined } from '@ant-design/icons-vue'

export const databaseRoutes: RouteRecordRaw[] = [
  {
    path: 'database',
    name: 'Database Management',
    component: () => import('@/views/database/DatabaseManagement.vue'),
    meta: {
      name: () => $gettext('Database Management'),
      icon: DatabaseOutlined,
    },
  },
]
