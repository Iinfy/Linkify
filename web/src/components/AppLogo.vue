<template>
  <div class="logo">
    <h1 class="title">Linkify</h1>
    <span class="subtitle">URL Shortener</span>
  </div>
</template>

<style scoped>
.logo {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: white;
  margin-bottom: 24px;
}

.logo .title {
  font-size: clamp(40px, 10vw, 72px);
  font-weight: 700;
  line-height: 1;
  margin: 0;
  text-align: center;
}

.logo .subtitle {
  color: white;
  font-size: clamp(14px, 3vw, 20px);
  font-weight: 600;
  padding: 4px 12px;
  border-radius: 4px;
  letter-spacing: 0.5px;
}

@media (max-width: 480px) {
  .logo {
    margin-bottom: 20px;
  }

}
</style>