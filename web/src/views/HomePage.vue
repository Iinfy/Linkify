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
import type { Link } from "@/types/link.ts";
import { ref } from "vue";
import AppLogo from "@/components/AppLogo.vue";
import LinkInput from "@/components/links/LinkInput.vue";
import RecentLinks from "@/components/links/RecentLinks.vue";
import Notification from "@/components/NotificationPanel.vue";
const links = ref<Link[]>([]);

function addLink(link: Link) {
    links.value.push(link);
}
</script>

<style>
@import "@/assets/main.css";
</style>
