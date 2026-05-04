<template>
    <div class="page-wrapper">
        <div class="main__section">
            <AppLogo />
            <LinkInput @linkCreated="addLink" />
            <RecentLinks :links="links" />
            <Notification @copied="addLink" :data="data" v-if="data" />
        </div>
    </div>
</template>

<script setup lang="ts">
import type { Link } from "@/types/link.ts";
import { ref } from "vue";
import AppLogo from "@/components/AppLogo.vue";
import LinkInput from "@/components/links/LinkInput.vue";
import RecentLinks from "@/components/links/RecentLinks.vue";
import Notification from "@/components/NotificationPanel.vue";
import type {NotificationData} from "@/types/notificationData.ts";
const links = ref<Link[]>([]);
const data = ref<NotificationData>();

function addLink(link: Link) {
    links.value.push(link);
    if (data.value) {
        data.value.title = "copied"
    }

  if (data.value) {
    data.value.subtitle = "Short link copied to clipboard"
  }

}

function triggerNotification(title: string, subtitle: string) {

}
</script>

<style>
@import "@/assets/main.css";
</style>
