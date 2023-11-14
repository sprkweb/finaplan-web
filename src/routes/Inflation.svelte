<script lang="ts" context="module">
    export interface InflationParams {
        inflation: string
        intervals: number
    }
</script>

<script lang="ts">
    import { getContext } from 'svelte'
    import type { Writable } from 'svelte/store'
    import type { InitParams } from './Init.svelte'

    const initParams = getContext<Writable<InitParams>>('init-params')

    export let value: InflationParams = {
        inflation: '4',
        intervals: 12
    }
</script>

<div class="row g-3 mb-3 align-items-center">
    <div class="col">
        <form class="form-floating">
            <input
                type="text"
                class="form-control"
                id="inflation"
                pattern="-?\d+(\.\d+)?%?"
                title="Percent (4%) or decimal number (0.04). Can be negative"
                bind:value={value.inflation}
            />
            <label for="inflation">Inflation rate</label>
        </form>
    </div>
    <div class="col-auto">per</div>
    <div class="col">
        <form class="form-floating">
            <input
                type="number"
                class="form-control"
                id="interval"
                min="0"
                bind:value={value.intervals}
            />
            <label for="interval">Interval</label>
        </form>
    </div>
    <div class="col-auto">{$initParams.intervalType}</div>
</div>
