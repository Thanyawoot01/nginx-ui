<script setup lang="ts">
import useSystemSettingsStore from '../store'

const systemSettingsStore = useSystemSettingsStore()
const { data, errors } = storeToRefs(systemSettingsStore)
</script>

<template>
  <AForm layout="vertical">
    <AFormItem
      :label="$gettext('Granted Access Paths')"
      :validate-status="errors?.backup?.granted_access_path ? 'error' : ''"
      :help="errors?.backup?.granted_access_path
        ? errors?.backup?.granted_access_path
        : $gettext('A list of directory paths that are allowed for backup operations.')"
    >
      <ASelect
        v-model:value="data.backup.granted_access_path"
        mode="tags"
        :placeholder="$gettext('Press enter to add a path')"
        :token-separators="[',']"
      />
    </AFormItem>
  </AForm>
</template>

<style lang="less" scoped>
</style>
