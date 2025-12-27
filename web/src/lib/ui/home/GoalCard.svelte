<script lang="ts">
    import {GoalViewModel} from "$lib/gen/zelvy/goal/goal_view_model";
    import GoalIcon from "$lib/ui/home/GoalIcon.svelte";

    type Props = {
        goal: GoalViewModel
        index: number
        dataTestId: string
    }

    const {goal, index, dataTestId}: Props = $props();
</script>

<article
        class="carousel-item bg-base-200 flex-col px-4 justify-evenly rounded-lg aspect-square lg:aspect-auto lg:rounded-lg w-[33%]
           md:rounded-none md:h-45 md:w-1/5" data-testid={dataTestId}>
    <div class="flex flex-row justify-between">
        <div class="flex flex-col gap-2">
            <h2 data-testid="summaryDetailsTemplateMetricValue" class="text-2xl font-semibold">{goal.displayValue}</h2>
            <p data-testid="summaryDetailsTemplateMetricName" class="text-base text-base-content/60">
                {goal.name}
            </p>
        </div>
        <GoalIcon type={goal.type} index={index}/>
    </div>
    <div class="flex flex-col gap-1">
        <div class="relative h-2 w-full">
            <div class="bg-primary h-2 rounded-l-full absolute top-0 left-0 {goal.progression === 100 ? 'rounded-r-full' : ''}"
                 style="width:{goal.progression}%"
            >
            </div>
            <div class="bg-primary-content h-2 w-full rounded-full"></div>
        </div>
    </div>
</article>
