<script setup lang="ts">
import { DesktopOutlined, HomeOutlined, LogoutOutlined, MenuUnfoldOutlined } from '@ant-design/icons-vue'
import { useElementSize } from '@vueuse/core'
import auth from '@/api/auth'
import NginxControl from '@/components/NginxControl'
import Notification from '@/components/Notification'
import ProcessingStatus from '@/components/ProcessingStatus'
import { SelfCheckHeaderBanner } from '@/components/SelfCheck'
import SwitchAppearance from '@/components/SwitchAppearance'

const emit = defineEmits<{
  clickUnFold: [void]
}>()

const router = useRouter()
const { message } = useGlobalApp()
import { useUserStore } from '@/pinia'
const user = useUserStore()

function logout() {
  auth.logout().then(() => {
    message.success($gettext('Logout successful'))
  }).then(() => {
    router.push('/login')
  })
}

const headerRef = useTemplateRef('headerRef') as Ref<HTMLElement>

const userWrapperRef = useTemplateRef('userWrapperRef')
const isWorkspace = computed(() => {
  return !!window.inWorkspace
})

const { width: headerWidth } = useElementSize(headerRef)

const { width: userWrapperWidth } = useElementSize(userWrapperRef)
</script>

<template>
  <div ref="headerRef" class="header">
    <div class="tool">
      <MenuUnfoldOutlined @click="emit('clickUnFold')" />
    </div>

    <SelfCheckHeaderBanner
      :header-weight="headerWidth"
      :user-wrapper-width="userWrapperWidth"
    />

    <ASpace
      ref="userWrapperRef"
      class="user-wrapper"
      :size="24"
    >
      <SwitchAppearance />

      <div v-if="!isWorkspace" class="workspace-entry">
        <RouterLink to="/workspace">
          <ATooltip :title="$gettext('Workspace')">
            <DesktopOutlined />
          </ATooltip>
        </RouterLink>
      </div>

      <ProcessingStatus />

      <Notification v-if="user.info?.role === 'admin'" :header-ref="headerRef" />

      <NginxControl v-if="['admin', 'webdev'].includes(user.info?.role || 'admin')" />

      <a href="/">
        <HomeOutlined />
      </a>

      <a v-if="!isWorkspace" @click="logout">
        <LogoutOutlined />
      </a>
    </ASpace>
  </div>
</template>

<style lang="less" scoped>
.header {
  height: 64px;
  padding: 0 20px 0 0;
  background: transparent;
  box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.05);
  width: 100%;
  position: relative;
  a {
    color: #000000;
  }
}

.dark {
  .header {
    box-shadow: 1px 1px 0 0 #404040;

    a {
      color: #fafafa;
    }
  }
}

.tool {
  position: absolute;
  left: 20px;
  @media (min-width: 600px) {
    display: none;
  }
}

.workspace-entry {
  @media (max-width: 600px) {
    display: none;
  }
}

.user-wrapper {
  position: absolute;
  right: 28px;
}

</style>
