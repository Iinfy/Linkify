<script setup lang="ts">
import LinkItem from "@/components/links/LinkItem.vue";
import AppLogo from "@/components/AppLogo.vue";
import StatsPanel from "@/components/StatsPanel.vue";
import { useRoute } from "vue-router";

const baseUrl = window.location.origin;
const route = useRoute();
const hash = route.params.hash as string;
</script>

<template>
  <div class="stats-page-wrapper">
    <AppLogo />
    <div class="main__section">
      <div class="linkitem-container">
        <LinkItem class="linkitem" :short="baseUrl + '/s/' + hash"></LinkItem>
      </div>
      <div class="stats-container">
        <StatsPanel :hash="hash" />
      </div>
    </div>
  </div>
</template>

<style>
@import "@/assets/main.css";

.stats-page-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
}

.stats-page-wrapper .main__section {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  width: 100%;
  max-width: 500px;
}

.linkitem-container {
  width: 100%;
  display: flex;
  justify-content: center;
  margin-bottom: 24px;

  .linkitem {
    width: 100%;
  }
}

.stats-container {
  width: 100%;
  display: flex;
  justify-content: center;

  .statsPanel {
    background-color: #212121;
    border: 1px solid #474745;
    border-radius: 20px;
    width: 100%;
  }
}
</style>
