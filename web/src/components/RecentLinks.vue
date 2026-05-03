<script setup>
import LinkItem from './LinkItem.vue'
import { ref } from 'vue'

const links = ref([
  { id: 1, original: 'https://google.com/', short: 'http://sh.rt/abc' },
  { id: 2, original: 'https://youtube.com/', short: 'http://sh.rt/xyz' },
  { id: 3, original: 'https://youtube.com/', short: 'http://sh.rt/xyz' },

])
</script>

<template>
  <div class="recent-links">
    <h3 class="recent-titles">Recent Links</h3>
    <div class="recent-links-list">
      <LinkItem
          v-for="link in links"
          :key="link.id"
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
</style>