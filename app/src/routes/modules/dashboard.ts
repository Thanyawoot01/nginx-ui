import type { RouteRecordRaw } from 'vue-router'
import { DashboardOutlined } from '@ant-design/icons-vue'

export const dashboardRoutes: RouteRecordRaw[] = [
  {
    path: 'dashboard',
    redirect: '/dashboard/server',
    name: 'Dashboard',
    meta: {
      name: () => $gettext('OS & Infra Monitor'),
      icon: DashboardOutlined,
      hideChildren: true,
      roles: ['admin'],
    },
    children: [
      {
        path: 'server',
        component: () => import('@/views/dashboard/ServerDashBoard.vue'),
        name: 'Server',
        meta: {
          name: () => $gettext('OS & Infra Monitor'),
        },
      },
    ],
  },
]
