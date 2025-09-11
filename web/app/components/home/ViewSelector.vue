<script setup lang="ts">
const updateView = async (view: "daily" | "overview") => {
  await navigateTo({
    path: "/",
    query: {
      view: view
    }
  });
};
</script>

<template>
  <div class="join w-full bg-base-200 rounded-full">
    <input
        class="join-item btn rounded-full w-1/2 h-8"
        type="radio"
        name="views"
        aria-label="Daily"
        :checked="$route.query.view === 'daily'"
        @change="updateView('daily')"
    />
    <input
        class="join-item btn rounded-full w-1/2 h-8"
        type="radio"
        name="views"
        aria-label="Overview"
        :checked="$route.query.view === 'overview'"
        @change="updateView('overview')"
    />
  </div>
</template>

<style scoped>

</style>
