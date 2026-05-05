<template>
  <div class="link-input-wrapper">
    <input
        ref="linkInput"
        type="text"
        class="link-input"
        placeholder="https://github.com/iinfy/linkify"
    />
    <button ref="shortBtn" class="short-btn" :class="{'success': isSuccess}" @click="hadleShortener">
      {{ buttonText }}
    </button>
  </div>
</template>
<script setup lang="ts">
import { useNotification } from '@/composables/useNotification'
const { show } = useNotification()
import {computed, ref} from "vue";

  const shortBtn = document.querySelector('.short-btn');
  const linkInput = ref<HTMLInputElement | null>(null);
  const base_url = window.location.origin;
  const buttonText = ref<string>("Create");
  const isSuccess = ref(false)

  const emit = defineEmits(['linkCreated']);
  const hadleShortener = async () => {
    let url = ""
    if (linkInput.value) {
      url = linkInput.value.value.trim()
    }
    try {
      if (!url.startsWith('https://') && !url.startsWith('http://')) {
        show("Invalid input", "URL must start with https/http", true)
        console.log('Url is not https/http')
        return
      }

      const res = await fetch(`/slink`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({url: url})
      })

      const data = await res.json()
      if (!res.ok) {
        show("Error", "Internal server error", true)
        return
      }



      buttonText.value = "Created";
      isSuccess.value = true
      setTimeout(() => isSuccess.value = false, 2000)

      setTimeout(() => {
        buttonText.value = "Create";
      }, 2000 )

      emit("linkCreated", {original: url, short: base_url + "/s/" + data.url})

    } catch (error) {
      console.error('Error:', error)

  }
  }
</script>

<style scoped>
.link-input-wrapper {
  position: relative;
  width: 100%;
  max-width: 500px;
}

.link-input-wrapper input {
  width: 100%;
  height: 60px;
  background-color: #212121;
  border: 1px solid #474745;
  border-radius: 20px;
  padding: 0 100px 0 20px;
  font-size: clamp(14px, 2.5vw, 17px);
  color: #e0e0dd;
  outline: none;
}

.link-input-wrapper input::placeholder {
  color: #5F5F5F;
  font-size: clamp(14px, 2.5vw, 17px);
}

.link-input-wrapper .short-btn {
  position: absolute;
  width: 5vw;
  top: 50%;
  right: 8px;
  transform: translateY(-50%);
  height: 44px;
  padding: 0 20px;
  background-color: #383837;
  color: #e0e0dd;
  border: 1px solid #5c5c5a;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: clamp(13px, 2vw, 15px);
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.4s ease;
}

.link-input-wrapper .short-btn:hover {
  background-color: #323232;
}

.short-btn.success {
  width: 5vw;
  background-color: green;
}
.short-btn.success:hover {

  background-color: green;
}

</style>