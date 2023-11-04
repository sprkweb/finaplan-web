interface Finaplan {
	add: (amount: number, each: number, start: number) => void;
	invest: (interest: number, interval: number, start: number, compound: boolean) => void;
	print: () => number[];
}

export { Finaplan };
