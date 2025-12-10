<script lang="ts">
    import Chart from "chart.js/auto";
    import {afterNavigate} from "$app/navigation";

    type Props = {
        data: {
            labels: string[],
            values: number[],
        }
    }

    const {data}: Props = $props();
    let canvas: HTMLCanvasElement;
    let chart: Chart<"bar">;

    afterNavigate(() => {
        if (!canvas) {
            return;
        }

        if (chart) {
            chart.destroy();
        }

        chart = new Chart(canvas, {
            type: "bar",
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
                plugins: {
                    legend: {
                        display: false,
                    },
                },
            },
        });
    });
</script>

<div class="flex items-center justify-center">
    <canvas bind:this={canvas}></canvas>
</div>
