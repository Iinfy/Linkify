<template>
  <div class="page-wrapper">
    <div class="main__section">
      <AppLogo />
      <LinkInput @linkCreated="addLink" />
      <RecentLinks :links="links" @qr="qrUrl = $event; qrOpen = true" />
      <Notification/>

    </div>
  </div>
  <QrModal v-if="qrOpen" :url="qrUrl" @close="qrOpen = false" />
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { Link } from "@/types/link.ts";
import AppLogo from "@/components/AppLogo.vue";
import LinkInput from "@/components/links/LinkInput.vue";
import RecentLinks from "@/components/links/RecentLinks.vue";
import Notification from "@/components/NotificationPanel.vue";
import QrModal from '@/components/QrModal.vue'
const links = ref<Link[]>([]);
const qrOpen = ref(false)
const qrUrl = ref()

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