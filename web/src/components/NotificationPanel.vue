<script setup lang="ts">
import { useNotification } from '@/composables/useNotification'
const { visible, title,subtitle } = useNotification()

</script>

<template>
  <transition name="slide">
  <div v-if="visible" class="notif-container">
    <img src="../assets/icons/check.png" class="check" alt="">
    <div class="text-group">
      <h2>{{title}}</h2>
      <p>{{subtitle}}</p>
    </div>
  </div>
  </transition>
</template>

<style scoped>

* {
  margin: 0;
  padding: 0;

}

.notif-container {
  display: flex;
  align-items: center;
  position: fixed;
  bottom: 0;
  right: 0;
  padding: 10px 15px;
  margin-right: 15px;
  margin-bottom: 15px;
  border: 1px solid #474745;
  background: #212121;
  height: 70px;
  width: 20vw;
  color: white;
  border-radius: 20px;


}

.text-group {
  display: flex;
  flex-direction: column;
  margin-left: 45px;
}

.notif-container h2 {
  font-size: 15px;
  font-weight: 500;
  line-height: 1.2;
}

.notif-container p {
  margin-top: 4px;
  font-size: 13px;
  color: #5F5F5F;
  line-height: 1.2;
}

.check {
  position: absolute;
  left: 20px;
  top: 50%;
  transform: translateY(-50%);
  width: 22px;
  height: 22px;
  object-fit: contain;
  filter: invert(1);
}

.slide-enter-active {
  transition: all 0.4s ease;
}

.slide-leave-active {
  transition: all 0.4s ease;
}

.slide-enter-from {
  transform: translateX(100%);
}

.slide-leave-to {
  transform: translateX(100%);
}

</style>