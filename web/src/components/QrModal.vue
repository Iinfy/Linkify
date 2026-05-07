<template>
  <div class="qr-container">
    <div class="qr-modal">
      <button class="close-btn" @click="close()">✕</button>
      <div class="qr-wrapper">
        <div ref="qrContainer"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue'
import QRCodeStyling from 'qr-code-styling'
import {showQRCode} from "@/composables/showQRCode.ts";
const { text, close } = showQRCode()


const qrContainer = ref<HTMLDivElement>()


onMounted(() => {
  if (!qrContainer.value) return
  console.log("mounted")

  const qr = new QRCodeStyling({
    width: 300,
    height: 300,
    type: "svg",
    data: text.value,
    dotsOptions: {
      color: "white",
      type: "rounded"
    },

    backgroundOptions: {
      color: "black",
    },
    imageOptions: {
      crossOrigin: "anonymous",
      margin: 20,
    }
  })

  qr.append(qrContainer.value)
})
</script>

<style scoped>
.qr-container {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.qr-modal {
  background-color: #212121;
  padding: 24px;
  border-radius: 20px;
  overflow: hidden;
  position: relative;
}

.qr-wrapper {
  border-radius: 24px;
  overflow: hidden;
  background: black;
  padding: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
  margin-top: 20px;
}

.close-btn {
  position: absolute;
  top: 12px;
  right: 12px;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: none;
  border: none;
  color: #aaaaaa;
  font-size: 22px;
  font-weight: 300;
  cursor: pointer;
  z-index: 10;
  border-radius: 50%;
  transition: all 0.2s;
}
</style>