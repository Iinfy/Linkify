

<template>
  <div class="qr-container">
    <div class="qr-modal">
      <button class="close-btn" @click="close()">✕</button>
      <div class="qr-wrapper">
        <div ref="qrContainer"></div>
      </div>
      <button class="save-qr" :class="{'success': isSuccess}" @click="downloadQr()">
        Save
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import QRCodeStyling from 'qr-code-styling'
import {ref, onMounted} from 'vue'
import {showQRCode} from "@/composables/showQRCode.ts";
import {useNotification} from "@/composables/useNotification.ts";
const { show } = useNotification()
const { text, close } = showQRCode()
const qr = ref<QRCodeStyling | null>(null)
const qrContainer = ref<HTMLDivElement>()
const isSuccess = ref(false)


onMounted(() => {
  if (!qrContainer.value) return

  qr.value = new QRCodeStyling({
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

  qr.value.append(qrContainer.value)
})

const downloadQr = async () => {
  if (!qr.value) return
  try {
    await qr.value.download({
      name: "qr-code",
      extension: "png",
  })
    isSuccess.value = true
    setTimeout(() => isSuccess.value = false, 2000)
  }catch (error) {
    console.log(error)
}
}
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
  padding: 48px 24px 24px 24px;
  border-radius: 20px;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
}

.qr-wrapper {
  border-radius: 24px;
  overflow: hidden;
  background: black;
  padding: 8px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
}

.close-btn {
  position: absolute;
  top: 16px;
  right: 16px;
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

.close-btn:hover {
  background: rgba(255,255,255,0.1);
  color: #ffffff;
}

.save-qr {
  background-color: #383837;
  color: #e0e0dd;
  border: 1px solid #5c5c5a;
  border-radius: 8px;
  padding: 10px 20px;
  min-width: 120px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.save-qr.success {
  background-color: #4682B4;
  border-color: #4682B4;
}

.save-qr:hover {
  background-color: #4a4a48;
  border-color: #7a7a78;
  transform: translateY(-2px);
}

.save-qr:active {
  transform: translateY(0);
}

.save-qr.success:hover {
  background-color: #4682B4;
}
</style>