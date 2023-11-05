<script lang="ts" context="module">
    export interface InitParams {
        intervalType: string
        intervalLength: number
        intervalAmount: number
    }
</script>

<script lang="ts">
    export let value: InitParams = {
        intervalType: 'months',
        intervalLength: 1,
        intervalAmount: 12
    }
</script>

<div class="row align-items-center mb-3">
    <div class="col-auto">Plan duration</div>
    <div class="col-auto">
        <div class="input-group">
            <input
                type="number"
                class="form-control"
                min="1"
                bind:value={value.intervalAmount}
                size="15"
            />
            <select class="form-select" bind:value={value.intervalType}>
                <option value="days">days</option>
                <option value="weeks">weeks</option>
                <option value="months">months</option>
                <option value="years">years</option>
            </select>
        </div>
    </div>
</div>
