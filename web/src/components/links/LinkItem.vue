<template>
  <div class="recent-item">
    <div class="text-wrapper">
      <a :href="short" class="url"  target="_blank" rel="noopener">{{ short }}</a>
      <p class="original-url">{{ original }}</p>
    </div>
    <button class="icon-btn copy-btn" :class="{'success': isSuccess}" @click="copyHandle()">
    <img src="@/assets/icons/copy.svg" alt="Copy">
  </button>
    <a :href="`/stats/${props.hash}`">
    <button class="icon-btn stats-btn">
      <img src="@/assets/icons/stats.svg" alt="stats">
    </button>
    </a>
    <button class="icon-btn qr-btn" @click="showQR(props.short ?? 'not_found')">
      <img src="@/assets/icons/qr.svg" alt="QR code">
    </button>
  </div>
</template>

<script setup lang="ts">
import { useNotification } from '@/composables/useNotification'
import {showQRCode} from "@/composables/showQRCode.ts";
const { show } = useNotification()
const {showQR} = showQRCode()
import {ref} from "vue";

const isSuccess = ref(false)
const buttonText = ref<string>("Copy");

const props = defineProps({
  original: String,
  short: String,
  hash: String,
})

const copyHandle = async () => {
  if (!props.short) return

  try {
    await navigator.clipboard.writeText(props.short)
    buttonText.value = "Copied";
    isSuccess.value = true
    setTimeout(() => isSuccess.value = false, 2000)
    setTimeout(() => {
      buttonText.value = "Copy";
    }, 2000 )
    show("Copied","Link copied to clipboard", false)
  } catch (e) {
    console.error(e)
  }
}
</script>

<style scoped>

.recent-item {
  position: relative;
  width: 100%;
  height: 60px;
  background-color: #212121;
  border: 1px solid #474745;
  border-radius: 20px;
  padding: 0 16px;
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
  padding-right: 120px;
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

.icon-btn {
  position: absolute;
  top: 8px;
  width: 28px;
  height: 28px;
  padding: 0;
  background-color: #383837;
  color: #e0e0dd;
  border: 1px solid #5c5c5a;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.icon-btn:hover {
  background-color: #323232;
}

.icon-btn img {
  filter: invert(1);
  width: 14px;
  height: 14px;
  display: block;
}

.copy-btn {
  right: 8px;
}

.qr-btn {
  right: 42px;
}

.stats-btn {
  right: 76px;
}

.copy-btn.success {
  background-color: #4682B4;
  border-color: #4682B4;
}

.copy-btn.success:hover {
  background-color: #4682B4;
}
</style>