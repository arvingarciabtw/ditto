<script>
	import { onMount } from "svelte";

	const ignoredKeys = new Set([
		"Alt",
		"AltGraph",
		"CapsLock",
		"Control",
		"Fn",
		"Meta",
		"NumLock",
		"ScrollLock",
		"Shift"
	]);
	/** @type {Record<string, string>} */
	const fingerByCode = Object.fromEntries(
		Object.entries({
			pinky:
				"Escape F1 F10 F11 F12 Backquote Digit1 Digit0 Minus Equal Backspace Tab KeyQ KeyP BracketLeft BracketRight Backslash CapsLock KeyA Semicolon Quote Enter ShiftLeft KeyZ Slash ShiftRight ControlLeft ControlRight NumpadSubtract NumpadAdd NumpadEnter",
			ring: "F2 F9 Digit2 Digit9 KeyW KeyO KeyS KeyL KeyX Period PageUp PageDown NumpadMultiply Numpad9 Numpad6 Numpad3 ArrowRight Numpad0 NumpadDecimal",
			middle:
				"F3 F8 Pause Digit3 Digit8 Home End KeyE KeyI KeyD KeyK KeyC Comma NumpadDivide Numpad8 Numpad5 ArrowUp Numpad2 ArrowDown",
			index:
				"F4 F5 F6 F7 PrintScreen ScrollLock Digit4 Digit5 Digit6 Digit7 Insert Delete NumLock KeyR KeyT KeyY KeyU KeyF KeyG KeyH KeyJ KeyV KeyB KeyN KeyM Numpad7 Numpad4 Numpad1 ArrowLeft",
			thumb: "Space"
		}).flatMap(([finger, codes]) => codes.split(" ").map((code) => [code, finger]))
	);
	let key = $state("");
	let finger = $state("");
	let visible = $state(false);

	onMount(() => {
		let hideTimer = 0;
		/** @type {(event: KeyboardEvent) => void} */
		const cast = (event) => {
			if (ignoredKeys.has(event.key)) return;

			key = event.key === " " ? "Space" : event.key;
			finger = fingerByCode[event.code] ?? "";
			visible = true;
			window.clearTimeout(hideTimer);
			hideTimer = window.setTimeout(() => (visible = false), 1500);
		};

		window.addEventListener("keydown", cast);

		return () => {
			window.removeEventListener("keydown", cast);
			window.clearTimeout(hideTimer);
		};
	});
</script>

<footer>
	<p class="developer">
		by <a href="https://arvingarcia.com" target="_blank" rel="external noopener noreferrer"
			>@arvingarciabtw</a
		>
	</p>
	<p class="source">
		<a
			href="https://github.com/arvingarciabtw/ditto"
			target="_blank"
			rel="external noopener noreferrer">@source</a
		>
	</p>
	<p class="coffee">
		<a href="https://ko-fi.com/arvingarciabtw" target="_blank" rel="external noopener noreferrer"
			>@ko-fi</a
		>
	</p>
	<kbd class="keycaster {finger}" class:visible aria-hidden="true">{key}</kbd>
</footer>

<style>
	footer {
		padding-top: 2rem;
		display: grid;
		grid-template-columns: repeat(3, max-content) 1fr;
		align-items: center;
		gap: 2rem;
	}
	.keycaster {
		justify-self: end;
		min-width: 2rem;
		height: 2.25rem;
		padding: 0 0.5rem;
		display: grid;
		place-items: center;
		border: 1px solid var(--finger-color, var(--overlay1));
		border-radius: 0.375rem;
		color: var(--finger-color, var(--text));
		font: inherit;
		line-height: 1.75;
		text-align: center;
		opacity: 0;
	}
	.keycaster.visible {
		opacity: 1;
	}
	.keycaster.pinky {
		--finger-color: var(--pink);
	}
	.keycaster.ring {
		--finger-color: var(--sapphire);
	}
	.keycaster.middle {
		--finger-color: var(--green);
	}
	.keycaster.index {
		--finger-color: var(--peach);
	}
	.keycaster.thumb {
		--finger-color: var(--red);
	}

	.developer a {
		color: var(--lavender);
	}
	.source a {
		color: var(--teal);
	}
	.coffee a {
		color: var(--peach);
	}
</style>
