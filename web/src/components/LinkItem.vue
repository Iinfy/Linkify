<template>
  <div class="recent-item">
    <div class="text-wrapper">
      <a :href="short" class="url" target="_blank" rel="noopener">{{ short }}</a>
      <p class="original-url">{{ original }}</p>
    </div>
    <button class="copy-btn">Copy</button>
  </div>
</template>

<script setup lang="ts">
defineProps({
  original: String,
  short: String,
})
</script>

<style>

.recent-item {
  position: relative;
  width: 100%;
  height: 60px;
  background-color: #212121;
  border: 1px solid #474745;
  border-radius: 20px;
  padding: 0 80px 0 16px;
  display: flex;
  align-items: center;
  overflow: hidden;
  flex-shrink: 0;
}

.text-wrapper {
  display: flex;
  flex-direction: column;
  min-width: 0;
  flex: 1;
  padding-right: 8px;
}

.url {
  color: #e0e0dd;
  font-size: clamp(12px, 2.5vw, 17px);
  text-decoration: none;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  display: block;
  margin-bottom: 2px;
}

.original-url {
  color: #5F5F5F;
  font-size: clamp(10px, 2vw, 13px);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin: 0;
  display: block;
}

.copy-btn {
  position: absolute;
  top: 50%;
  right: 8px;
  transform: translateY(-50%);
  height: 36px;
  padding: 0 14px;
  background-color: #383837;
  color: #e0e0dd;
  border: 1px solid #5c5c5a;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: clamp(11px, 2vw, 14px);
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.4s ease;
}
</style>