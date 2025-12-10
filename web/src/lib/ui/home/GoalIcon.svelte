<script lang="ts">
    import {CookingPot, Droplet, Dumbbell, Flame, HeartPulse, CircleQuestionMark} from "@lucide/svelte";
    import {GoalType, goalTypeToJSON} from "$lib/gen/zelvy/goal/goal_type_enum";

    type Props = {
        type: string
        index: number
    }

    const {type, index}: Props = $props();
    const backgroundColor = index % 2 === 0 ? "bg-primary/30" : "bg-secondary/30";
    const color = index % 2 === 0 ? "text-primary" : "text-secondary";
    let IconComponent = $state(CircleQuestionMark);
    switch (type) {
        case goalTypeToJSON(GoalType.KCAL_CONSUMED):
            IconComponent = CookingPot;
            break;
        case goalTypeToJSON(GoalType.MAIN_WORKOUT_DURATION):
            IconComponent = Dumbbell;
            break;
        case goalTypeToJSON(GoalType.EXTRA_WORKOUT_DURATION):
            IconComponent = HeartPulse;
            break;
        case goalTypeToJSON(GoalType.KCAL_BURNED):
            IconComponent = Flame;
            break;
        case goalTypeToJSON(GoalType.MILLILITER_DRANK):
            IconComponent = Droplet;
            break;
    }

</script>


<div class="hidden lg:flex w-12 h-12 justify-center items-center rounded-lg {backgroundColor}">
    <IconComponent class={color}/>
</div>
