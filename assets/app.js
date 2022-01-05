const go = new Go();

WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then(
	(result) => {
		go.run(result.instance);
	}
);

// EXEMPLE 1

const number1 = document.querySelector('#number1');
const number2 = document.querySelector('#number2');
const btn = document.querySelector('#btn');
const form1 = document.querySelector('#ex1');
const span = document.querySelector('#result');

form1.addEventListener('submit', (e) => {
	e.preventDefault();

	const result = add(+number1.value, +number2.value);
	span.textContent = result;

	number1.value = number2.value = '';
});

// EXEMPLE 2

const userInput = document.querySelector('#userInput');
const form2 = document.querySelector('#ex2');
const capitalSpan = document.querySelector('#capitalized');

form2.addEventListener('submit', (e) => {
	e.preventDefault();

	const capitalStr = capitalized(userInput.value);

	capitalSpan.textContent = capitalStr;

	userInput.value = '';
});