<script lang="ts">
    import {onMount} from "svelte";
    import Chart from "chart.js/auto";

    type Props = {
        data: {
            labels: string[],
            values: number[],
        }
    }

    let {data}: Props = $props();
    let canvas: HTMLCanvasElement;

    onMount(() => {
        if (!canvas) {
            return;
        }

        return new Chart(canvas, {
            type: "line",
            data: {
                labels: data.labels,
                datasets: [
                    {
                        label: "Period",
                        data: data.values,
                        backgroundColor: "rgba(8, 92, 167, 0.70)",
                        type: "bar"

                    }
                ]
            },
            options: {
                layout: {
                    padding: 0,
                },
                plugins: {
                    legend: {
                        display: false,
                    },
                },
            },
        });
    });
</script>

<div class="flex items-center justify-center bg-base-200 rounded-lg p-6">
    <canvas bind:this={canvas}></canvas>
</div>
