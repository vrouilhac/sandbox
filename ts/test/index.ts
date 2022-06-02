const test: Test = {
	a: 34
};

interface Test {
	a: number;
}

const t = (tes?: Test) => {
	const isValid = 
		tes !== undefined && tes !== null && !isNaN(tes.a);

	const value: number = isValid ? tes.a : 0;

};
