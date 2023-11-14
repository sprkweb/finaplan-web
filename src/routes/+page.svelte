<script lang="ts">
    import { onMount } from 'svelte'

    import { loadFinaplan, type Finaplan } from './finaplan'

    import Print from './Print.svelte'
    import Init, { type InitParams } from './Init.svelte'
    import Add, { type AddParams } from './Add.svelte'
    import Invest, { type InvestParams } from './Invest.svelte'
    import type { InflationParams } from './Inflation.svelte'
    import Inflation from './Inflation.svelte'

    onMount(() => {
        loadFinaplan()
            .then(() => calc())
            .catch(e => console.log(e))
    })

    enum StepType {
        Add = 'add',
        Invest = 'invest',
        Inflation = 'inflation',
        Print = 'print'
    }

    interface Step {
        name: string
        type: StepType
        addParams?: AddParams
        investParams?: InvestParams
        inflationParams?: InflationParams
        error?: string
        output?: number[]
    }

    let initParams: InitParams

    let steps: Step[] = [
        {
            name: 'Start capital',
            type: StepType.Add,
            addParams: {
                amount: '1000',
                each: 0,
                start: 0
            }
        },
        {
            name: 'Income',
            type: StepType.Add,
            addParams: {
                amount: '100',
                each: 1,
                start: 0
            }
        },
        {
            name: 'Invest',
            type: StepType.Invest,
            investParams: {
                interest: '10%',
                interval: 12,
                start: 0,
                compound: true
            }
        },
        {
            name: 'Inflation',
            type: StepType.Inflation,
            inflationParams: {
                inflation: '4%',
                intervals: 12
            }
        },
        {
            name: 'Print',
            type: StepType.Print
        }
    ]

    let calc = () => {
        let plan: Finaplan = window.InitFinaplan(
            initParams.intervalType,
            initParams.intervalLength,
            initParams.intervalAmount
        )
        steps.every(step => {
            switch (step.type) {
                case StepType.Add:
                    if (step.addParams == undefined) {
                        break
                    }
                    {
                        const { error } = plan.add(
                            step.addParams.amount,
                            step.addParams.each,
                            step.addParams.start
                        )
                        if (error != undefined) {
                            step.error = error
                            return false
                        } else {
                            step.error = undefined
                        }
                    }
                    break

                case StepType.Invest:
                    if (step.investParams == undefined) {
                        break
                    }
                    {
                        const { error } = plan.invest(
                            step.investParams.interest,
                            step.investParams.interval,
                            step.investParams.start,
                            step.investParams.compound
                        )
                        if (error != undefined) {
                            step.error = error
                            return false
                        } else {
                            step.error = undefined
                        }
                    }
                    break

                case StepType.Inflation:
                    if (step.inflationParams == undefined) {
                        break
                    }
                    {
                        const { error } = plan.inflation(
                            step.inflationParams.inflation,
                            step.inflationParams.intervals
                        )
                        if (error != undefined) {
                            step.error = error
                            return false
                        } else {
                            step.error = undefined
                        }
                    }
                    break

                case StepType.Print:
                    {
                        const { result, error } = plan.print()
                        if (error != undefined) {
                            step.error = error
                            return false
                        } else {
                            step.output = result
                        }
                    }
                    break
            }
            return true
        })
        steps = steps
    }
</script>

<div class="navbar sticky-top bg-body-tertiary">
    <div class="container">
        <a class="navbar-brand" href=".">Finaplan</a>
        <button type="submit" class="btn btn-primary m-0" on:click={calc}>
            Calculate
        </button>
    </div>
</div>

<div class="container">
    <div class="card my-3">
        <div class="card-body">
            <Init bind:value={initParams} />
        </div>
    </div>

    {#each steps as step}
        <div class="card my-3">
            <div class="card-body">
                <h5
                    class="card-title"
                    bind:innerHTML={step.name}
                    contenteditable
                />

                {#if step.error}
                    <div class="error">{step.error}</div>
                {/if}

                {#if step.type == StepType.Add}
                    <Add bind:value={step.addParams} />
                {:else if step.type == StepType.Invest}
                    <Invest bind:value={step.investParams} />
                {:else if step.type == StepType.Inflation}
                    <Inflation bind:value={step.inflationParams} />
                {:else if step.type == StepType.Print}
                    <Print value={step.output} />
                {/if}
            </div>
        </div>
    {/each}
</div>

<style>
    .error {
        color: red;
    }
</style>
