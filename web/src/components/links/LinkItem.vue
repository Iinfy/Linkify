<template>
  <div class="recent-item">
    <div class="text-wrapper">
      <a :href="short" class="url"  target="_blank" rel="noopener">{{ short }}</a>
      <p class="original-url">{{ original }}</p>
    </div>
    <button class="copy-btn" :class="{'success': isSuccess}" @click="copyHandle()">{{buttonText}}</button>
    <button class="qr-btn" @click="showQR(props.short ?? 'not_found')">QR</button>
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

  top: 13%;
  right: 8px;

  height: 44px;
  padding: 0 20px;
  background-color: #383837;
  color: #e0e0dd;
  border: 1px solid #5c5c5a;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  width: clamp(70px, 8vw, 120px);
  transition: background-color 0.5s ease, color 0.4s ease, border-color 0.4s ease;
}

.qr-btn {
  position: absolute;
  top: 13%;
  right: 90px;        /* ← сдвигаем влево от copy-btn */
  height: 44px;
  padding: 0 14px;
  background-color: #383837;
  color: #e0e0dd;
  border: 1px solid #5c5c5a;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  font-size: 13px;
  transition: background-color 0.3s;
}

.copy-btn.success{
  background-color: #4682B4;
}

.copy-btn.success:hover {
  background-color: #4682B4;
}

.copy-btn:hover {
  background-color: #323232;
}


@media (max-width: 480px) {
  .copy-btn {
    height: 38px;
    margin-top: 3px;
    font-size: 13px;
    right: 6px;
  }


}
</style>