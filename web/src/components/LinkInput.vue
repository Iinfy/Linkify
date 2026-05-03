<template>
  <div class="link-input-wrapper">
    <input
        type="text"
        class="link-input"
        placeholder="https://github.com/iinfy/linkify"
    />
    <button class="short-btn">Create</button>
  </div>
</template>

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
</style>