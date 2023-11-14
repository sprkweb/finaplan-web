<script lang="ts" context="module">
    export interface AddParams {
        amount: string
        each: number
        start: number
    }
</script>

<script lang="ts">
    import { getContext } from 'svelte'
    import type { Writable } from 'svelte/store'
    import type { InitParams } from './Init.svelte'

    const initParams = getContext<Writable<InitParams>>('init-params')

    export let value: AddParams = {
        amount: '100',
        each: 1,
        start: 0
    }
</script>

<div class="row g-3 mb-3">
    <div class="col">
        <form class="form-floating">
            <input
                type="text"
                class="form-control"
                id="amount"
                pattern="-?\d+(\.\d+)?"
                title="Decimal number (4321.05). Can be negative"
                bind:value={value.amount}
            />
            <label for="amount">Amount of money</label>
        </form>
    </div>
    <div class="col">
        <form class="form-floating">
            <input
                type="number"
                class="form-control"
                id="start"
                min="0"
                bind:value={value.start}
            />
            <label for="start">Start at</label>
        </form>
    </div>
</div>

<div class="row align-items-center">
    <div class="col-auto">
        <div class="form-check">
            <input
                class="form-check-input"
                type="checkbox"
                id="repeat"
                checked={value.each != 0}
                on:change={e => (value.each = e.target?.checked ? 1 : 0)}
            />
            <label class="form-check-label" for="repeat">
                Repeat {#if value.each != 0}each{/if}
            </label>
        </div>
    </div>
    <div class="col">
        <input
            type="number"
            class="form-control"
            class:invisible={value.each == 0}
            min="1"
            bind:value={value.each}
        />
    </div>
    <div class="col-auto" class:invisible={value.each == 0}>
        {$initParams.intervalType}
    </div>
</div>
