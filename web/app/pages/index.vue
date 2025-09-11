<template>
  <HomeViewSelector/>
  <section v-if="$route.query.view === 'daily'" class="flex flex-col gap-6">
    <section>
      <h3 class="text-2xl">2025-09-11</h3>
      <p>Winner: Bob</p>
    </section>

    <section class="carousel w-full bg-base-900 gap-2 md:gap-0 md:rounded-lg">
      <HomeDailyGoalCard v-for="goal in [0, 2, 3, 4]" :goal="goal"/>
    </section>

    <section class="flex flex-col gap-2 w-full md:mx-auto">
      <HomeDailyWorkoutCard v-for="workout in [0, 2, 3, 4]" :workout="workout"/>
    </section>
  </section>

  <section v-else class="flex flex-col gap-6">
    <section class="grid grid-cols-2 gap-3">
      <HomeOverviewStatCard v-for="statCard in [0, 2, 3, 4]" :stat="statCard"/>
    </section>
  </section>
</template>

<script setup lang="ts">
</script>
