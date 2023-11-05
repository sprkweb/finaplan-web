<script lang="ts">
	import Chart from 'chart.js/auto';
	import { afterUpdate } from 'svelte';

	export let value: number[] | undefined;

	let canvas: HTMLCanvasElement;
    let chart: Chart;

	afterUpdate(() => {
		if (value != undefined) {
            if (chart != undefined) {
                chart.destroy()
            }
			chart = new Chart(canvas, {
				type: 'bar',
				data: {
					labels: value.map((_, i) => i),
					datasets: [
						{
							data: value,
							borderWidth: 1
						}
					]
				},
				options: {
					animation: false,
					scales: {
						y: {
							beginAtZero: true
						}
					},
					plugins: {
						legend: {
							display: false
						}
					}
				}
			});
		}
	});
</script>

{#if value != undefined}
	<canvas bind:this={canvas} />
{/if}
