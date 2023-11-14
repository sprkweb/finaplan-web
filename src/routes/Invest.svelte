<script lang="ts" context="module">
    export interface InvestParams {
        interest: string
        interval: number
        start: number
        compound: boolean
    }
</script>

<script lang="ts">
    import { getContext } from 'svelte'
    import type { Writable } from 'svelte/store'
    import type { InitParams } from './Init.svelte'

    const initParams = getContext<Writable<InitParams>>('init-params')

    export let value: InvestParams = {
        interest: '10',
        interval: 12,
        start: 0,
        compound: true
    }
</script>

<div class="row g-3 mb-3 align-items-center">
    <div class="col">
        <form class="form-floating">
            <input
                type="text"
                class="form-control"
                id="interest"
                pattern="-?\d+(\.\d+)?%?"
                title="Percent (4%) or decimal number (0.04). Can be negative"
                bind:value={value.interest}
            />
            <label for="interest">Rate</label>
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
                bind:value={value.interval}
            />
            <label for="interval">Interval</label>
        </form>
    </div>
    <div class="col-auto">{$initParams.intervalType}</div>
</div>

<form class="form-floating mb-3">
    <input
        type="number"
        class="form-control"
        id="start"
        bind:value={value.start}
    />
    <label for="start">Start at</label>
</form>

<div class="form-check">
    <input
        class="form-check-input"
        type="checkbox"
        id="repeat"
        bind:checked={value.compound}
    />
    <label class="form-check-label" for="compound">Compound interest</label>
</div>
