<script lang="ts">
    import {GoalViewModel} from "$lib/gen/zelvy/goal/goal_view_model";

    type Props = {
        goal: GoalViewModel
    }

    const {goal}: Props = $props();
</script>

<article data-testid="summaryDetailsTemplateMetric"
         class="carousel-item bg-base-200 flex-col px-4 justify-evenly rounded-lg aspect-square w-[43%]
           md:rounded-none md:px-2 md:aspect-auto md:h-56 md:w-1/5">
    <p data-testid="summaryDetailsTemplateMetricName" class="text-base text-base-content/60 md:text-lg">
        {goal.name}
    </p>
    <h2 data-testid="summaryDetailsTemplateMetricValue" class="text-3xl font-semibold">{goal.displayValue}</h2>
    <div class="flex flex-col gap-1">
        <p class="text-base-content/60 text-sm md:text-lg">Goal: <span
                data-testid="summaryDetailsTemplateMetricThreshold">{goal.displayThreshold}</span></p>
        <div class="relative h-2 w-full">
            <div class="bg-primary h-2 rounded-l-full absolute top-0 left-0 {goal.progression === 100 ? 'rounded-r-full' : ''}"
                 style="width:{goal.progression}%"
            >
            </div>
            <div class="bg-primary-content h-2 w-full rounded-full"></div>
        </div>
    </div>
</article>
