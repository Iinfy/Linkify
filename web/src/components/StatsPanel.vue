<script setup lang="ts">

</script>

<template>
  <div class="stats-card">
    <h2 class="stats-count">100 People opened this link</h2>
  </div>
</template>

<style scoped>
.stats-card {
  width: 100%;
  height: 60px;
  background-color: #212121;
  border: 1px solid #474745;
  border-radius: 20px;
  padding: 0 16px;
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

.stats-count {
  color: #e0e0dd;
  font-size: clamp(12px, 2.5vw, 17px);
  font-weight: 600;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
