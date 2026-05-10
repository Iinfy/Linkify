<template>
  <div class="page-wrapper">
    <div class="main__section">
      <AppLogo />
      <LinkInput @linkCreated="addLink" />
      <RecentLinks :links="links" />
      <Notification/>

    </div>
  </div>
  <QrModal v-if="visible"/>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { Link } from "@/types/link.ts";
import AppLogo from "@/components/AppLogo.vue";
import LinkInput from "@/components/links/LinkInput.vue";
import RecentLinks from "@/components/links/RecentLinks.vue";
import Notification from "@/components/NotificationPanel.vue";
import QrModal from '@/components/QrModal.vue'
import {showQRCode} from "@/composables/showQRCode.ts";
const links = ref<Link[]>([]);
const {visible} = showQRCode()

function addLink(link: Link) {
  links.value.unshift(link);
  if (links.value.length > 3) {
    links.value.pop()
  }
}
</script>

<style>
@import "@/assets/main.css";
</style>