<template>
    <div class="page-wrapper">
        <div class="main__section">
            <AppLogo />
            <LinkInput @linkCreated="addLink" />
            <RecentLinks :links="links" />
            <Notification/>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { Link } from "@/types/link.ts";
import AppLogo from "@/components/AppLogo.vue";
import LinkInput from "@/components/links/LinkInput.vue";
import RecentLinks from "@/components/links/RecentLinks.vue";
import Notification from "@/components/NotificationPanel.vue";
const links = ref<Link[]>([]);

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
