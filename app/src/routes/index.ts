import type { RouteRecordRaw } from 'vue-router'
import { createRouter, createWebHashHistory } from 'vue-router'
import { useNProgress } from '@/lib/nprogress/nprogress'
import { useUserStore } from '@/pinia'
import { authRoutes } from './modules/auth'

// Dimension 1: User & Access Control
import { userRoutes } from './modules/user'
// Dimension 2: OS & Infra (Dashboard)
import { dashboardRoutes } from './modules/dashboard'
// Dimension 3: Web Server Management
import { configRoutes } from './modules/config'
import { nginxLogRoutes } from './modules/nginx_log'
import { sitesRoutes } from './modules/sites'
import { terminalRoutes } from './modules/terminal'
// Dimension 4: Database Management
import { databaseRoutes } from './modules/database'
// Dimension 5: Security & Backup
import { backupRoutes } from './modules/backup'
import { securityRoutes } from './modules/security'
import { systemRoutes } from './modules/system'

import { errorRoutes } from './modules/error'
import 'nprogress/nprogress.css'

// Combine child routes for the main layout (5-dimension LAMP/LEMP structure)
const mainLayoutChildren: RouteRecordRaw[] = [
  // มิติที่ 2: OS & Infra
  ...dashboardRoutes,
  // มิติที่ 1: User & Access Control
  ...userRoutes,
  // มิติที่ 3: Web Server Management
  ...sitesRoutes,
  ...configRoutes,
  ...nginxLogRoutes,
  ...terminalRoutes,
  // มิติที่ 4: Database Management
  ...databaseRoutes,
  // มิติที่ 5: Security & Backup
  ...backupRoutes,
  ...systemRoutes,
  ...securityRoutes,
]

// Main routes configuration
export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/layouts/BaseLayout.vue'),
    redirect: '/dashboard',
    meta: {
      name: () => $gettext('Home'),
    },
    children: mainLayoutChildren,
  },
  {
    path: '/workspace',
    name: 'Workspace',
    component: () => import('@/views/workspace/WorkSpace.vue'),
    meta: {
      name: () => $gettext('Workspace'),
    },
  },
  ...authRoutes,
  ...errorRoutes,
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

const nprogress = useNProgress()

router.beforeEach((to, _, next) => {
  document.title = `${to?.meta.name?.() ?? ''} | Nginx UI`

  nprogress.start()

  const user = useUserStore()

  if (to.meta.noAuth || user.isLogin)
    next()
  else
    next({ path: '/login', query: { next: to.fullPath } })
})

router.afterEach(() => {
  nprogress.done()
})

export default router
