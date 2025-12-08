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
            type: "radar",
            data: {
                labels: data.labels,
                datasets: [
                    {
                        label: "Period",
                        data: data.values,
                        backgroundColor: "rgba(8, 92, 167, 0.70)",
                    }
                ]
            },
            options: {
                layout: {
                    padding: 15,
                },
                scales: {
                    r: {
                        ticks: {
                            display: false,
                        },
                        angleLines: {
                            display: false,
                        },
                        grid: {
                            color: "rgba(255, 255, 255, 0.2)",
                        },
                        pointLabels: {
                            color: "rgba(255, 255, 255, 0.5)",
                        },
                    },
                },
                plugins: {
                    legend: {
                        position: "bottom",
                        labels: {
                            color: "rgba(255, 255, 255, 0.5)",
                        },
                    },
                },
            },
        });
    });
</script>

<div class="relative h-fit w-full">
    <canvas bind:this={canvas}></canvas>
</div>
