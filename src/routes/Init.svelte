<script lang="ts" context="module">
    export interface InitParams {
        intervalType: string
        intervalLength: number
        intervalAmount: number
    }
</script>

<script lang="ts">
    import { getContext } from 'svelte'
    import type { Writable } from 'svelte/store'

    const initParams = getContext<Writable<InitParams>>('init-params')
</script>

<div class="row align-items-center">
    <div class="col-auto">Plan duration</div>
    <div class="col-auto">
        <div class="input-group">
            <input
                type="number"
                class="form-control"
                min="1"
                bind:value={$initParams.intervalAmount}
                size="15"
            />
            <select class="form-select" bind:value={$initParams.intervalType}>
                <option value="days">days</option>
                <option value="weeks">weeks</option>
                <option value="months">months</option>
                <option value="years">years</option>
            </select>
        </div>
    </div>
</div>
