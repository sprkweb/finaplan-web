<script lang="ts">
	import '$lib/go/wasm_exec';
	import wasm from '$lib/go/main.wasm?url';
	import { onMount } from 'svelte';
	import type { Finaplan } from '$lib/types';
	import Print from '$lib/steps/Print.svelte';
	import Init, { type InitParams } from '$lib/steps/Init.svelte';
	import Add, { type AddParams } from '$lib/steps/Add.svelte';
	import Invest, { type InvestParams } from '$lib/steps/Invest.svelte';

	onMount(() => {
		console.log('Starting load');
		if (WebAssembly) {
			const go = new Go(); // eslint-disable-line no-undef

			WebAssembly.instantiateStreaming(fetch(wasm), go.importObject).then((result) => {
				console.log('Starting WebAssembly');
				go.run(result.instance);
			});
		} else {
			console.log('WebAssembly is not supported in your browser');
		}
	});

	enum StepType {
		Add = 'add',
		Invest = 'invest',
		Print = 'print'
	}

	interface Step {
		name: string;
		type: StepType;
		addParams?: AddParams;
		investParams?: InvestParams;
		output?: number[];
	}

    let initParams: InitParams;

	let steps: Step[] = [
		{
			name: 'Start capital',
			type: StepType.Add,
			addParams: {
				amount: 1000,
				each: 0,
				start: 0
			}
		},
		{
			name: 'Income',
			type: StepType.Add,
			addParams: {
				amount: 100,
				each: 1,
				start: 0
			}
		},
		{
			name: 'Invest',
			type: StepType.Invest,
			investParams: {
				interest: 1.1,
				interval: 12,
				start: 0,
				compound: true
			}
		},
		{
			name: 'Print',
			type: StepType.Print
		}
	];

	let calc = () => {
		let plan: Finaplan = initFinaplan(initParams.intervalType, initParams.intervalLength, initParams.intervalAmount); // eslint-disable-line no-undef
		steps.forEach((step) => {
			switch (step.type) {
				case StepType.Add:
					if (step.addParams == undefined) {
						break;
					}
					plan.add(step.addParams.amount, step.addParams.each, step.addParams.start);
					break;

				case StepType.Invest:
					if (step.investParams == undefined) {
						break;
					}
					plan.invest(
						step.investParams.interest,
						step.investParams.interval,
						step.investParams.start,
						step.investParams.compound
					);
					break;

				case StepType.Print:
					step.output = plan.print();
					break;
			}
		});
		steps = steps;
	};
</script>

<div class="container">
	<div class="card my-3">
		<div class="card-body">
            <Init bind:value={initParams} />
			<button type="submit" class="btn btn-primary" on:click={calc}>Calculate</button>
		</div>
	</div>

	{#each steps as step}
		<div class="card my-3">
			<div class="card-body">
				<h5 class="card-title" bind:innerHTML={step.name} contenteditable />

				{#if step.type == StepType.Add}
                    <Add bind:value={step.addParams} />
				{:else if step.type == StepType.Invest}
                    <Invest bind:value={step.investParams} />
				{:else if step.type == StepType.Print}
					<Print value={step.output} />
				{/if}
			</div>
		</div>
	{/each}
</div>

<style>
</style>
