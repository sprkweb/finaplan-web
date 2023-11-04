<script lang="ts">
	import Chart from 'chart.js/auto';
	import { afterUpdate } from 'svelte';

	export let data: number[] | undefined;

	let canvas: HTMLCanvasElement;
    let chart: Chart;

	afterUpdate(() => {
		if (data != undefined) {
            if (chart != undefined) {
                chart.destroy()
            }
			chart = new Chart(canvas, {
				type: 'bar',
				data: {
					labels: data.map((_, i) => i),
					datasets: [
						{
							data: data,
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

{#if data != undefined}
	<canvas bind:this={canvas} />
{/if}
