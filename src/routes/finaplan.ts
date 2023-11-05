// This file loads WASM implementation of Finaplan

import '$lib/go/wasm_exec'
import wasm from '$lib/go/main.wasm?url'

declare global {
    export interface Window {
        Go: {
            new (): {
                run: (inst: WebAssembly.Instance) => Promise<void>
                importObject: WebAssembly.Imports
            }
        }
        initFinaplan: (
            intervalType: string,
            intervalLength: number,
            intervalAmount: number
        ) => Finaplan
    }
}

export interface Finaplan {
    add: (amount: number, each: number, start: number) => void
    invest: (
        interest: number,
        interval: number,
        start: number,
        compound: boolean
    ) => void
    print: () => number[]
}

const until = (f: () => boolean): Promise<void> => {
    return new Promise(resolve => {
        const intervalCode = setInterval(() => {
            if (f()) {
                resolve()
                clearInterval(intervalCode)
            }
        }, 10)
    })
}

export async function loadFinaplan() {
    if (!WebAssembly.instantiateStreaming) {
        // polyfill
        WebAssembly.instantiateStreaming = async (resp, importObject) => {
            const source = await (await resp).arrayBuffer()
            return await WebAssembly.instantiate(source, importObject)
        }
    }

    console.log('Loading WebAssembly')

    if (!WebAssembly) {
        throw new Error('WebAssembly is not supported in your browser')
    }
    const go = new window.Go()
    const result = await WebAssembly.instantiateStreaming(
        fetch(wasm),
        go.importObject
    )

    console.log('Loaded WebAssembly')

    go.run(result.instance)

    // wait until WASM create the function
    await until(() => window.initFinaplan != undefined)
    return window.initFinaplan
}
