<script setup lang="ts">
import LinkItem from "@/components/links/LinkItem.vue";
import { type PropType, ref } from "vue";
import type { Link } from "@/types/link.ts";

defineProps({ links: Array as PropType<Link[]> });
</script>

<template>
  <div class="recent-links">
    <h3 class="recent-titles">Recent Links</h3>
    <div class="recent-links-list">
      <LinkItem
          v-for="link in links"
          :original="link.original"
          :short="link.short"
      />
    </div>
  </div>
</template>

<style scoped>
.recent-links {
  width: 100%;
  margin-top: 24px;
}

.recent-titles {
  color: white;
  font-size: 15px;
  font-weight: 600;
  margin: 0 0 12px;
}

.recent-links-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
  max-height: 220px;
  overflow: hidden;
  width: 100%;
}

@media (max-width: 480px) {
  .recent-links {
    margin-top: 20px;
  }

  .recent-links-list {
    gap: 10px;
    max-height: 180px;
  }
}
</style>