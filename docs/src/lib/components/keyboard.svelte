<script>
	import { onMount } from "svelte";
	import { SvelteSet } from "svelte/reactivity";

	const pressedCodes = new SvelteSet();

	const keyboard = {
		// prettier-ignore
		top: ["╭───┬───┬───┬───┬───┬───┬──┬───┬───┬───┬───┬──┬───┬───┬───┬───┬──┬───┬───┬───┬──────────────────╮"],
		rows: [
			// prettier-ignore
			{
				content: ["│", 
					{ ch: "Esc", f: "pinky", code: "Escape"}, 
                    "│   │", 
					{ ch: "F1 ", f: "pinky", code: "F1"}, 
                    "│", 
					{ ch: "F2 ", f: "ring", code: "F2"}, 
                    "│", 
					{ ch: "F3 ", f: "middle", code: "F3"}, 
                    "│", 
					{ ch: "F4 ", f: "index", code: "F4"}, 
                    "│  │", 
					{ ch: "F5 ", f: "index", code: "F5"}, 
                    "│", 
					{ ch: "F6 ", f: "index", code: "F6"}, 
                    "│", 
					{ ch: "F7 ", f: "index", code: "F7"}, 
                    "│", 
					{ ch: "F8 ", f: "middle", code: "F8"}, 
                    "│  │", 
					{ ch: "F9 ", f: "ring", code: "F9"}, 
                    "│", 
					{ ch: "F10", f: "pinky", code: "F10"}, 
                    "│", 
					{ ch: "F11", f: "pinky", code: "F11"}, 
                    "│", 
					{ ch: "F12", f: "pinky", code: "F12"}, 
                    "│  │",
					{ ch: "Prt", f: "index", code: "PrintScreen"}, 
                    "│",
					{ ch: "Slk", f: "index", code: "ScrollLock"}, 
                    "│",
					{ ch: "Fls", f: "middle", code: "Pause"}, 
                    "│                  │",
                ],
				sep: [
					"├───┴───┴───┴───┴───┴───┴──┴───┴───┴───┴───┴──┴───┴───┴───┴───┤  ├───┴───┴───┤  ╭───┬───┬───┬───┤"
				]
			},
			{
				content: [
					"│",
					{ ch: " ` ", f: "pinky", code: "Backquote" },
					"│",
					{ ch: " 1 ", f: "pinky", code: "Digit1" },
					"│",
					{ ch: " 2 ", f: "ring", code: "Digit2" },
					"│",
					{ ch: " 3 ", f: "middle", code: "Digit3" },
					"│",
					{ ch: " 4 ", f: "index", code: "Digit4" },
					"│",
					{ ch: " 5 ", f: "index", code: "Digit5" },
					"│",
					{ ch: " 6 ", f: "index", code: "Digit6" },
					"│",
					{ ch: " 7 ", f: "index", code: "Digit7" },
					"│",
					{ ch: " 8 ", f: "middle", code: "Digit8" },
					"│",
					{ ch: " 9 ", f: "ring", code: "Digit9" },
					"│",
					{ ch: " 0 ", f: "pinky", code: "Digit0" },
					"│",
					{ ch: " - ", f: "pinky", code: "Minus" },
					"│",
					{ ch: " = ", f: "pinky", code: "Equal" },
					"│",
					{ ch: "   <--   ", f: "pinky", code: "Backspace" },
					"│  ",
					"│",
					{ ch: "Ins", f: "index", code: "Insert" },
					"│",
					{ ch: "Hme", f: "middle", code: "Home" },
					"│",
					{ ch: "PgU", f: "ring", code: "PageUp" },
					"│  ",
					"│",
					{ ch: "Nlk", f: "index", code: "NumLock" },
					"│",
					{ ch: " / ", f: "middle", code: "NumpadDivide" },
					"│",
					{ ch: " * ", f: "ring", code: "NumpadMultiply" },
					"│",
					{ ch: " - ", f: "pinky", code: "NumpadSubtract" },
					"│"
				],
				sep: [
					"├───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴─────────┤  ├───┴───┴───┤  ├───┴───┴───┴───┤"
				]
			},
			{
				content: [
					"│",
					{ ch: "  Tab  ", f: "pinky", code: "Tab" },
					"│",
					{ ch: " Q ", f: "pinky", code: "KeyQ" },
					"│",
					{ ch: " W ", f: "ring", code: "KeyW" },
					"│",
					{ ch: " E ", f: "middle", code: "KeyE" },
					"│",
					{ ch: " R ", f: "index", code: "KeyR" },
					"│",
					{ ch: " T ", f: "index", code: "KeyT" },
					"│",
					{ ch: " Y ", f: "index", code: "KeyY" },
					"│",
					{ ch: " U ", f: "index", code: "KeyU" },
					"│",
					{ ch: " I ", f: "middle", code: "KeyI" },
					"│",
					{ ch: " O ", f: "ring", code: "KeyO" },
					"│",
					{ ch: " P ", f: "pinky", code: "KeyP" },
					"│",
					{ ch: " [ ", f: "pinky", code: "BracketLeft" },
					"│",
					{ ch: " ] ", f: "pinky", code: "BracketRight" },
					"│",
					{ ch: "  \\  ", f: "pinky", code: "Backslash" },
					"│  │",
					{ ch: "Del", f: "index", code: "Delete" },
					"│",
					{ ch: "End", f: "middle", code: "End" },
					"│",
					{ ch: "PgD", f: "ring", code: "PageDown" },
					"│  │",
					{ ch: " 7 ", f: "index", code: "Numpad7" },
					"│",
					{ ch: " 8 ", f: "middle", code: "Numpad8" },
					"│",
					{ ch: " 9 ", f: "ring", code: "Numpad9" },
					"│   │"
				],
				sep: [
					"├───────┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴─────┤  ╰───┴───┴───╯  ├───┴───┴───┤ ",
					{ ch: "+", f: "pinky", code: "NumpadAdd" },
					" │"
				]
			},
			{
				content: [
					"│",
					{ ch: " Caps ", f: "pinky", code: "CapsLock" },
					"│",
					{ ch: " A ", f: "pinky", code: "KeyA" },
					"│",
					{ ch: " S ", f: "ring", code: "KeyS" },
					"│",
					{ ch: " D ", f: "middle", code: "KeyD" },
					"│",
					{ ch: " F ", f: "index", code: "KeyF" },
					"│",
					{ ch: " G ", f: "index", code: "KeyG" },
					"│",
					{ ch: " H ", f: "index", code: "KeyH" },
					"│",
					{ ch: " J ", f: "index", code: "KeyJ" },
					"│",
					{ ch: " K ", f: "middle", code: "KeyK" },
					"│",
					{ ch: " L ", f: "ring", code: "KeyL" },
					"│",
					{ ch: " ; ", f: "pinky", code: "Semicolon" },
					"│",
					{ ch: " ' ", f: "pinky", code: "Quote" },
					"│",
					{ ch: "  Enter   ", f: "pinky", code: "Enter" },
					"│                 │",
					{ ch: " 4 ", f: "index", code: "Numpad4" },
					"│",
					{ ch: " 5 ", f: "middle", code: "Numpad5" },
					"│",
					{ ch: " 6 ", f: "ring", code: "Numpad6" },
					"│   │"
				],
				sep: [
					"├──────┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴──────────┤      ╭───╮      ├───┴───┴───┴───┤"
				]
			},
			{
				content: [
					"│",
					{ ch: "   Shift   ", f: "pinky", code: "ShiftLeft" },
					"│",
					{ ch: " Z ", f: "pinky", code: "KeyZ" },
					"│",
					{ ch: " X ", f: "ring", code: "KeyX" },
					"│",
					{ ch: " C ", f: "middle", code: "KeyC" },
					"│",
					{ ch: " V ", f: "index", code: "KeyV" },
					"│",
					{ ch: " B ", f: "index", code: "KeyB" },
					"│",
					{ ch: " N ", f: "index", code: "KeyN" },
					"│",
					{ ch: " M ", f: "index", code: "KeyM" },
					"│",
					{ ch: " , ", f: "middle", code: "Comma" },
					"│",
					{ ch: " . ", f: "ring", code: "Period" },
					"│",
					{ ch: " / ", f: "pinky", code: "Slash" },
					"│",
					{ ch: "  Shift  ", f: "pinky", code: "ShiftRight" },
					"│      │",
					{ ch: " ↑ ", f: "middle", code: "ArrowUp" },
					"│      │",
					{ ch: " 1 ", f: "index", code: "Numpad1" },
					"│",
					{ ch: " 2 ", f: "middle", code: "Numpad2" },
					"│",
					{ ch: " 3 ", f: "ring", code: "Numpad3" },
					"│   │"
				],
				sep: [
					"├───────────┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴─────────┤  ╭───┴───┴───╮  ├───┴───┴───┤ ",
					{ ch: "E", f: "pinky", code: "NumpadEnter" },
					" │"
				]
			},
			{
				content: [
					"│",
					{ ch: " Ctrl ", f: "pinky", code: "ControlLeft" },
					"│",
					{ ch: " ⌘ ", f: "ring", code: "MetaLeft" },
					"│",
					{ ch: " Alt ", f: "ring", code: "AltLeft" },
					"│",
					{ ch: "         Space         ", f: "thumb", code: "Space" },
					"│",
					{ ch: " Alt ", f: "ring", code: "AltRight" },
					"│",
					{ ch: " ⌘ ", f: "ring", code: "MetaRight" },
					"│",
					{ ch: "Fn ", f: "ring" },
					"│",
					{ ch: " Ctrl ", f: "pinky", code: "ControlRight" },
					"│  │",
					{ ch: " ← ", f: "index", code: "ArrowLeft" },
					"│",
					{ ch: " ↓ ", f: "middle", code: "ArrowDown" },
					"│",
					{ ch: " → ", f: "ring", code: "ArrowRight" },
					"│  │",
					{ ch: "   0   ", f: "ring", code: "Numpad0" },
					"│",
					{ ch: " . ", f: "ring", code: "NumpadDecimal" },
					"│   │"
				]
			}
		],
		bottom: [
			"╰──────┴───┴─────┴───────────────────────┴─────┴───┴───┴──────┴──┴───┴───┴───┴──┴───────┴───┴───╯"
		]
	};

	onMount(() => {
		/** @type {(event: KeyboardEvent) => void} */
		const press = ({ code }) => {
			pressedCodes.add(code);
		};
		/** @type {(event: KeyboardEvent) => void} */
		const release = ({ code }) => {
			pressedCodes.delete(code);
		};
		const clear = () => pressedCodes.clear();

		window.addEventListener("keydown", press);
		window.addEventListener("keyup", release);
		window.addEventListener("blur", clear);

		return () => {
			window.removeEventListener("keydown", press);
			window.removeEventListener("keyup", release);
			window.removeEventListener("blur", clear);
		};
	});
</script>

<div class="keyboard-wrapper">
	<div class="line top-line">
		{#each keyboard.top as char}
			<p class="char box-drawing">{char}</p>
		{/each}
	</div>
	{#each keyboard.rows as row}
		<div class="line content">
			{#each row.content as char}
				{#if typeof char === "object"}
					<p class={`char ${char.f}`} class:pressed={pressedCodes.has(char.code)}>{char.ch}</p>
				{:else}
					<p class="char box-drawing">{char}</p>
				{/if}
			{/each}
		</div>
		{#if row.sep != null || row.sep != undefined}
			<div class="line sep">
				{#each row.sep as char}
					{#if typeof char === "object"}
						<p class={`char ${char.f}`} class:pressed={pressedCodes.has(char.code)}>
							{char.ch}
						</p>
					{:else}
						<p class="char box-drawing">{char}</p>
					{/if}
				{/each}
			</div>
		{/if}
	{/each}
	<div class="line bottom-line">
		{#each keyboard.bottom as char}
			<p class="char box-drawing">{char}</p>
		{/each}
	</div>
</div>
<blockquote>this keyboard updates... try typing! •ᴗ•</blockquote>

<style>
	.keyboard-wrapper {
		--cell-height-scale: 1.3;

		display: flex;
		flex-direction: column;
		font-size: 14px;
	}
	.line {
		display: flex;
		align-items: center;
		height: calc(1.07em * var(--cell-height-scale));
	}
	.char {
		margin: 0;
		padding: 0;
		white-space: pre; /* to prevent html ignoring extra spaces in strings */
		color: var(--overlay1);
	}
	.char.box-drawing {
		transform: scaleY(var(--cell-height-scale));
	}
	.char.pinky {
		color: var(--pink);
	}
	.char.ring {
		color: var(--sapphire);
	}
	.char.middle {
		color: var(--green);
	}
	.char.index {
		color: var(--peach);
	}
	.char.thumb {
		color: var(--red);
	}
	.char.pressed {
		/* font-style: italic; */
		font-weight: 900;
		color: var(--rosewater);
	}
	blockquote {
		padding-left: 1rem;
		margin-top: 2rem;
		border-left: 4px solid var(--surface1);
		color: var(--overlay1);
	}

	@media (max-width: 900px) {
		.keyboard-wrapper {
			font-size: 12px;
		}
	}
	@media (max-width: 750px) {
		.keyboard-wrapper {
			font-size: 10px;
		}
	}
</style>
